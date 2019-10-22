package wordsegment

import (
	"hash"
	"hash/crc64"
	"strconv"
)

const (
	DEFAULTSIMHASH64   = "0000000000000000000000000000000000000000000000000000000000000000"
	DEFAULTSIMHASHSIZE = 64
)

// 64位hash计算器
func crc64Hasher() hash.Hash64 {
	tab := crc64.MakeTable(crc64.ECMA)
	return crc64.New(tab)
}

// 获取64位simhash
func GetSimHash(segs []Segment) string {
	if l := len(segs); l > 0 {
		ret := []byte(DEFAULTSIMHASH64)
		var vector [DEFAULTSIMHASHSIZE]float64
		h := crc64Hasher()
		for i := 0; i < l; i++ {
			// 分词hash
			h.Reset()
			h.Write([]byte(segs[i].Word))
			hash := strconv.FormatUint(h.Sum64(), 2)
			hash = strPadLeft(hash, DEFAULTSIMHASHSIZE, '0')
			// 加权
			for ii := 0; ii < DEFAULTSIMHASHSIZE; ii++ {
				if hash[ii] == '1' {
					vector[ii] += segs[i].Weight
				} else {
					vector[ii] -= segs[i].Weight
				}
			}
		}
		// 降维，构建simhash
		for i := 0; i < DEFAULTSIMHASHSIZE; i++ {
			if vector[i] > 0 {
				ret[i] = '1'
			}
		}
		return string(ret)
	} else {
		return DEFAULTSIMHASH64
	}
}
