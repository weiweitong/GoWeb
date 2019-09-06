package minHeap

import (
	"testing"
)

func Test_MinHeap(t *testing.T) {
	heap := NewMinHeap()
	if heap.Length() == 0 {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}

	heap.Insert(&Url{1, "1"})
	heap.Insert(&Url{2, "2"})
	heap.Insert(&Url{3, "3"})
	heap.Insert(&Url{4, "4"})
	heap.Insert(&Url{5, "5"})

	if heap.Length() != 5 {
		t.Error("Insert 5 to Heap Failed")
	} else {
		t.Log("Insert 5 to Heap Pass")
	}

	v, _ := heap.Min()
	if v.Frequency != 1 {
		t.Error("Top is min Failed")
	} else {
		t.Log("Top is min Pass")
	}

	v, _ = heap.DeleteMin()
	if v.Frequency != 1 {
		t.Error("Delete Top min Failed")
	} else {
		t.Log("Delete Top min Pass")
	}

	v, _ = heap.DeleteMin()
	if v.Frequency != 2 {
		t.Error("Delete Top min Failed")
	} else {
		t.Log("Delete Top min Pass")
	}
}
