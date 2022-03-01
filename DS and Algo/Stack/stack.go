package stack

//Stack is a struct containing a generic slice declared using interface.
type Stack struct {
	s []interface{}
}

//Push the value to top of the stack.
func (stk *Stack) Push(value interface{}) {
	stk.s = append(stk.s, value)
}

//Pop removes and returns the top element of the stack.
func (stk *Stack) Pop() interface{} {
	if stk.IsEmpty() {
		return nil
	}
	val := stk.s[len(stk.s)-1]
	stk.s = stk.s[:len(stk.s)-1]
	return val
}

//Peek returns the top element of the stack.
func (stk *Stack) peek() interface{} {
	if stk.IsEmpty() {
		return nil
	}
	return stk.s[len(stk.s)-1]
}

//IsEmpty is used to check if stack is empty.
//Returns true if stack is empty else false.
func (stk *Stack) IsEmpty() bool {
	if len(stk.s) == 0 {
		return true
	}
	return false
}
