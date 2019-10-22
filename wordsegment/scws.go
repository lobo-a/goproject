package wordsegment

/*
#cgo CFLAGS: -I${SRCDIR}/libscws -DLOGGING_LEVEL=LL_WARNING -O3 -Wall
#cgo LDFLAGS: -L${SRCDIR}/libscws
// 需根据自己的运行平台进行配置
#include <config.h>
#include <scws.c>
#include <charset.c>
#include <darray.c>
#include <lock.c>
#include <pool.c>
#include <rule.c>
#include <crc32.c>
#include <xdb.c>
#include <xdict.c>
#include <xtree.c>
*/
import "C"
import "unsafe"

type scwsT struct {
	Cscws C.scws_t
}

// 获取scws分词器实例
func NewScws() *scwsT {
	scws := C.scws_new()
	C.scws_set_charset(scws, C.CString("utf8"))
	C.scws_set_debug(scws, C.SCWS_NA)
	C.scws_set_ignore(scws, C.SCWS_YEA)
	s := scwsT{Cscws: scws}
	return &s
}

// 设置字典 etc/scws_dict_utf8.xdb
func (s *scwsT) ScwsSetDict(path string) {
	scws_dict_utf8 := C.CString(path)
	C.scws_set_dict(s.Cscws, scws_dict_utf8, C.SCWS_XDICT_XDB)
	C.free(unsafe.Pointer(scws_dict_utf8))
}

// 设置rules规则 etc/scws_rules_utf8.ini
func (s *scwsT) ScwsSetRules(path string) {
	scws_rules_utf8 := C.CString(path)
	C.scws_set_rule(s.Cscws, scws_rules_utf8)
	C.free(unsafe.Pointer(scws_rules_utf8))
}

// 释放相关资源
func (s *scwsT) SourcesFree() {
	C.scws_free(s.Cscws)
}

// 对str进行分词，并返回权重排序topNum的切片
func (s *scwsT) GetTopSegments(str string, topNum int) (ret []Segment) {
	if topNum < 1 {
		topNum = 20
	}
	xattr := C.CString("")
	text := C.CString(str)
	scws := s.Cscws
	C.scws_send_text(scws, text, C.int(C.strlen(text)))
	tops := C.scws_get_tops(scws, C.int(topNum), xattr)
	defer func() {
		C.scws_free_tops(tops)
		C.free(unsafe.Pointer(xattr))
		C.free(unsafe.Pointer(text))

	}()
	if tops != nil {
		num := 0
		cur := tops
		for {
			num++
			if cur == nil || num > topNum {
				break
			}
			Segment := Segment{Word: C.GoString(cur.word), Weight: float64(cur.weight)}
			ret = append(ret, Segment)
			cur = cur.next
		}
	}
	return
}
