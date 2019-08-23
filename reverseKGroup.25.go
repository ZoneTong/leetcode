package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	var dummy = &ListNode{}
	dummy.Next = head

	var pre, end = dummy, dummy

	for end.Next != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		var start = pre.Next
		var next = end.Next
		end.Next = nil
		pre.Next = reverse(start)
		start.Next = next
		pre = start
		end = pre
	}
	return dummy.Next
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	var curr = head
	for curr != nil {
		var next = curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}
