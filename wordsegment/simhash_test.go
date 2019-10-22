package wordsegment

import (
	"testing"
)

func TestGetSimHash(t *testing.T) {
	scws := NewScws()
	scws.ScwsSetDict("../etc/scws_dict_utf8.xdb")
	scws.ScwsSetRules("../etc/scws_rules_utf8.ini")
	defer scws.SourcesFree()
	strs := [][2]string{
		{"2008年中国网络游戏的实际销售收入达183.8亿元人民币，比2007年增长了百分之76.6", "0010111010110101101011110010110001100100000101110110101110011010"},
		{"", "0000000000000000000000000000000000000000000000000000000000000000"},
	}
	for i := 0; i < len(strs); i++ {
		segs := scws.GetTopSegments(strs[i][0], 20)
		simhash := GetSimHash(segs)
		if simhash != strs[i][1] {
			t.Errorf("simhash: %q, want: %q\n", simhash, strs[i][1])
		}
	}
}

func TestHamDist(t *testing.T) {
	h1 := "0010111010110101101011110010110001100100000101110110101110011010"
	h2 := "0000000000000000000000000000000000000000000000000000000000000000"
	dist, err := HamDist(h1, h2)
	if err == nil {
		if dist != 34 {
			t.Errorf("dist:%v, want: 34", dist)  
		}
	}else{
		t.Errorf("error: %v", err)
	}
}