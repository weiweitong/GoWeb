package genURL

import (
	"bufio"
	"math/rand"
	"os"
	"regexp"
	"testing"
)

func Test_GenerateURL(t *testing.T) {
	filePath := "./gen.txt"
	// 增加随机性
	numChannel := rand.Intn(10)
	numChannel += 2
	writeTime := rand.Intn(1024 * 1024)
	GenerateURL(filePath, numChannel, writeTime)

	f, _ := os.Open(filePath)
	defer f.Close()
	buf := bufio.NewReader(f)
	r := regexp.MustCompile(`https://(.+)`)
	// 测试URL出现总次数是否与设定的一致
	count := 0
	for line, _, err := buf.ReadLine(); err == nil; line, _, err = buf.ReadLine() {
		s := string(line)
		if m := r.MatchString(s); !m {
			// 测试生成的是否是URL
			t.Error("Generation is Not URL")
		}
		count++
	}
	if count != numChannel * writeTime {
		t.Error("Generation is not equal to the designed.")
	} else {
		t.Log("Generation Pass")
	}
}
