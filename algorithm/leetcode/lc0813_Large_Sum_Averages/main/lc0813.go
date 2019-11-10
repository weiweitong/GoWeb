package main

import "fmt"

func max(x, y float64) float64 {
	if x < y {
		return y
	}
	return x
}

func largestSumOfAverages(A []int, K int) float64 {
	l := len(A)
	sum := make([]int, l+1)
	for i := 1; i <= l; i++ {
		sum[i] = sum[i-1] + A[i-1]
	}
	dp := make([][]float64, l+1)
	for i := range dp {
		dp[i] = make([]float64, K+1)
	}
	for i := 1; i <= l; i++ {
		dp[i][1] = 1.0 * float64(sum[i]) / float64(i)
		for k := 2; k <= K && k <= i; k++ {
			for j := 1; j < i; j++ {
				dp[i][k] = max(dp[i][k], dp[j][k-1] + float64(sum[i]-sum[j]) / float64(i-j))
			}
		}
	}
	return dp[l][K]
}

func main() {
	A := []int{9, 1, 2, 3, 9}
	K := 3
	fmt.Println(largestSumOfAverages(A, K))
}