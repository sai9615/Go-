package main

import (
	"DataStructandAlgo/queue"
	"errors"
	"fmt"
	"reflect"
)

func push(myQueue *queue.Queue, value interface{}) {
	myQueue.Add(value)
}

func pop(myQueue *queue.Queue) (error, interface{}) {
	if myQueue.IsEmpty() {
		return errors.New("Stack is empty"), nil
	} else {
		for i := 0; i < myQueue.Size()-1; i++ {
			val := myQueue.Remove()
			myQueue.Add(val)
		}
		value := myQueue.Remove()
		return nil, value
	}
}

func printStack(myQueue *queue.Queue) {
	arr := []interface{}{}
	for i := myQueue.Size() - 1; i >= 0; i-- {
		value := myQueue.Remove()
		arr = append(arr, value)
		myQueue.Add(value)
	}
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Println(arr[i])
	}
}

func main() {
	myQueue := &queue.Queue{}
	push(myQueue, 1)
	push(myQueue, 2)
	push(myQueue, 3)
	push(myQueue, 5)
	printStack(myQueue)
	err, val := pop(myQueue)
	vals := reflect.ValueOf(val).Type().String()
	fmt.Println()
	if err != nil {
		fmt.Println(err)
	} else {
		switch vals {
		case "int":
			println(val.(int))
		case "float64":
			println(val.(float64))
		case "float32":
			println(val.(float32))
		case "string":
			println(val.(string))
		case "bool":
			println(val.(bool))
		}
	}
	fmt.Println()
	printStack(myQueue)
}
