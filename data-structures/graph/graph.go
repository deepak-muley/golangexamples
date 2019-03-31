package graph

import (
	"fmt"
	"sync"

	"github.com/deepak-muley/golangexamples/data-structures/queue"
)

type GraphNode struct {
	value   string
	visited bool
}

func NewGraphNode(value string) *GraphNode {
	gn := GraphNode{value, false}
	return &gn
}

func (graphNode GraphNode) String() string {
	return fmt.Sprintf("%v", graphNode.value)
}

func (graphNode GraphNode) Print() {
	fmt.Printf("|%v| ", graphNode.String())
}

type Graph struct {
	//vertices are array/slice of graph nodes
	vertices []*GraphNode
	//edges are map of a graph node to multiple graph nodes
	edges  map[GraphNode][]*GraphNode
	rwLock sync.RWMutex
}

func (graph *Graph) AddNode(graphNode *GraphNode) {
	//Take exclusive lock
	graph.rwLock.Lock()
	defer graph.rwLock.Unlock()

	graph.vertices = append(graph.vertices, graphNode)

}

func (graph *Graph) AddEdge(directedEdge bool, graphNode1, graphNode2 *GraphNode) {
	//Take exclusive lock
	graph.rwLock.Lock()
	defer graph.rwLock.Unlock()

	if graph.edges == nil {
		graph.edges = make(map[GraphNode][]*GraphNode)
	}

	if directedEdge {
		graph.edges[*graphNode1] = append(graph.edges[*graphNode1], graphNode2)
	} else {
		graph.edges[*graphNode1] = append(graph.edges[*graphNode1], graphNode2)
		graph.edges[*graphNode2] = append(graph.edges[*graphNode2], graphNode1)
	}
}

func (graph *Graph) Print() {
	//Take Read lock
	graph.rwLock.RLock()
	defer graph.rwLock.RUnlock()

	s := ""
	for i := 0; i < len(graph.vertices); i++ {
		s += graph.vertices[i].String() + " -> "
		connectedEdges := graph.edges[*graph.vertices[i]]
		for j := 0; j < len(connectedEdges); j++ {
			s += connectedEdges[j].String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
}

func (graph *Graph) PrintDFS() {
	gn := graph.vertices[0]
	gn.Print()
	graph.traverseDFS(gn)
}

func (graph *Graph) traverseDFS(gn *GraphNode) {
	edges := graph.edges[*gn]
	for i := 0; i < len(edges); i++ {
		edges[i].Print()
		graph.traverseDFS(edges[i])
	}
}

func (graph *Graph) PrintBFS() {
	gn := graph.vertices[0]
	var q *queue.Queue = new(queue.Queue)
	q.EnqueueGeneric(gn)
	graph.traverseBFS(q)
}

func (graph *Graph) traverseBFS(q *queue.Queue) (err error) {
	for {
		if q.IsEmpty() {
			return nil
		}
		node, err := q.DequeueGeneric()
		if err != nil {
			return err
		}
		gn, ok := node.(*GraphNode)
		if !ok {
			return fmt.Errorf("dequeued node is not of type *GraphNode")
		}
		gn.Print()
		edges := graph.edges[*gn]
		for i := 0; i < len(edges); i++ {
			q.EnqueueGeneric(edges[i])
		}
	}
}
