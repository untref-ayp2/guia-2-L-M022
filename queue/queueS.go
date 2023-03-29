package queue

import (
	"errors"
	"guia2/stack"
)

type QueueS struct {
	a stack.Stack
	b stack.Stack
}

func (q *QueueS) Enqueue(v any) {
	q.a.Push(v)
}

func (q *QueueS) Dequeue() (any, error) {
	if q.IsEmpty() {

		return nil, errors.New("la cola esta vacia")
	}
	if q.b.IsEmpty() {
		for !q.a.IsEmpty() {
			aux, _ := q.a.Pop()
			q.b.Push(aux)
		}
	}
	return q.b.Pop()
}

func (q *QueueS) IsEmpty() bool {
	return q.a.IsEmpty() && q.b.IsEmpty()
}

func (q *QueueS) Front() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("la cola esta vacia")
	}
	if q.b.IsEmpty() {
		for !q.a.IsEmpty() {
			aux, _ := q.a.Pop()
			q.b.Push(aux)
		}
	}
	return q.b.Top()
}
