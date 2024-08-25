package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}

	fmt.Println(postorderTraversal(&root))

}

func postorderTraversal(root *TreeNode) []int {
	var result = make([]int, 0)

	if root == nil {
		return result
	}

	left := postorderTraversal(root.Left)
	result = append(result, left...)

	right := postorderTraversal(root.Right)
	result = append(result, right...)

	result = append(result, root.Val)

	return result
}
