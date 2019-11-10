/*
	github.com/weiweitong/leetcode/
	Thinking and Written by weiweitong
	copyright@weiweitong
*/
package main

import "fmt"

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	dp := make([][]*TreeNode, n+1)
	dp[0] = make([]*TreeNode, 0)
	if n == 0 {
		return dp[0];
	}
	dp[0] = append(dp[0], nil)
	for len := 1; len <= n; len++ {
		dp[len] = make([]*TreeNode, 0)
		for root := 1; root <= len; root++ {
			left := root - 1
			right := len - root
			for _, leftTree := range dp[left] {
				for _, rightTree := range dp[right] {
					treeRoot := &TreeNode{root, leftTree, clone(rightTree, root)}
					dp[len] = append(dp[len], treeRoot)
				}
			}
		}
	}
	return dp[n]
}

func clone(n *TreeNode, bias int) *TreeNode {
	if n == nil {
		return nil
	}
	node := &TreeNode{n.Val + bias, clone(n.Left, bias), clone(n.Right, bias)}
	return node
}

func main() {
	n := 3
	treeArr := generateTrees(n)
	//fmt.Printf("%v", treeArr)
	fmt.Println("[ ")
	for _, node := range treeArr {
		fmt.Print("  [ ")
		printTree(node)
		fmt.Println("]")
	}
	fmt.Println("]")
}

func printTree(node *TreeNode) {
	if node == nil {
		fmt.Print("null ")
		return
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, node)
	for len(queue) > 0 {
		tree := queue[0]
		queue = queue[1:]
		if tree == nil {
			fmt.Print("null ")
			continue
		}
		fmt.Printf("%v ", tree.Val)
		queue = append(queue, tree.Left)
		queue = append(queue, tree.Right)
	}
}
