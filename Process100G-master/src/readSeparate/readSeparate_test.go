package readSeparate

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"testing"
)

// 每次运行前，先运行genURL的test文件，然后删除本包下的tmp文件夹
func TestReadFile(t *testing.T) {
	filePath := "../genURL/gen.txt"
	BatchSize := 1024 * 1024 * 100
	NumFile := 100
	tmpPath := "./tmp"
	err := ReadFile(filePath, SetPartition, BatchSize, NumFile, tmpPath)
	if err != nil {
		t.Error("Read file is wrong")
	}

	// 增加随机性
	ra := rand.Intn(NumFile)
	tmpPath = tmpPath + "/" + strconv.Itoa(ra) + ".txt"
	f, err := os.Open(tmpPath)
	if err != nil {
		t.Error("Cannot open tmp txt")
	} else {
		t.Log("open tmp txt success")
	}
	defer f.Close()
	buf := bufio.NewReader(f)

	// 测试一个URL在小文件中出现次数是否与大文件中出现次数的一致，
	// 不妨取随机一个小文件中的第一行URL作为测试
	line, _, err := buf.ReadLine()
	target := string(line)
	targetSum := 0
	for ; err == nil; line, _, err = buf.ReadLine() {
		s := string(line)
		if s == target {
			targetSum++
		}
	}
	// 求出总文件中出现的次数
	f2, err := os.Open(filePath)
	if err != nil {
		t.Error("Cannot open file txt")
	} else {
		t.Log("open file txt success")
	}
	defer f2.Close()
	buf2 := bufio.NewReader(f2)
	count := 0
	for l, _, err := buf2.ReadLine(); err == nil; l, _, err = buf2.ReadLine() {
		s := string(l)
		if s == target {
			count++
		}
	}
	//t.Log(ra)
	//t.Log(target)
	//t.Log(count)
	//t.Log(targetSum)
	if count != targetSum && count+1 != targetSum && count-1 != targetSum {
		t.Error("After Separate is not equal to the designed.")
	} else {
		t.Log("Generation Pass")
	}
}
