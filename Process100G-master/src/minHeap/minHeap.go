package minHeap

import (
	"fmt"
	"math"
)

/*
最小堆采用数组形式，其中每个元素包含url网址和其频率
为了对该ArrU数组进行访问，
 */
type MinHeap struct {
	ArrU []*Url
}

// Url结构里包含它的名字和总次数
type Url struct {
	Frequency int64
	Name      string
}

// 最小堆的初始化，返回一个数组的封装
// 数组内仅有一个不含任何信息的节点，此节点仅充当标志物
func NewMinHeap() *MinHeap {
	first := &Url{math.MinInt64, "None"}
	h := &MinHeap{ArrU: []*Url{first}}
	return h
}

// 返回最小堆的总长度，为实际总长度减一
func (H *MinHeap) Length() int {
	return len(H.ArrU) - 1
}

// 返回最小堆的堆顶，即返回它的最小值
func (H *MinHeap) Min() (*Url, error) {
	if len(H.ArrU) > 1 {
		return H.ArrU[1], nil
	}
	return nil, fmt.Errorf("heap is empty")
}

/*
最小堆中插入元素时，先插入到最后
再将最末节点与它的母节点(即i/2)相比较
 */
func (H *MinHeap) Insert(v *Url) {

	H.ArrU = append(H.ArrU, v)
	i := len(H.ArrU) - 1
	for ; (H.ArrU[i/2]).Frequency > v.Frequency; i /= 2 {
		H.ArrU[i] = H.ArrU[i/2]
	}

	H.ArrU[i] = v
}

/*
删除最小堆堆顶，并返回堆顶的元素
最小堆的删除是将它的最小值与最末尾值互换，
再将最小堆顶点下沉，在数组中即每次与i*2和i*2+1的两个子节点比较，
与两个子节点中较小的那个子节点交换，然后继续下沉操作，直到到达最底层(没有子节点了)
然后删除掉最末尾的值
 */
func (H *MinHeap) DeleteMin() (*Url, error) {
	if len(H.ArrU) <= 1 {
		return nil, fmt.Errorf("MinHeap is empty")
	}
	minElement := H.ArrU[1]
	lastElement := H.ArrU[len(H.ArrU)-1]
	var i, child int
	for i = 1; i*2 < len(H.ArrU); i = child {
		child = i * 2
		if child < len(H.ArrU)-1 && H.ArrU[child+1].Frequency < H.ArrU[child].Frequency {
			child++
		}
		if lastElement.Frequency > H.ArrU[child].Frequency {
			H.ArrU[i] = H.ArrU[child]
		} else {
			break
		}
	}
	H.ArrU[i] = lastElement
	H.ArrU = H.ArrU[:len(H.ArrU)-1]
	return minElement, nil
}
