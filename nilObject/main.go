package main

import "fmt"

type Object struct {
	Value int
}

func (o Object) foo(v int) bool {
	return o.bar(v)
}

func (this *Object) bar(v int) bool {
	if this == nil {
		return false
	}
	return this.Value > v
}

func main() {
	obj1 := Object{Value: 3}
	fmt.Println(obj1.foo(1))
	var obj2Ptr *Object
	fmt.Println(obj2Ptr.bar(1))
}
