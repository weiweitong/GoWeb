/*
	github.com/weiweitong/leetcode/
	Thinking and Written by weiweitong
	copyright@weiweitong
*/
package main

import "fmt"

func stoneGameII(piles []int) int {
	sum := make([]int, len(piles)+1)
	for i := len(piles)-1; i >= 0; i-- {
		sum[i] = sum[i+1] + piles[i]
	}
	hasCal := make(map[[2]int]int)
	return dfs(piles, sum, 0, 1, hasCal)
}

// idx 从当前位置开始拿石子
// M 题目中的M
// sum 当前idx开始的石子数
// hasCal 如果已经计算过，则直接返回
func dfs(piles, sum []int, idx, M int, hasCal map[[2]int]int) int {
	if v, ok := hasCal[[2]int{idx, M}]; ok {
		return v
	}
	if idx >= len(piles) {
		return 0
	}
	if idx + M*2 >= len(piles) {
		return sum[idx]
	}
	maxT := 0
	for i := 1; i <= 2*M; i++ {
		maxT = max(maxT, sum[idx] - dfs(piles, sum, idx+i, max(i, M), hasCal))
	}
	hasCal[[2]int{idx, M}] = maxT
	return maxT
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	piles := []int{2, 7, 9, 4, 4}

	fmt.Println(stoneGameII(piles))
}
