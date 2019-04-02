package ds

import "fmt"

type Graph struct {
	nodes map[string]*GraphNode
}

func NewGraph() *Graph {
	gr := new(Graph)
	gr.nodes = make(map[string]*GraphNode)
	return gr
}

func (gr *Graph) AddRootNode(s string) *GraphNode {
	gnode := GraphNode{children: make(map[string]*GraphNode)}
	gr.nodes[s] = &gnode
	return &gnode
}

func (gr *Graph) GetRootNode(s string) *GraphNode {
	if gnode, ok := gr.nodes[s]; ok {
		return gnode
	}
	return nil
}

func (gr *Graph) Print() {
	fmt.Println("Printing current graph")
	for _, v := range gr.nodes {
		v.Print()
	}
}

type GraphNode struct {
	tag      string
	children map[string]*GraphNode
}

func (gn *GraphNode) SetTag(tag string) {
	gn.tag = tag
}

func (gn *GraphNode) GetTag() string {
	return gn.tag
}

func (gn *GraphNode) AddChild(val string) *GraphNode {
	gnode := GraphNode{children: make(map[string]*GraphNode)}
	gn.children[val] = &gnode
	return &gnode
}

func (gn *GraphNode) GetChild(val string) *GraphNode {
	if gnode, ok := gn.children[val]; ok {
		return gnode
	}
	return nil
}

func (gn *GraphNode) Print() {
	for k, v := range gn.children {
		fmt.Printf("%v (tag:%s) => ", k, v.tag)
		v.Print()
	}
	fmt.Println()
}
