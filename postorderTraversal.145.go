package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	var result = make([]int, 0, 4)
	if root == nil {
		return result
	}

	var stack = make([]*TreeNode, 0, 4)
	stack = append(stack, root)
	top := 0
	var pre *TreeNode
	peek := func() *TreeNode {
		return stack[top]
	}

	pop := func() *TreeNode {
		n := stack[top]
		stack = stack[:top]
		top--
		return n
	}

	push := func(n *TreeNode) {
		stack = append(stack, n)
		top++
	}

	for top >= 0 {
		peeknode := peek()
		for peeknode.Left != nil || peeknode.Right != nil {
			if peeknode.Left != nil {
				push(peeknode.Left)
				peeknode = peeknode.Left
			} else {
				push(peeknode.Right)
				peeknode = peeknode.Right
			}

		}

		for {
			node := pop()
			result = append(result, node.Val)
			pre = node

			if top < 0 {
				break
			}

			peeknode = peek()
			if peeknode.Left == pre && peeknode.Right != nil {
				push(peeknode.Right)
				break
			}
		}
	}
	return result
}
