package main

import (
	"errors"
	"fmt"
)

//node Structure.
type node struct {
	data int
	next *node
}

//linkedList is the head of the linked list to be created.
type linkedList struct {
	head *node
}

//createlinkedList is used to create a new linked list.
func createlinkedList() *linkedList {
	return &linkedList{}
}

//insertIntoLinkedList will insert a new node at the end of the linked list.
func (myList *linkedList) insertIntoLinkedList(value int) {
	newNode := node{}
	newNode.data = value
	if myList.head == nil {
		myList.head = &newNode
		return
	}
	temp := myList.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = &newNode
}

//printLinkedList is used to print the linked list.
func (myList *linkedList) printLinkedList() error {
	currentNode := myList.head
	if currentNode == nil {
		return nil
	} else {
		for currentNode != nil {
			fmt.Println(currentNode.data)
			currentNode = currentNode.next
		}
	}
	return nil
}

//insertAt inserts a new node at the given position.
// Consider the position of first node starts from 1.
func (myList *linkedList) insertAt(position, value int) {
	newNode := node{data: value}
	tempCurrNode := myList.getNodeAt(position)
	tempPrevNode := myList.getNodeAt(position - 1)
	tempPrevNode.next = &newNode
	newNode.next = tempCurrNode
}

//getNodeAt gives you the node at requested position.
func (myList *linkedList) getNodeAt(position int) *node {
	temp := myList.head
	if position <= 1 {
		return temp
	}
	for i := 0; i < position-1; i++ {
		temp = temp.next
	}
	return temp
}

//deleteNodeAt is used to delete a node at the given position.
func (myList *linkedList) deleteNodeAt(position int) error {
	if position < 1 {
		return errors.New("Position cannot be less than 1")
	}
	if position == 1 {
		tempNextNode := myList.getNodeAt(position + 1)
		myList.head = tempNextNode
		return nil
	}
	tempNextNode := myList.getNodeAt(position + 1)
	tempPrevNode := myList.getNodeAt(position - 1)
	tempPrevNode.next = tempNextNode
	return nil
}

//main function.
func main() {
	linkedList := createlinkedList()
	linkedList.insertIntoLinkedList(10)
	linkedList.insertIntoLinkedList(15)
	linkedList.insertIntoLinkedList(20)
	linkedList.insertAt(2, 30)
	linkedList.printLinkedList()
	err := linkedList.deleteNodeAt(2)
	if err != nil {
		fmt.Println(err)
	}
	linkedList.printLinkedList()
}
