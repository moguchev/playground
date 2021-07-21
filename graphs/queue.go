package main

type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	h := *q
	var el interface{}
	el, *q = h[0], h[1:]
	return el
}

func (q *Queue) Len() int {
	return len(*q)
}

func NewQueue() *Queue {
	return &Queue{}
}
