package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func send(msg *string) {
	time.Sleep(10 * time.Second)
	fmt.Println(*msg)
}

func sender(w http.ResponseWriter, r *http.Request) {
	msg := "HELLO"
	fmt.Println("send")
	go send(&msg)
	fmt.Println("Done")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Done"))
}

type Node struct {
	ID       string
	Fot      int
	Sum      int
	Children []Node
}

var tree = Node{
	ID:  "0",
	Fot: 1,
	Children: []Node{
		Node{
			ID:  "0.1",
			Fot: 1,
			Children: []Node{
				Node{
					ID:  "0.1.1",
					Fot: 1,
					Children: []Node{
						Node{
							ID:  "0.1.1.1",
							Fot: 1,
							Children: []Node{
								Node{
									ID:  "0.1.1.1.1",
									Fot: 1,
									Children: []Node{
										Node{
											ID:  "0.1.1.1.1.1",
											Fot: 1,
										},
									},
								},
							},
						},
					},
				},
				Node{
					ID:  "0.1.2",
					Fot: 1,
					Children: []Node{
						Node{
							ID:  "0.1.2.1",
							Fot: 1,
						},
					},
				},
			},
		},
		Node{
			ID:  "0.2",
			Fot: 1,
			Children: []Node{
				Node{
					ID:  "0.2.1",
					Fot: 1,
					Children: []Node{
						Node{
							ID:  "0.2.1.1",
							Fot: 1,
						},
					},
				},
				Node{
					ID:  "0.2.2",
					Fot: 1,
					Children: []Node{
						Node{
							ID:  "0.2.2.1",
							Fot: 1,
						},
					},
				},
				Node{
					ID:  "0.2.3",
					Fot: 1,
					Children: []Node{
						Node{
							ID:  "0.2.3.1",
							Fot: 1,
						},
					},
				},
				Node{
					ID:  "0.2.4",
					Fot: 1,
					Children: []Node{
						Node{
							ID:  "0.2.4.1",
							Fot: 1,
						},
					},
				},
				Node{
					ID:  "0.2.5",
					Fot: 1,
					Children: []Node{
						Node{
							ID:  "0.2.5.1",
							Fot: 1,
						},
					},
				},
			},
		},
		Node{
			ID:  "0.3",
			Fot: 1,
		},
	},
}

func recurciveRun(root *Node) {
	if root == nil {
		return
	}

	for i := range root.Children {
		recurciveRun(&root.Children[i])
	}

	for i := range root.Children {
		root.Sum += root.Children[i].Sum
	}
	root.Sum += root.Fot
	time.Sleep(time.Second / 10) // imitation hard calculation
}

func recurciveRunParallel(root *Node, lvl int, wg *sync.WaitGroup) {
	defer wg.Done()
	if root == nil {
		return
	}
	// fmt.Println(root.ID, ":", lvl)

	var childGroup sync.WaitGroup
	childGroup.Add(len(root.Children))
	for i := range root.Children {
		go recurciveRunParallel(&root.Children[i], lvl+1, &childGroup)
	}

	for i := range root.Children {
		root.Sum += root.Children[i].Sum
	}
	root.Sum += root.Fot
	time.Sleep(time.Second / 10) // imitation hard calculation
}

func recurciveRunHalfParallel(root *Node, lvl int, wg *sync.WaitGroup) {
	defer wg.Done()
	if root == nil {
		return
	}

	var childGroup sync.WaitGroup
	childGroup.Add(len(root.Children))
	for i := range root.Children {
		if lvl < 3 {
			go recurciveRunHalfParallel(&root.Children[i], lvl+1, &childGroup)
		} else {
			recurciveRunHalfParallel(&root.Children[i], lvl+1, &childGroup)
		}
	}
	childGroup.Wait()

	for i := range root.Children {
		root.Sum += root.Children[i].Sum
	}
	root.Sum += root.Fot
	time.Sleep(time.Second / 10) // imitation hard calculation
}

func main() {
	start := time.Now()
	recurciveRun(&tree)
	fmt.Println(time.Now().Sub(start))
	fmt.Println(tree.Sum)

	var wg sync.WaitGroup
	wg.Add(1)
	start = time.Now()
	recurciveRunParallel(&tree, 1, &wg)
	fmt.Println(time.Now().Sub(start))
	wg.Wait()

	var wg1 sync.WaitGroup
	wg1.Add(1)
	start = time.Now()
	recurciveRunHalfParallel(&tree, 1, &wg1)
	fmt.Println(time.Now().Sub(start))
	wg1.Wait()
}
