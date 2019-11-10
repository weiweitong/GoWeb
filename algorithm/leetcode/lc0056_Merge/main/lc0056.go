/*
	github.com/weiweitong/leetcode/
	Thinking and Written by weiweitong
	copyright@weiweitong
*/
package main

import (
	"fmt"
	"sort"
)

type twoDim [][]int

func (t twoDim) Len() int {
	return len(t)
}

func (t twoDim) Less(i, j int) bool {
	if t[i][0] == t[j][0] {
		if t[i][1] >= t[j][1] {
			return true
		}
	} else if t[i][0] < t[j][0] {
		return true
	}
	return false
}

func (t twoDim) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func merge(intervals [][]int) [][]int {
	length := len(intervals)
	if length < 2 {
		return intervals
	}
	list := make(twoDim, length)
	res := make([][]int, 0)
	for i := range list {
		list[i] = intervals[i]
	}
	sort.Sort(list)

	for i, next := 0, 0; i < length; i = next {
		for j := i+1; ; j++ {
			if j == length || list[i][1] < list[j][0] {
				res = append(res, list[i])
				next = j
				break
			}
			list[i][1] = max(list[i][1], list[j][1])
		}
	}

	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	arr := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	fmt.Println(merge(arr))

	a2 := [][]int{
		{1, 4},
		{4, 5},
	}
	fmt.Println(merge(a2))
}
