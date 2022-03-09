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

//main function.
func main() {
	myTrees := createTree()
	values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	myTrees.createLevelOrderTree(values)
	myTrees.printLevelOrderTree()
}
