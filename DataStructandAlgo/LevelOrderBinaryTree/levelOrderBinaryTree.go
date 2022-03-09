package main

import (
	"DataStructandAlgo/queue"
	"fmt"
)

//node of the tree contains data and pointer to left and right node.
type node struct {
	data        int
	left, right *node
}

//tree is a structure which has root of type node.
type tree struct {
	root *node
}

//createTree function is used to create a new tree.
func createTree() *tree {
	return &tree{}
}

//createLevelOrderTree is used to create a new tree in level order.
func (myTree *tree) createLevelOrderTree(arr []int) {
	myQueue := &queue.Queue{}
	for _, val := range arr {
		var newNode *node = &node{val, nil, nil}
		myQueue.Add(newNode)
		//fmt.Println(myQueue.Peek().(*node).data)
		if myTree.root == nil {
			myTree.root = newNode
		} else if myQueue.Peek().(*node).left == nil {
			temp := myQueue.Peek().(*node) //use type assertion.
			temp.left = newNode
		} else {
			temp := myQueue.Peek().(*node)
			temp.right = newNode
			myQueue.Remove()
		}
	}
}

//printLevelOrderTree function is used to print the Binary Tree level wise.
func (myTree *tree) printLevelOrderTree() {
	myPrintQueue := &queue.Queue{}
	myPrintQueue.Add(myTree.root)
	for !myPrintQueue.IsEmpty() {
		currentNode := myPrintQueue.Remove().(*node)
		fmt.Println(currentNode.data)
		if currentNode.left != nil {
			myPrintQueue.Add(currentNode.left)
		}
		if currentNode.right != nil {
			myPrintQueue.Add(currentNode.right)
		}
	}
}

//printPreOrderTraversal function is used to call the printPreOrderTraversal function which has the same name as current function name but different receiver type.
//In go we can have two functions with same name if they have different receiver types else go doesn't support function overloading or overwriting.
func (myTree *tree) printPreOrderTraversal() {
	printPreOrderTraversal(myTree.root)
	fmt.Println()
}

//printPreOrderTraversal function used to traverse and print the tree in pre-order fashion i.e root, left and then right node.
func printPreOrderTraversal(currentNode *node) {
	if currentNode == nil {
		return
	}
	fmt.Println(currentNode.data)
	printPreOrderTraversal(currentNode.left)
	printPreOrderTraversal(currentNode.right)
}

//printInOrderTraversal function is used to call the printInOrderTraversal function to traverse and print the tree in inorder fashion.
func (myTree *tree) printInorderTraversal() {
	printInorderTraversal(myTree.root)
	fmt.Println()
}

//printInOrderTraversal function used to traverse and print the tree in In-order fashion i.e left, root and then right node.
func printInorderTraversal(currentNode *node) {
	if currentNode == nil {
		return
	}
	printInorderTraversal(currentNode.left)
	fmt.Println(currentNode.data)
	printInorderTraversal(currentNode.right)
}

//printPostOrderTraversal function is used to call the printPostOrderTraversal function to traverse and print the tree in post-order fashion.
func (myTree *tree) printPostOrderTraversal() {
	printPostOrderTraversal(myTree.root)
	fmt.Println()
}

//printPostOrderTraversal function used to traverse and print the tree in post-order fashion i.e left, right and then root node.
func printPostOrderTraversal(currentNode *node) {
	if currentNode == nil {
		return
	}
	printPostOrderTraversal(currentNode.left)
	printPostOrderTraversal(currentNode.right)
	fmt.Println(currentNode.data)
}

//main function.
func main() {
	myTrees := createTree()
	values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	myTrees.createLevelOrderTree(values)
	fmt.Println("-------------------- LEVEL ORDER TRAVERSAL --------------------")
	myTrees.printLevelOrderTree()
	fmt.Println("-------------------- PRE-ORDER TRAVERSAL --------------------")
	myTrees.printPreOrderTraversal()
	fmt.Println("-------------------- IN-ORDER TRAVERSAL --------------------")
	myTrees.printInorderTraversal()
	fmt.Println("-------------------- POST-ORDER TRAVERSAL --------------------")
	myTrees.printPostOrderTraversal()
}
