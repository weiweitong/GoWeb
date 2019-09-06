package fileToHeap

import (
	"bufio"
	"os"
	"randHash"
	"strconv"
	"testing"
)


// 运行前，先运行genURL和readSeparate中的Test文件
// 测试是否可以合并得出含有100个最大元素的最小堆
func TestReduce(t *testing.T) {

	tmpPath := "../readSeparate/tmp"
	BatchSize := 1024 * 1024 * 100
	NumFile, NumTop := 100, 100
	heap := Reduce(BatchSize, NumFile, NumTop, tmpPath)

	for heap.Length() != 0 {
		a, err := heap.DeleteMin()
		if err != nil {
			t.Error("wrong with deleteMin")
		} else {
			freq := int(a.Frequency)
			url := a.Name
			partition := tmpPath + "/" + strconv.Itoa(int(randHash.TongHash64(url)) % NumFile) + ".txt"
			// 求出该小文件中出现的次数
			f, err := os.Open(partition)
			if err != nil {
				t.Error("Cannot open file txt")
				return
			} else {
				t.Log("open file txt success")
			}
			buf := bufio.NewReader(f)
			count := 0
			for l, _, err := buf.ReadLine(); err == nil; l, _, err = buf.ReadLine() {
				s := string(l)
				if s == url {
					count++
				}
			}
			if count != freq {
				t.Error("Cannot match the biggest in heap with the tmp file")
			} else {
				t.Log("success match the biggest in heap with the tmp file")
			}
			f.Close()
		}

	}

}
