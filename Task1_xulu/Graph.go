package main

import "fmt"

type Graph struct {
	vertices int
	adjList  map[*Node][]*Node
}

func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		adjList:  make(map[*Node][]*Node),
	}
}

func (g *Graph) AddEdge(source, dest *Node) {
	g.adjList[source] = append(g.adjList[source], dest)
	g.adjList[dest] = append(g.adjList[dest], source)
}

func (g *Graph) DFSUtil(vertex *Node, visited map[*Node]bool) {
	visited[vertex] = true
	fmt.Printf("%d ", vertex.Result)

	for _, v := range g.adjList[vertex] {
		if !visited[v] {
			g.DFSUtil(v, visited)
		}
	}
}

func (g *Graph) DFS(startVertex *Node) {
	visited := make(map[*Node]bool)
	g.DFSUtil(startVertex, visited)
}
