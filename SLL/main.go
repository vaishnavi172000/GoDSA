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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var result [][]int

func leafPathSum(root *TreeNode, targetSum int, path []int, currSum int) {
	fmt.Println("val targetSum path currsum result", root.Val, targetSum, path, currSum, result)
	if root.Left == nil && root.Right == nil {
		sum := currSum + root.Val
		if sum == targetSum {
			path = append(path, root.Val)
			result = append(result, path)
		}
		return
	}
	path = append(path, root.Val)
	currSum += root.Val
	if root.Left != nil {
		leafPathSum(root.Left, targetSum, path, currSum)
	}
	if root.Right != nil {
		leafPathSum(root.Right, targetSum, path, currSum)
	}
	fmt.Println("%%%%%%val targetSum path currsum result", root.Val, targetSum, path, currSum, result)
}
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return result
	}
	var path []int
	leafPathSum(root, targetSum, path, 0)
	return result

}

func main() {
	l := new(SLL)
	l.Create()
	l.Print()
}
