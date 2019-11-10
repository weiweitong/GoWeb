/*
	github.com/weiweitong/leetcode/
	Thinking and Written by weiweitong
	copyright@weiweitong
*/
package main

import "fmt"

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func convert(s string, numRows int) string {
	if len(s) < 2 || numRows < 2 {
		return s
	}
	div := numRows + (numRows - 2)

	// 将字符按照v字形来切分数组
	var divS []string
	for i := 0; i < len(s); i += div {
		divS = append(divS, s[i:Min(i+div, len(s))])
	}
	res := ""
	for _, each := range divS {
		res += string(each[0])
	}

	for i, j := 1, div-1; i <= j; i,j = i+1, j-1 {
		for _, each := range divS {
			if i < len(each) {
				res += string(each[i])
			}
			if j < len(each) && i < j {
				res += string(each[j])
			}
		}
	}
	return res
}

func main() {
	x := "123"
	fmt.Println(x[0:len(x)-1])
	s := "LEETCODEISHIRING"
	fmt.Println(convert(s, 3))
	fmt.Println(convert(s, 4))
}
