package main

import(
	funcs "DataStructure/functions"
)

func main() {
	// Test: myarray----------------
	// funcs.MyArrayInt()
	// funcs.MyArrayString()
	// funcs.MyArrayStudent()
	// funcs.MyArrayTest()

	// Test: stack -----------------
	// funcs.TestStack()
	// funcs.TestStack2("()[]{([])}")

	// Test: arrayQueue-------------
	// funcs.TestArrayQueue()

	// Test: LoopQueue--------------
	// funcs.TestLoopQueue()
	// funcs.TestLoopQueue2()

	// Test: LinkedList-------------
	// funcs.TestLinkedList()

	// Test: LinkedListStack--------
	// funcs.TestLinkedListStack()

	// Test: LinkedListQueue--------
	// funcs.TestLinkedListQueue()

	// Test: test efficiency--------
	// testEfficiency()
	
	// Test: Recursion - Sum
	// funcs.TestSum()

	
}

func testEfficiency() {
	cnt := 100000
	funcs.ArrayQueueEfficiency(cnt)
	funcs.LoopQueueEfficiency(cnt)
	funcs.LinkedListQueueEfficiency(cnt)
	funcs.ArrayStackEfficiency(cnt)
	funcs.LinkedListStackEfficiency(cnt)

	/*
	Result : 
		ArrayQueue Use time :  10.203s
		LoopQueue Use time :  22ms
		LinkedListQueue Use time :  7ms
		ArrayStack Use time :  27ms
		LinkedListStack Use time :  13ms
	*/
}

// leetcode : 203.Remove Linked List Elements
/*
Remove all elements from a linked list of integers that have value val.

Example:

Input:  1->2->6->3->4->5->6, val = 6
Output: 1->2->3->4->5
*/
type ListNode struct {
	Next *ListNode
	Val int
}

func removeElements(head *ListNode, val int) *ListNode {
    for head != nil && head.Val == val {
		delNode := head
		head = head.Next
		delNode.Next = nil
	}

	if head == nil {
		return head
	}

	prev := head;
	for prev.Next != nil {
		if prev.Next.Val == val {
			delNode := prev.Next
			prev.Next = delNode.Next
			delNode.Next = nil
		} else {
			prev = prev.Next
		}
	}
	return head
}

func removeElements2(head *ListNode, val int) *ListNode {
    for head != nil && head.Val == val {
		head = head.Next
	}

	if head == nil {
		return head
	}

	prev := head;
	for prev.Next != nil {
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}
	return head
}

func removeElements3(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{nil, -1}
	dummyHead.Next = head

	prev := dummyHead
	for prev.Next != nil {
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}
	return dummyHead.Next
}

// 递归方法：
func removeElements4(head *ListNode, val int) *ListNode{
	if head == nil {
		return head
	}
	head.Next = removeElements4(head.Next, val)
	if head.Val == val {
		return head.Next
	} else {
		return head
	}
}