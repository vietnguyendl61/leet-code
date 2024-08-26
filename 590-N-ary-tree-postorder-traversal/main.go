package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func main() {
	root := Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{
						Val: 5,
					},
					{
						Val: 6,
					},
				},
			},
			{
				Val: 2,
			},
			{
				Val: 4,
			},
		},
	}
	fmt.Println(postorder(&root))

}

func postorder(root *Node) []int {
	var result []int
	if root == nil {
		return nil
	}

	if root.Children == nil {
		return []int{root.Val}
	} else {
		for i := range root.Children {
			result = append(result, postorder(root.Children[i])...)
		}
		result = append(result, root.Val)
	}

	return result
}
