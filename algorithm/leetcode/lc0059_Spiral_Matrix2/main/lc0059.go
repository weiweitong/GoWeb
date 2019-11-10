/*
	github.com/weiweitong/leetcode/
	Thinking and Written by weiweitong
	copyright@weiweitong
*/
package main

import "fmt"

func generateMatrix(n int) [][]int {
	if n < 2 {
		return [][]int{{1}}
	}
	arr := make([][]int, n)
	for i := range arr {
		arr[i] = make([]int, n)
	}
	count := 1
	round := 0
	end := n * n
	i, j := 0, 0
	for count <= end {
		for count <= end && j < n-round {
			arr[i][j] = count
			count++
			j++
		}
		j--
		i++
		for count <= end && i < n-round {
			arr[i][j] = count
			count++
			i++
		}
		i--
		j--
		for count <= end && j >= round {
			arr[i][j] = count
			count++
			j--
		}
		j++
		i--
		for count <= end && i > round {
			arr[i][j] = count
			count++
			i--
		}
		i++
		j++
		fmt.Println(arr)
		round++
	}
	return arr
}


func main() {
	fmt.Println( generateMatrix(5))
}
