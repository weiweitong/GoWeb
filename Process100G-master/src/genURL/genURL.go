package genURL

import (
	"math/rand"
	"os"
	"strconv"
	"sync"
)

var mu sync.Mutex

func GenerateURL(filename string, numChannel, writeTime int)  {

	f, err := os.OpenFile(filename, os.O_CREATE | os.O_WRONLY | os.O_TRUNC, 0666 )
	checkError(err)

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

	// int64 max: 9223372036854775807
	// int32 max: 2147483647
	// 1024 * 1024 * 1024 url ~= 100GB 应该是接近100G的样子

	wait := make(chan int)
	for i := 0; i < numChannel; i++ {
		go writeFile(f, urlArr[i%10], wait, writeTime)
	}
	<- wait
}

// 创造数据
func writeFile(f *os.File, url string, wait chan<- int, writeTime int) {
	// fmt.Println("start " + url)
	for j := 100; j < writeTime + 100; j++ {
		a := rand.Intn(j)
		b := []byte(url + strconv.Itoa(a) + "\n")
		// 将所有可以在 File Write 的操作全部放在Lock锁前，
		// 这样让1024的协程早点处理完并等待，快速完成写入后马上释放锁，
		// 虽然写入总有单线程的瓶颈，但速度快了两倍以上
		mu.Lock()
		_, _ = f.Write(b)
		mu.Unlock()
	}
	wait <- 0
}

func checkError(err error) {
	if err != nil {
		//log.Fatal(err)
		panic(err)
	}
}

