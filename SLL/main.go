package main

import (
	"fmt"
	"math"
)

// Node structure for SLL
type Node struct {
	val  int
	next *Node
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// SLL structure
type SLL struct {
	// head of the SLL
	head *Node
}

// 1
// getNode function to create a new node
func (n *Node) getNode(val int) {
	n.val = val
}

// 2.
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

// 3.
// Print function to print the SLL
func (l *SLL) Print() {
	temp := l.head
	for temp != nil {
		fmt.Println(temp.val)
		temp = temp.next
	}
}

// 4.
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

// 5.
// https://leetcode.com/problems/rotate-list/?envType=problem-list-v2&envId=linked-list
// Given the head of a linked list, rotate the list to the right by k places.
// Example 1:

// Input: head = [1,2,3,4,5], k = 2
// Output: [4,5,1,2,3]
func rotateRight(head *Node, k int) *Node {
	fast := head
	//slow := head;
	Dummy := new(Node)
	Dummy.next = head
	slow := Dummy
	if head == nil || head.next == nil {
		return head
	}
	n := 0
	for fast != nil {
		fast = fast.next
		n++
	}
	k = k % n
	if k == 0 {
		return head
	}

	fast = Dummy
	for i := 0; i < k; i++ {
		fast = fast.next
	}
	for fast.next != nil {
		fast = fast.next
		slow = slow.next
	}

	fast = slow.next
	slow.next = nil

	slow = fast

	for fast != nil && fast.next != nil {
		fast = fast.next
	}
	fast.next = head
	return slow

}

// 6.
// Given the head of a sorted linked list,
// delete all duplicates such that each element appears only once.
// Return the linked list sorted as well.
// Input: head = [1,1,2]
// Output: [1,2]

func deleteDuplicates(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	temp := head
	for temp.next != nil {
		temp2 := temp
		key := temp.val
		for temp.next != nil && temp.next.val == key {
			temp = temp.next
		}
		if temp2 != temp {
			temp2.next = temp.next
		} else {
			temp = temp.next
		}
	}
	return head

}

//7.
// https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/description/?envType=problem-list-v2&envId=linked-list
// Given the head of a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list. Return the linked list sorted as well.

// Example 1:

// Input: head = [1,2,3,3,4,4,5]
// Output: [1,2,5]
// Example 2:

// Input: head = [1,1,1,2,3]
// Output: [2,3]

func deleteDuplicates2(head *Node) *Node {
	temp := head
	Dummy := new(Node)
	prev := Dummy
	Dummy.next = head
	for temp != nil && temp.next != nil {
		temp1 := temp
		key := temp1.val
		if temp.val == key && temp.next.val == key {
			for temp != nil && temp.val == key {
				temp = temp.next
			}
			prev.next = temp
		} else {
			prev = temp
			temp = temp.next
		}

	}
	return Dummy.next

}

// 8
// https://leetcode.com/problems/swap-nodes-in-pairs/?envType=problem-list-v2&envId=linked-list
// Given a linked list, swap every two adjacent nodes and return its head.
// You must solve the problem without modifying the values in the list's nodes
// (i.e., only nodes themselves may be changed.)
// Example 1:
// Input: head = [1,2,3,4]
// Output: [2,1,4,3]

// Explanation:
// Example 2:
// Input: head = []
// Output: []

// Example 3:
// Input: head = [1]
// Output: [1]

// Example 4:
// Input: head = [1,2,3]
// Output: [2,1,3]

func swapPairs(head *Node) *Node {
	Dummy := new(Node)
	Dummy.next = head
	temp := head
	for temp != nil && temp.next != nil {
		ele := temp.next.val
		temp.next.val = temp.val
		temp.val = ele
		temp = temp.next.next
	}
	return Dummy.next

}

// 9
func reverse(head *Node) *Node {
	curr := head
	var prev *Node
	var next *Node
	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next

	}
	return prev

}

// 10
func FindKthNode(head *Node, k int) *Node {
	temp := head
	cnt := 1
	for cnt != k {
		if temp == nil {
			return nil
		}
		temp = temp.next
		cnt++
	}
	return temp
}

// 11
// https://leetcode.com/problems/reverse-nodes-in-k-group/?envType=problem-list-v2&envId=linked-list
// Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.
// k is a positive integer and is less than or equal to the length of the linked list.
// If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.

