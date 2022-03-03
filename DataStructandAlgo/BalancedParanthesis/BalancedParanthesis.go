package main

import (
	"DataStructandAlgo/stack"
	"fmt"
)

//isBalancedParanthesis is used to check if the expression has balanced symbols or not.
func isBalancedParanthesis(exp string) bool {
	myStack := stack.Stack{}
	for _, ch := range exp {
		switch ch {
		case '{', '[', '(':
			myStack.Push(ch)
		case '}':
			val := myStack.Pop()
			if val != '{' {
				return false
			}
		case ']':
			val := myStack.Pop()
			if val != '[' {
				return false
			}
		case ')':
			val := myStack.Pop()
			if val != '(' {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isBalancedParanthesis("{() ({})}"))
	fmt.Println(isBalancedParanthesis("{((}"))
}
