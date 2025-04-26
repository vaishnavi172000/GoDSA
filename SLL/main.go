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

func main() {
	l := new(SLL)
	l.Create()
	l.Print()
}
