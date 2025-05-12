package main

import (
	"fmt"
	"math"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

// 1
// getNode function to create a new node
func getNode(value int) *Node {
	return &Node{data: value}
}

// 2
// insert function to insert a new node
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

// 3
// inorder traversal recursive
func inorder(root *Node) {
	if root == nil {
		return
	}
	inorder(root.left)
	fmt.Printf("%d ", root.data)
	inorder(root.right)
}

// 4
// postorder traversal recursive
func postOrder(root *Node) {
	if root == nil {
		return
	}
	postOrder(root.left)
	postOrder(root.right)
	fmt.Printf("%d ", root.data)
}

// 5
// preorder traversal recursive
func preOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.data)
	preOrder(root.left)
	preOrder(root.right)
}

// 6
// Search for a key in BST
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

// 7
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

// 8
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
// Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.
// According to the definition of LCA on Wikipedia:
// “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as
// descendants (where we allow a node to be a descendant of itself).”

func lowestCommonAncestor(root, p, q *Node) *Node {
	if root == nil {
		return nil
	}

	curr := root.data
	if curr < p.data && curr < q.data {
		return lowestCommonAncestor(root.right, p, q)
	}
	if curr > p.data && curr > q.data {
		return lowestCommonAncestor(root.left, p, q)
	}
	return root

}

var maxI int

// 9
// https://leetcode.com/problems/binary-tree-maximum-path-sum/
func maxSum(root *Node) int {
	maxI = max(root.data, maxI)
	if root.left == nil && root.right == nil {

		return root.data

	}
	//fmt.Println("%%%%%%")
	lh := 0
	rh := 0
	if root.left != nil {
		lh = maxSum(root.left)
		maxI = max(maxI, lh)
	}
	if root.right != nil {
		rh = maxSum(root.right)
		maxI = max(maxI, rh)
	}
	maxI = max(maxI, root.data+lh+rh)
	return root.data + max(lh, rh)

}
func maxPathSum(root *Node) int {
	maxI = math.MinInt
	if root == nil {
		return 0
	}
	val := maxSum(root)
	//fmt.Println("val max", val, maxI)

	if val > maxI {
		//fmt.Println("dbfhjdfewh")
		return val
	}

	return maxI

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
