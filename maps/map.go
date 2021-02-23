package main

import (
	"fmt"
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

func main() {
	m := make(map[string]string)
	err := foo(m, "key")
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range m {
		fmt.Printf("%s:%s\n", k, v)
	}

	s := map[int64]bool{1: false, 2: false, 3: false, 4: false}
	l := []int64{1, 2}

	bar(s, l)
	fmt.Println(s)
}
