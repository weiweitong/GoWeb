/*
	github.com/weiweitong/leetcode/
	Thinking and Written by weiweitong
	copyright@weiweitong
*/
package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	arr := convert(head)
	res := make([]int, len(arr))
	stack := make([]int, len(arr))
	pos := -1
	for i := 0; i < len(arr); i++ {
		for pos >= 0 && arr[stack[pos]] < arr[i] {
			res[stack[pos]] = arr[i]
			pos--
		}
		pos++
		stack[pos] = i
	}
	return res
}

func convert(head *ListNode) []int {
	res := make([]int, 0, 1024)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func reconvert(arr []int) *ListNode {
	if len(arr) < 1 {
		return nil
	}
	head := &ListNode{}
	head.Val = arr[0]
	if len(arr) < 2 {
		head.Next = nil
		return head
	}
	var tmp = &ListNode{}
	head.Next = tmp
	for i := 1; i < len(arr); i++ {
		tmp.Val = arr[i]
		if i+1 < len(arr) {
			tmp.Next = &ListNode{}
			tmp = tmp.Next
		} else {
			tmp.Next = nil
			break
		}
	}
	return head
}

func main() {
	arr := []int{2, 7, 3, 4, 5, 4, 3, 5}
	head := reconvert(arr)
	res := nextLargerNodes(head)
	fmt.Println(res)
}
