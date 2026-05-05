package main

import "fmt"

func main() {

	node5 := ListNode{
		Val:  5,
		Next: nil,
	}

	node4 := ListNode{
		Val:  4,
		Next: &node5,
	}

	node3 := ListNode{
		Val:  3,
		Next: &node4,
	}

	node2 := ListNode{
		Val:  2,
		Next: &node3,
	}

	node1 := ListNode{
		Val:  1,
		Next: &node2,
	}

	currentNode := rotateRight(&node1, 2000000000)

	fmt.Printf("Current head node value %d and next node value %d", currentNode.Val, currentNode.Next.Val)

}