package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	if n < 1 {
		return
	}

	cities := make([]Point, 0, n)

	var line string
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		scanner.Scan()
		line = scanner.Text()

		values := strings.SplitN(line, " ", 2)
		x, _ := strconv.Atoi(values[0])
		y, _ := strconv.Atoi(values[1])

		cities = append(cities, Point{
			X: x, Y: y,
		})
	}

	scanner.Scan()
	line = scanner.Text()
	distance, _ := strconv.Atoi(line)

	scanner.Scan()
	line = scanner.Text()

	values := strings.SplitN(line, " ", 2)
	start, _ := strconv.Atoi(values[0])
	end, _ := strconv.Atoi(values[1])

	g := NewListGraph(len(cities))

	for i := 0; i < len(cities); i++ {
		for j := i + 1; j < len(cities); j++ {
			if Distance(cities[i], cities[j]) <= distance {
				g.AddEdge(i+1, j+1)
			}
		}
	}

	if start == end {
		fmt.Println(0)
		return
	}

	count, stop := -1, true
	BFS(g, start, func(vertex, lvl int) bool {
		if vertex == end {
			count = lvl
			return stop
		}
		return !stop
	})

	fmt.Println(count)
}

func Distance(a, b Point) int {
	return Abs(a.X-b.X) + Abs(a.Y-b.Y)
}

func Abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
