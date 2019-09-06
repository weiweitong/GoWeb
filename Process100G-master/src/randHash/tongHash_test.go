package randHash

import (
	"testing"
)



func Test_TongHash64(t *testing.T) {

	urlArr := [10]string{
		"https://www.google.com/search?query=pingcap/search?query=",
		"https://www.infoq.cn/article/jjp0c2bR4*Ulld0wb88r/search?query=",
		"https://www.amap.com/search?query=pingcap&city=110000&geoobj=116.354861%7C40.038564%7C116.358089%7C40.045304&zoom=18/search?query=",
		"https://pingcap.com/about-cn/recruit/engineering/qa-engineer/search?query=",
		"https://pingcap.com/about-cn/recruit/join/#campus/search?query=",

		"https://github.com/weiweitong/tidb/blob/master/bindinfo/bind_test.go/search?query=",
		"https://www.nowcoder.com/discuss/224008?type=post&order=create&pos=&page=1/search?query=",
		"https://www.jianshu.com/p/4bd4f88e24e4/search?query=",
		"https://github.com/aylei/interview/search?query=",
		"https://www.zhihu.com/search?type=people&q=Ed%20Huang",
	}

	if TongHash64(urlArr[0] + "37648") != uint64(663337119048726480) {
		t.Error("URL 1 Failed")
	} else {
		t.Log("URL 1 Pass")
	}
	if TongHash64(urlArr[1] + "539") != uint64(1951549810135895647) {
		t.Error("URL 2 Failed")
	} else {
		t.Log("URL 2 Pass")
	}
	if TongHash64(urlArr[2] + "4683") != uint64(6595538339775324920) {
		t.Error("URL 3 Failed")
	} else {
		t.Log("URL 3 Pass")
	}
	if TongHash64(urlArr[3] + "7842") != uint64(3541448068784443734) {
		t.Error("URL 4 Failed")
	} else {
		t.Log("URL 4 Pass")
	}
	if TongHash64(urlArr[4] + "6739") != uint64(3695940101946128463) {
		t.Error("URL 5 Failed")
	} else {
		t.Log("URL 5 Pass")
	}



	if TongHash64(urlArr[5] + "5922") != uint64(1497386836596738461) {
		t.Error("URL 6 Failed")
	} else {
		t.Log("URL 6 Pass")
	}

	if TongHash64(urlArr[6] + "7311") != uint64(8005963952974628570) {
		t.Error("URL 7 Failed")
	} else {
		t.Log("URL 7 Pass")
	}

	if TongHash64(urlArr[7] + "8913") != uint64(752487054441504907) {
		t.Error("URL 8 Failed")
	} else {
		t.Log("URL 8 Pass")
	}

	if TongHash64(urlArr[8] + "90377") != uint64(3318101722004543183) {
		t.Error("URL 9 Failed")
	} else {
		t.Log("URL 9 Pass")
	}

	if TongHash64(urlArr[9] + "67429") != uint64(8663923265148642088) {
		t.Error("URL 10 Failed")
	} else {
		t.Log("URL 10 Pass")
	}
}
