package wordsegment

import (
	"errors"
	"bytes"
)

type Segment struct {
	Word   string
	Weight float64
}

type Tokenizer interface {
	GetTopSegments(str string, topNum int) (ret []Segment)
	SourcesFree()
}

type SimHasher interface {
	GetSimHash(segs []Segment) string
}
// 补充字符到lenth长度
func strPadLeft(str string, length int, sep byte) string {
	if l := len(str); l < length {
		prefix := bytes.Repeat([]byte{sep}, length-l)
		buf := bytes.NewBuffer(prefix)
		buf.WriteString(str)
		return buf.String()
	} else {
		return str
	}
}
// 海明距离
func HamDist(h1, h2 string) (int, error){
	if l := len(h1); l == len(h2) {
		dist := 0
		for i := 0; i < l; i++ {
			if h1[i] != h2[i] {
				dist++
			}
		}
		return dist, nil
	}else{
		return 0, errors.New("diff hashs length")
	}
}
