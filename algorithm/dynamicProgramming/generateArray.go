package main

import "fmt"

/*
problem:
[1, 2, 3, 4] ==> products: [24, 12, 8, 6]
a1, a2, a3, a4

Algorithm: O(n) time, O(1) space.

use dynamic programming:

the first round:
brr ==>
1 => a1 => a1*a2 => a1*a2*a3
get b4
tmp = a4

b3 = a1 * a2 * a4 = b3 * tmp
tmp *= a3 = a3 * a4

b2 = b2 * a3
tmp *= a2 = a2 * a3 * a4

b1 = b1 * tmp

*/
func generate(arr []int) []int {
	l := len(arr)
	brr := make([]int, l)

	brr[0] = 1
	for i := 1; i < l; i++ {
		brr[i] = arr[i-1] * brr[i-1]
	}

	// brr: 1, arr[0], arr[0] * arr[1], arr[0]*arr[1]*arr[2]

	tmp := 1

	for i := l-1; i >= 0; i-- {
		brr[i] = brr[i] * tmp
		tmp *= arr[i]
	}

	return brr
}

func main() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{1, 2, 0, 4}

	fmt.Println(generate(arr1))
	fmt.Println("")
	fmt.Println(generate(arr2))
}



