package main

import (
	"fmt"
)

func main() {
	i := []int{1, 2, 3, 4, 5}
	j := i[0 : len(i)-1]
	fmt.Println(j)
}
