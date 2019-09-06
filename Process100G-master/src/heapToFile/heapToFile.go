package heapToFile

import (
	"minHeap"
	"os"
	"strconv"
)

// 将最小堆的内容输出到文件
func HeapToFile(heap *minHeap.MinHeap) {
	f, err := os.OpenFile("./src/data/output.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()
	CheckError(err)

	for heap.Length() != 0 {
		item, _ := heap.DeleteMin()
		_, err = f.Write([]byte("Frequence: " + strconv.FormatInt(item.Frequency, 10) + " | Url: " + item.Name + "\n"))
		CheckError(err)
	}
}



func CheckError(err error) {
	if err != nil {
		//log.Fatal(err)
		panic(err)
	}
}

