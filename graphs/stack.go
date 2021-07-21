package main

type Stack []interface{}

func (s *Stack) Push(v interface{}) {
	*s = append(*s, v)
}

func (s *Stack) Pop() interface{} {
	h := *s
	var el interface{}
	el, *s = h[len(h)-1], h[0:len(h)-1]
	return el
}

func (s *Stack) Len() int {
	return len(*s)
}

func NewStack() *Stack {
	return &Stack{}
}
