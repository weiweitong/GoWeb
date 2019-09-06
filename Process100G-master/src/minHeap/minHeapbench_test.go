package minHeap

import (
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkMinHeap_DeleteMin(b *testing.B) {
	b.StopTimer()
	arr := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		arr[i] = i
	}
	// 打乱顺序
	for i := 0; i < b.N; i++ {
		t := rand.Intn(b.N)
		arr[i], arr[t] = arr[t], arr[i]
	}
	// 建堆
	heap := NewMinHeap()
	for i := 0; i < b.N; i++ {
		s := strconv.Itoa(arr[i])
		heap.Insert(&Url{int64(arr[i]), s})
	}

	b.StartTimer()
	for heap.Length() != 0 {
		_, _ = heap.DeleteMin()
	}
}