// You may not alter the values in the list's nodes, only nodes themselves may be changed.

// Example 1:

// Input: head = [1,2,3,4,5], k = 2
// Output: [2,1,4,3,5]
// Example 2:

// Input: head = [1,2,3,4,5], k = 3
// Output: [3,2,1,4,5]
func reverseKGroup(head *Node, k int) *Node {

	temp := head
	Dummy := new(Node)
	next := head
	prev := Dummy
	for temp != nil {
		kthNode := FindKthNode(temp, k)
		if kthNode == nil {
			prev.next = temp
			break
		}
		next = kthNode.next
		kthNode.next = nil
		reverse(temp)
		prev.next = kthNode
		prev = temp
		temp = next
	}

	return Dummy.next

}

// 12
// https://leetcode.com/problems/reverse-linked-list-ii/description/?envType=problem-list-v2&envId=linked-list
// Given the head of a singly linked list and two integers left and right where
// left <= right, reverse the nodes of the list from position left to position right, and return the reversed list.\
// Example 1:
// Input: head = [1,2,3,4,5], left = 2, right = 4
// Output: [1,4,3,2,5]
// Example 2:

// Input: head = [5], left = 1, right = 1
// Output: [5]
func reverseBetween(head *Node, left int, right int) *Node {
	Dummy := new(Node)
	Dummy.next = head
	prev := Dummy
	temp := head
	cnt := 1
	if left == right {
		return head
	}
	for temp != nil && cnt != left {
		prev = temp
		temp = temp.next
		cnt++
	}
	leftNode := temp
	for cnt != right && temp != nil {
		cnt++
		temp = temp.next
	}
	next := temp.next
	temp.next = nil
	prev.next = reverse(leftNode)
	leftNode.next = next
	return Dummy.next

}

// 13
// https://leetcode.com/problems/linked-list-cycle/
// Given head, the head of a linked list, determine if the linked list has a cycle in it.

// There is a cycle in a linked list if there is some node in the list that can be reached
// again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

// Return true if there is a cycle in the linked list. Otherwise, return false.
func hasCycle(head *Node) bool {
	if head == nil || head.next == nil {
		return false
	}
	fast := head
	slow := head
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			return true
		}

	}
	return false

}

// 14
// https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/?envType=problem-list-v2&envId=linked-list
// Given head which is a reference node to a singly-linked list.
// The value of each node in the linked list is either 0 or 1.
// The linked list holds the binary representation of a number.
// Return the decimal value of the number in the linked list.
// The most significant bit is at the head of the linked list.
// Example 1:

// Input: head = [1,0,1]
// Output: 5
// Explanation: (101) in base 2 = (5) in base 10
// Example 2:

// Input: head = [0]
// Output: 0

func getDecimalValue(head *Node) int {
	power := 0
	head = reverse(head)
	num := 0
	for head != nil {
		num += head.val * int(math.Pow(float64(2), float64(power)))
		power += 1
		head = head.next
	}
	return num

}

// 15
func findMiddle(head *Node) *Node {
	fast := head
	slow := head
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	return slow
}

// 16
// https://leetcode.com/problems/palindrome-linked-list/
// Given the head of a singly linked list, return true if it is a palindrome or false otherwise.
// Example 1:
// Input: head = [1,2,2,1]
// Output: true
// Example 2:
// Input: head = [1,2]
// Output: false
func isPalindrome(head *Node) bool {
	mid := findMiddle(head)
	mid = reverse(mid)
	for mid != nil {
		if head.val != mid.val {
			return false
		}
		head = head.next
		mid = mid.next
	}
	return true

}

func FindMiddle1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := new(ListNode)
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		prev = slow
		fast = fast.Next.Next
		slow = slow.Next
	}
	if prev.Next != nil {
		prev.Next = nil
	}
	return slow

}
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	mid := FindMiddle1(head)

	head1 := new(TreeNode)
	head1.Val = mid.Val

	if mid == head {
		return head1
	}

	head1.Left = sortedListToBST(head)
	head1.Right = sortedListToBST(mid.Next)
	return head1

}

func main() {
	l := new(SLL)
	l.Create()
	l.Print()
}
