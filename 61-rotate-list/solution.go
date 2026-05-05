package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode{
	//find last node
	if head == nil || head.Next == nil || k == 0{
		return head
	}

	currentNode := head
	n := 0
	done := false
	needToShiftHead := 0
	for !done{
		n++
		if currentNode.Next == nil{
			needToShiftHead = n - (k % n) 
			if needToShiftHead == 0{
				return head
			}
			currentNode.Next = head
			done = true
			break
		}
		currentNode = currentNode.Next
	}

	previousHead := new(ListNode)
	for range needToShiftHead{
		previousHead = head
		head = head.Next
	}

	previousHead.Next = nil
	return head
}
