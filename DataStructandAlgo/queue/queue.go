package queue

type Queue struct {
	queue []interface{}
}

func (myQueue *Queue) Add(value interface{}) {
	myQueue.queue = append(myQueue.queue, value)
}

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

func (myQueue *Queue) Peek() interface{} {
	if myQueue.IsEmpty() {
		return nil
	}
	return myQueue.queue[0]
}

func (myQueue *Queue) IsEmpty() bool {
	return len(myQueue.queue) == 0
}

func (myQueue *Queue) Size() int {
	return len(myQueue.queue)
}
