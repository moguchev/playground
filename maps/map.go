package main

import (
	"fmt"
)

func foo(out map[string]string, key string) error {
	out[key] = "foo"
	return nil
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
}
