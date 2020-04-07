package linkedlist

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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dumm := &ListNode{Next:head}

	slow := dumm
	fast := dumm
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	delNode := slow.Next
	slow.Next = slow.Next.Next
	delNode.Next = nil

	return dumm.Next
}
