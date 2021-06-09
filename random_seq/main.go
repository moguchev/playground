package main

import (
	"fmt"
)

func main() {
	fmt.Println(random(5))
}

func random(n int) []int {
	// rand.Seed(time.Now().UnixNano())

	// array := make([]int, 0, n)

	// // m := make(map[int]bool, n)
	// k := make(map[int]int, n)

	// for i := 0; i < n; {
	// 	v := rand.Int()

	// 	_, exist := k[v]
	// 	// exist = false
	// 	// value = default (0, "", false, nil, nil, {})
	// 	if !exist {
	// 		array = append(array, v)
	// 		k[v] = v
	// 		i++
	// 	}

	// 	// if !m[v] {
	// 	// 	array = append(array, v)
	// 	// 	m[v] = true
	// 	// 	i++
	// 	// }
	// }

	// for index, value := range array {
	// 	fmt.Printf("%d: %d(%d)", index, array[index], value)
	// }

	// for key, value := range k {
	// 	fmt.Printf("%d: %d %s %v %x", key, value, "", []int{1, 3, 34, 45}, 1, )
	// }

	str := "ABCDabcd1234абвг" /// == []byte

	fmt.Printf("len: %d\n", len(str)) // байты

	for i, ch := range str {
		fmt.Printf("%d: %c\n", i, ch)
	}

	return nil
}
