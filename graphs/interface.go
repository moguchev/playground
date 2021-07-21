package main

type Graph interface {
	AddEdge(from, to int)
	VerticesCount() int

	GetNextVertices(vertex int) []int
	// GetPrevVertices(vertex int) []int
}

type ListGraph struct {
	adjacencylists map[int][]int
}

func NewListGraph(size int) *ListGraph {
	return &ListGraph{
		adjacencylists: make(map[int][]int, size),
	}
}

func (lg *ListGraph) AddEdge(from, to int) {
	lg.adjacencylists[from] = append(lg.adjacencylists[from], to)
}

func (lg *ListGraph) VerticesCount() int {
	return len(lg.adjacencylists)
}

func (lg *ListGraph) GetNextVertices(vertex int) []int {
	return lg.adjacencylists[vertex]
}

func BFS(g Graph, vertex int, visit func(vertex, lvl int) bool) {
	visited := make(map[int]bool, g.VerticesCount())
	nextVertices := NewQueue()

	level := make(map[int]int)

	visited[vertex] = true
	nextVertices.Push(vertex)
	level[vertex] = 0

	for nextVertices.Len() > 0 {
		curr, _ := nextVertices.Pop().(int)
		lvl := level[curr]

		if stop := visit(curr, lvl); stop {
			break
		}

		for _, nextVertex := range g.GetNextVertices(curr) {
			if !visited[nextVertex] {
				visited[nextVertex] = true
				nextVertices.Push(nextVertex)
				if _, ok := level[nextVertex]; !ok {
					level[nextVertex] = lvl + 1
				}
			}
		}
	}
}

func DFS(g Graph, vertex int, visit func(int)) {
	visited := make([]bool, g.VerticesCount())
	nextVertices := NewStack()

	nextVertices.Push(vertex)

	visited[vertex] = true

	for nextVertices.Len() > 0 {
		curr, _ := nextVertices.Pop().(int)

		visit(curr)

		for _, nextVertex := range g.GetNextVertices(curr) {
			if !visited[nextVertex] {
				visited[nextVertex] = true
				nextVertices.Push(nextVertex)
			}
		}
	}
}
