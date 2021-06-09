package main

import (
	"fmt"
	"reflect"
)

func foo(out map[string]string, key string) error {
	out[key] = "foo"
	return nil
}

func bar(m map[int64]bool, s []int64) {
	for i := range s {
		m[s[i]] = true
	}
}

type Foo interface {
	Foo()
}

type A struct{}

func (a *A) Foo() {
	fmt.Println("A")
}

type B struct{}

func (a *B) Foo() {
	fmt.Println("B")
}

type C struct{}

func (a *C) Foo() {
	fmt.Println("C")
}

func main() {
	// m := make(map[string]string)
	// err := foo(m, "key")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// for k, v := range m {
	// 	fmt.Printf("%s:%s\n", k, v)
	// }

	// s := map[int64]bool{1: false, 2: false, 3: false, 4: false}
	// l := []int64{1, 2}

	// bar(s, l)
	// fmt.Println(s)

	a, b, c := new(A), new(B), new(C)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))

	foos := Foos{make(map[reflect.Type]Foo)}
	foos.Add(a)
	foos.Add(b)
	foos.Add(c)

	for _, f := range foos.m {
		f.Foo()
	}

	aa, _ := foos.Get(reflect.TypeOf(a))
	bb, _ := foos.Get(reflect.TypeOf(b))
	cc, _ := foos.Get(reflect.TypeOf(c))

	aa.Foo()
	bb.Foo()
	cc.Foo()
}

type Foos struct {
	m map[reflect.Type]Foo
}

func (f *Foos) Add(foo Foo) {
	fmt.Println(reflect.TypeOf(foo))
	f.m[reflect.TypeOf(foo)] = foo
}

func (f *Foos) Get(restriction reflect.Type) (Foo, bool) {
	foo, ok := f.m[restriction]
	return foo, ok
}
