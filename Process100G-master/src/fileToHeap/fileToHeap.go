package fileToHeap

import (
	"minHeap"
	"readSeparate"
	"strconv"
)

// 删除每个小文件中的最小堆，并将他们全部合并成一个含有NumTop个元素的大堆
func Reduce(BatchSize, NumFile, NumTop int, tmpPath string) *minHeap.MinHeap {
	heap := minHeap.NewMinHeap()
	for i := 0; i < NumFile; i++ {
		NextHeap, err := getMinHeapFromFile(tmpPath + "/" + strconv.Itoa(i) + ".txt", BatchSize, NumFile, NumTop, tmpPath)
		if err != nil {
			continue
		}
		heap = mergeTwoHeap(heap, NextHeap, NumTop)
	}
	return heap
}



// 根据文件名，生成对应的最小堆，先将前NumTop个rul和对应堆次数加入堆
// 随后每次新加入堆的url出现的次数必须大于最小堆的堆顶
// 这样最小堆中存放的是最大的NumTop个元素
func getMinHeapFromFile(filePath string, BatchSize, NumFile, NumTop int, tmpPath string) (*minHeap.MinHeap, error) {
	CountMap := make(map[string]int64)

	/*
		此出构造一个函数addToHashMap，为的是对切割后的小文件
		对URL统计频率，并将结果储存到CountMap中
	*/
	var addToHashMap func([]string, string, int)
	addToHashMap = func(keys []string, tmpPath string, NumFile int) {
		for _, key := range keys {
			if _, ok := CountMap[key]; ok {
				CountMap[key]++
			} else if key != "" {
				CountMap[key] = 1
			}
		}
	}

	err := readSeparate.ReadFile(filePath, addToHashMap, BatchSize, NumFile, tmpPath)
	CheckError(err)

	heap := minHeap.NewMinHeap()
	for k, v := range CountMap {
		if heap.Length() < NumTop {
			heap.Insert(&minHeap.Url{v, k})
			continue
		}
		min, _ := heap.Min()
		if min.Frequency <= v {
			heap.DeleteMin()
			heap.Insert(&minHeap.Url{v, k})
		}
	}
	return heap, nil
}

// 合并两个最小堆，得到仅有一个NumTop的最小堆，因为我们最终只需要的最大的NumTop个值
func mergeTwoHeap(oldH, newH *minHeap.MinHeap, NumTop int) *minHeap.MinHeap {
	if newH == nil || newH.Length() == 0 {
		return oldH
	}
	for newH.Length() != 0 {
		value, _ := newH.DeleteMin()
		if oldH.Length() < NumTop {
			oldH.Insert(value)
			continue
		}
		min, _ := oldH.Min()
		if min.Frequency <= value.Frequency {
			oldH.DeleteMin()
			oldH.Insert(value)
		}
	}
	return oldH
}

func CheckError(err error) {
	if err != nil {
		//log.Fatal(err)
		panic(err)
	}
}

