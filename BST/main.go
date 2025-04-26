package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func getNode(value int) *Node {
	return &Node{data: value}
}

func insert(value int, root *Node) *Node {
	if root == nil {
		root = getNode(value)
		return root
	}

	if root.data < value {
		root.right = insert(value, root.right)
	} else {
		root.left = insert(value, root.left)

	}
	return root
}

func inorder(root *Node) {
	if root == nil {
		return
	}
	inorder(root.left)
	fmt.Printf("%d ", root.data)
	inorder(root.right)
}

func postOrder(root *Node) {
	if root == nil {
		return
	}
	postOrder(root.left)
	postOrder(root.right)
	fmt.Printf("%d ", root.data)
}

func preOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.data)
	preOrder(root.left)
	preOrder(root.right)
}

func Search(root *Node, key int) bool {
	if root == nil {
		return false
	}
	if root.data == key {
		return true
	}
	if root.data < key {
		return Search(root.right, key)
	}
	return Search(root.left, key)
}

// https://leetcode.com/problems/path-sum-ii/?envType=problem-list-v2&envId=binary-tree
// Given the root of a binary tree and an integer targetSum,
// return all root-to-leaf paths where the sum of the node values in the path equals targetSum.
// Each path should be returned as a list of the node values, not node references.
// A root-to-leaf path is a path starting from the root and ending at any leaf node. A leaf is a node with no children.
// Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
// Output: [[5,4,11,2],[5,8,4,5]]
// Explanation: There are two paths whose sum equals targetSum:
// 5 + 4 + 11 + 2 = 22
// 5 + 8 + 4 + 5 = 22
var result [][]int

func leafPathSum(root *Node, targetSum int, path []int, currSum int) {
	if root.left == nil && root.right == nil {
		sum := currSum + root.data
		if sum == targetSum {
			path = append(path, root.data)
			result = append(result, path)
		}
		return
	}
	var path1 []int
	path1 = append(path1, path...)
	path1 = append(path1, root.data)
	currSum += root.data
	if root.left != nil {
		leafPathSum(root.left, targetSum, path1, currSum)
	}
	if root.right != nil {
		leafPathSum(root.right, targetSum, path1, currSum)
	}
}
func pathSumtoLeafs(root *Node, targetSum int) [][]int {
	result = [][]int{}
	if root == nil {
		return result
	}
	var path []int
	leafPathSum(root, targetSum, path, 0)
	return result

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var Dummy, temp *ListNode
	Dummy = new(ListNode)
	temp = Dummy
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			temp.Next = list2
			temp = temp.Next
			list2 = list2.Next
		} else {
			temp.Next = list1
			temp = temp.Next
			list1 = list1.Next
		}
	}
	if list1 != nil {
		temp.Next = list1
	}
	if list2 != nil {
		temp.Next = list2
	}
	return Dummy.Next

}

func main() {
	var root *Node

	// Insert keys into BST
	keys := []int{50, 30, 20, 40, 70, 60, 80}
	for _, key := range keys {
		root = insert(key, root)
	}

	fmt.Print("Inorder traversal: ")
	inorder(root)
	fmt.Println()

	fmt.Println("Preorder traversal: ")
	preOrder(root)
	fmt.Println()

	fmt.Println("Postorder traversal: ")
	postOrder(root)
	fmt.Println()

	//Search for a key
	keyToSearch := 60
	if Search(root, keyToSearch) {
		fmt.Printf("Key %d found in BST\n", keyToSearch)
	} else {
		fmt.Printf("Key %d not found in BST\n", keyToSearch)
	}
}
