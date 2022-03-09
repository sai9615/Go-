package queue

//Queue Structure Defined using slice and interface{} to support generic data types.
type Queue struct {
	queue []interface{}
}

//Add function is used to add the value to the queue.
func (myQueue *Queue) Add(value interface{}) {
	myQueue.queue = append(myQueue.queue, value)
}

//Remove function is used to remove the front element from queue and return it.
func (myQueue *Queue) Remove() interface{} {
	if myQueue.IsEmpty() {
		return nil
	}
	value := myQueue.queue[0]
	if len(myQueue.queue) == 1 {
		//empty the slice
		myQueue.queue = myQueue.queue[:0]
	} else {
		myQueue.queue = myQueue.queue[1:len(myQueue.queue)]
	}
	return value
}

//Peek function is used to get the element at front of the queue.
func (myQueue *Queue) Peek() interface{} {
	if myQueue.IsEmpty() {
		return nil
	}
	return myQueue.queue[0]
}

//IsEmpty function is used to check if the queue is empty.
//Returns boolean value.
func (myQueue *Queue) IsEmpty() bool {
	return len(myQueue.queue) == 0
}

//Size function is used to return the size of the queue which is integer value.
func (myQueue *Queue) Size() int {
	return len(myQueue.queue)
}
