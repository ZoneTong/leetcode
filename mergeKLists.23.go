package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	} else if l == 1 {
		return lists[0]
	}

	ln0 := mergeKLists(lists[:l>>1])
	ln1 := mergeKLists(lists[l>>1:])
	var head, pre *ListNode
	if ln0 == nil {
		return ln1
	}

	if ln1 != nil && ln0.Val > ln1.Val {
		ln0, ln1 = ln1, ln0
	}

	head = ln0
	pre = head
	ln0 = head.Next
	for ln0 != nil && ln1 != nil {
		if ln0.Val > ln1.Val { // 插入
			pre.Next = ln1
			pre = ln1
			ln1 = ln1.Next
			pre.Next = ln0
			// printList(head)
		} else {
			pre = ln0
			ln0 = ln0.Next
		}
	}

	if ln0 == nil {
		pre.Next = ln1
	}

	return head
}

func printList(list *ListNode) {
	for list != nil {
		fmt.Printf("%d->", list.Val)
		list = list.Next
	}
	fmt.Println()
}
