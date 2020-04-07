package linkedlist

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dum := &ListNode{Next: head}

	cur, next := dum.Next, dum.Next.Next
	for next != nil {
		cur.Next = next.Next
		next.Next = dum.Next
		dum.Next = next

		next = cur.Next
	}

	return dum.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	dum := &ListNode{Next: head}

	prev, cur := dum, dum.Next
	for cur != nil {
		if cur.Val == val {
			delNode := cur
			prev.Next = cur.Next
			delNode.Next = nil

			cur = prev.Next
		} else {
			prev = prev.Next
			cur = cur.Next
		}

	}

	return dum.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func OddEvenList(head *ListNode) *ListNode {
	dum := &ListNode{Next: head}
	odd := &ListNode{}
	oddStart := odd

	idx := -1

	for dum.Next != nil {
		idx++

		if idx%2 == 0 { // even
			dum = dum.Next
		} else { // odd
			node := dum.Next
			dum.Next = dum.Next.Next
			node.Next = nil

			odd.Next = node
			odd = odd.Next
		}
	}

	dum.Next = oddStart.Next

	return head
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func IsPalindrome(head *ListNode) bool {

	return false
}

func (node *ListNode) String() string {
	res := ""
	head := node
	idx := 0
	for head != nil && idx < 50 {
		res += fmt.Sprintf("%v-> ", head.Val)
		head = head.Next
		idx++
	}

	return res
}
