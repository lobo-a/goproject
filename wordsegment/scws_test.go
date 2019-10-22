package wordsegment

import (
	"testing"
)

func TestScwsGetTops(t *testing.T) {
	str := "2008年中国网络游戏的实际销售收入达183.8亿元人民币，比2007年增长了百分之76.6"
	scws := NewScws()
	scws.ScwsSetDict("../etc/scws_dict_utf8.xdb") // 使用相对路径时，需注意为你执行文件的相对路径
	scws.ScwsSetRules("../etc/scws_rules_utf8.ini")
	defer scws.SourcesFree()
	t.Log(scws.GetTopSegments(str, 10))
}

func BenchmarkScwsGetTops(b *testing.B) {
	b.StopTimer()
	// 48979	     24524 ns/op	     832 B/op	      15 allocs/op  创建一次scws性能
	// 2034	    573132 ns/op	     832 B/op	      15 allocs/op。   每次都创建scws性能
	str := "2008年中国网络游戏的实际销售收入达183.8亿元人民币，比2007年增长了百分之76.6"
	scws := NewScws()
	scws.ScwsSetDict("../etc/scws_dict_utf8.xdb")
	scws.ScwsSetRules("../etc/scws_rules_utf8.ini")
	defer scws.SourcesFree()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		scws.GetTopSegments(str, 10)
	}
}
