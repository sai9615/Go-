package main

import (
	"errors"
	"fmt"
)

type node struct {
	prev *node
	data int
	next *node
}

type linkedList struct {
	head *node
	tail *node
}

func (myList *linkedList) insertAtFront(value int) {
	newNode := &node{data: value}
	if myList.head == nil {
		myList.head = newNode
		myList.tail = newNode
	} else {
		myList.head.prev = newNode
		newNode.next = myList.head
		myList.head = newNode
	}
}

func (myList *linkedList) insertAtTail(value int) {
	newNode := &node{data: value}
	if myList.tail == nil {
		myList.tail = newNode
		myList.head = newNode
	} else {
		myList.tail.next = newNode
		newNode.prev = myList.tail
		myList.tail = newNode
	}
}

func (myList *linkedList) printLinkedList() {
	temp := myList.head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
	fmt.Println()
}

func (myList *linkedList) reverseLinkedList() {
	currentNode := myList.head
	tempNode := &node{}
	for currentNode != nil {
		tempNode = currentNode.next
		currentNode.next = currentNode.prev
		currentNode.prev = tempNode
		if currentNode.prev == nil {
			myList.tail = myList.head
			myList.head = currentNode
			return
		}
		currentNode = currentNode.prev
	}
}

func (myList *linkedList) deleteNodeAt(position int) {
	nodeAt, err := myList.getNodeAt(position)
	if err != nil {
		fmt.Print(err)
		return
	} else {
		//If first node
		if position == 1 {
			myList.head = nodeAt.next
			nodeAt.next = nil
		} else if nodeAt.next == nil {
			//If last node
			myList.tail = nodeAt.prev
			nodeAt.prev.next = nil
		} else {
			nodeAt.prev.next = nodeAt.next
			nodeAt.next.prev = nodeAt.prev
		}
	}
}

func (myList *linkedList) sortedInsert(value int) {
	newNode := &node{data: value}
	temp := myList.head
	if temp == nil {
		myList.head = newNode
		myList.tail = newNode
		return
	} else if temp.data >= value {
		newNode.next = temp
		temp.prev = newNode
		temp = newNode
		return
	}
	for temp.next != nil && temp.data >= value {
		temp = temp.next
	}
	if temp.next == nil {
		temp.next = newNode
		newNode.prev = temp
		myList.tail = newNode
	} else {
		newNode.next = temp.next
		newNode.prev = temp
		temp.next = newNode
		newNode.next.prev = newNode
	}

}

func (myList *linkedList) getNodeAt(position int) (*node, error) {
	if position < 1 {
		return nil, errors.New("Please enter a position greater than 1")
	}
	temp := myList.head
	for i := 0; i < position; i++ {
		temp = temp.next
	}
	return temp, nil
}

func main() {
	list := &linkedList{}
	list.insertAtTail(5)
	list.insertAtTail(10)
	list.insertAtTail(15)
	list.insertAtTail(30)
	list.insertAtFront(1)
	list.printLinkedList()
	println()
	list.deleteNodeAt(2)
	list.printLinkedList()
	println()
	list.sortedInsert(6)
	list.printLinkedList()
	println()
	list.reverseLinkedList()
	list.printLinkedList()
	println()
}
