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
		if t[i][1] <= t[j][1] {
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

func main() {
	var arr twoDim = [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	sort.Sort(arr)
	fmt.Println(arr)
}
