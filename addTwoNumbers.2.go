// 2. 两数相加
// https://leetcode-cn.com/problems/add-two-numbers
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head = &ListNode{}
	var parent, l3 *ListNode
	var gte10 int
	for l1 != nil && l2 != nil {
		if parent == nil {
			l3 = head
		} else {
			l3 = &ListNode{}
			parent.Next = l3
		}
		parent = l3

		l3.Val = l1.Val + l2.Val + gte10
		if l3.Val >= 10 {
			l3.Val -= 10
			gte10 = 1
		} else {
			gte10 = 0
		}
		l1, l2 = l1.Next, l2.Next
	}

	if l1 != nil {
		parent.Next = l1

	} else {
		parent.Next = l2
	}
	l3 = parent.Next
	for l3 != nil && gte10 > 0 {
		l3.Val += 1
		if l3.Val >= 10 {
			l3.Val -= 10
			gte10 = 1
		} else {
			gte10 = 0
		}
		parent = l3
		l3 = l3.Next
	}
	if l3 == nil && gte10 > 0 {
		parent.Next = &ListNode{Val: 1}
	}

	return head
}
