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

func profitableSchemes(G int, P int, group []int, profit []int) int {
	var MOD int64 = int64(math.Pow10(9)) + 7
	N := len(group)

	dp := make([][][]int64, 2)

	for i := range dp {
		dp[i] = make([][]int64, P+1)
		for j := range dp[i] {
			dp[i][j] = make([]int64, G+1)
		}
	}
	dp[0][0][0] = 1

	for idx := 0; idx < N; idx++ {
		p0 := profit[idx]	// the current crime profit
		g0 := group[idx]

		cur := dp[idx%2]
		cur2 := dp[(idx+1)%2]

		for jp := 0; jp <= P; jp++ {
			for jg := 0; jg <= G; jg++ {
				cur2[jp][jg] = cur[jp][jg]
			}
		}

		for p1 := 0; p1 <= P; p1++ { // p1 the current profit
			p2 := min(p0+p1, P); // p2: the new profit after committing this crime
			for g1 := 0; g1 <= G - g0; g1++ {
				g2 := g0 + g1
				cur2[p2][g2] += cur[p1][g1]
				cur2[p2][g2] = cur2[p2][g2] % MOD
			}
		}
	}

	var ans int64

	for _, x := range dp[N%2][P] {
		ans += x
	}
	return int(ans % MOD)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	G, P := 5, 3
	group := []int{2, 2}
	profit := []int{2, 3}
	fmt.Println(profitableSchemes(G, P, group, profit))

	G, P = 10, 5
	group = []int{2, 3, 5}
	profit = []int{6, 7, 8}
	fmt.Println(profitableSchemes(G, P, group, profit))
}
