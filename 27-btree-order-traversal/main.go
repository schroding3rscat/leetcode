package main

import (
	"fmt"
)

func main() {
	root := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Printf("%+v\n", levelOrder(root))

	root = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 5},
					Right: nil,
				},
				Right: nil,
			},
			Right: nil,
		},
		Right: nil,
	}
	fmt.Printf("%+v\n", levelOrder(root))
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{{root.Val}}
	levels(&res, root.Left, root.Right)
	return res
}

func levels(res *[][]int, nodes ...*TreeNode) {
	if len(nodes) == 0 {
		return
	}
	children := make([]*TreeNode, 0, len(nodes))
	tmp := []int{}
	for _, n := range nodes {
		if n == nil {
			continue
		}

		tmp = append(tmp, n.Val)
		children = append(children, n.Left, n.Right)
	}
	if len(tmp) > 0 {
		*res = append(*res, tmp)
	}
	levels(res, children...)
}

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
