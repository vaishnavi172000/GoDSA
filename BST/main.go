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
