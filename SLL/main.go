package main

import "fmt"

// Node structure for SLL
type Node struct {
	val  int
	next *Node
}

// SLL structure
type SLL struct {
	// head of the SLL
	head *Node
}

// getNode function to create a new node
func (n *Node) getNode(val int) {
	n.val = val
}

// Create function to create a SLL
func (l *SLL) Create() {
	var val int
	var temp *Node
	for {
		fmt.Println("Enter Node value or -1 to stop")
		fmt.Scan(&val)
		if val == -1 {
			return
		}
		// if head is nil
		// then create a new node and assign it to head
		if l.head == nil {
			l.head = new(Node)
			l.head.getNode(val)
			temp = l.head
			continue
		}
		temp.next = new(Node)
		temp = temp.next
		temp.getNode(val)
	}
}

// Print function to print the SLL
func (l *SLL) Print() {
	temp := l.head
	for temp != nil {
		fmt.Println(temp.val)
		temp = temp.next
	}
}

// https://leetcode.com/problems/merge-two-sorted-lists/description/?envType=problem-list-v2&envId=linked-list
// You are given the heads of two sorted linked lists list1 and list2.
// Merge the two lists into one sorted list.
// The list should be made by splicing together the nodes of the first two lists.
// Return the head of the merged linked list.
func mergeTwoLists(list1 *Node, list2 *Node) *Node {
	var Dummy, temp *Node
	Dummy = new(Node)
	temp = Dummy
	for list1 != nil && list2 != nil {
		if list1.val > list2.val {
			temp.next = list2
			temp = temp.next
			list2 = list2.next
		} else {
			temp.next = list1
			temp = temp.next
			list1 = list1.next
		}
	}
	if list1 != nil {
		temp.next = list1
	}
	if list2 != nil {
		temp.next = list2
	}
	return Dummy.next

}

func main() {
	l := new(SLL)
	l.Create()
	l.Print()
}
