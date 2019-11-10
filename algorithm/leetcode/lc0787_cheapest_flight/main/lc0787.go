/*
	github.com/weiweitong/leetcode/
	Thinking and Written by weiweitong
	copyright@weiweitong
*/
package main

import (
	"fmt"
	"math"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	preMoney := make([]int, n)
	curMoney := make([]int, n)
	city := make([][]int, n)
	for i := 0; i < n; i++ {
		preMoney[i] = math.MaxInt32
		curMoney[i] = math.MaxInt32
		city[i] = make([]int, n)
	}
	preMoney[src] = 0
	curMoney[src] = 0
	for i := 0; i < len(flights); i++ {
		city[flights[i][0]][flights[i][1]] = flights[i][2]
	}
	queue := make([]int, 1)
	queue[0] = src
	curRound, nextRound := 1, 0
	for i := 0; i <= K && len(queue) != 0; {
		tmp := queue[0]
		queue = queue[1:]
		curRound--
		for j := 0; j < n; j++ {
			price := city[tmp][j]
			if price > 0 && preMoney[tmp] != math.MaxInt32 && preMoney[tmp]+price < curMoney[j] {
				curMoney[j] = preMoney[tmp]+price
				if j != dst {
					queue = append(queue, j)
					nextRound++
				}
			}
		}
		if curRound == 0 {
			for x := 0; x < n; x++ {
				preMoney[x] = curMoney[x]
			}
			curRound = nextRound
			nextRound = 0
			i++
		}
	}
	if preMoney[dst] < math.MaxInt32 {
		return preMoney[dst]
	}
	return -1
}

func main() {
	n := 3
	edges := [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{0, 2, 500},
	}
	src := 0
	dst := 2
	k := 1
	fmt.Println(findCheapestPrice(n, edges, src, dst, k))

	n = 4
	edges = [][]int {{0,1,1}, {0,2,5}, {1,2,1}, {2,3,1}}
	src = 0
	dst = 3
	k = 1
	fmt.Println(findCheapestPrice(n, edges, src, dst, k))
}
