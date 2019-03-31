package graph

import (
	"fmt"
	"testing"

	"github.com/deepak-muley/golangexamples/data-structures/graph"
)

func Test_graph_DFS_BFS(t *testing.T) {
	var g graph.Graph

	gnA := graph.NewGraphNode("A")
	gnB := graph.NewGraphNode("B")
	gnC := graph.NewGraphNode("C")
	gnD := graph.NewGraphNode("D")
	gnE := graph.NewGraphNode("E")
	gnF := graph.NewGraphNode("F")
	g.AddNode(gnA)
	g.AddNode(gnB)
	g.AddNode(gnC)
	g.AddNode(gnD)
	g.AddNode(gnE)
	g.AddNode(gnF)

	g.AddEdge(true, gnA, gnB)
	g.AddEdge(true, gnA, gnC)
	g.AddEdge(true, gnA, gnD)
	g.AddEdge(true, gnD, gnB)
	g.AddEdge(true, gnC, gnE)
	g.AddEdge(true, gnB, gnF)
	g.Print()
	fmt.Println("\n----- Print DFS -----")
	g.PrintDFS()
	fmt.Println("\n----- Print BFS -----")
	g.PrintBFS()
	fmt.Println("")
}
