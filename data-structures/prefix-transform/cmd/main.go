package main

import (
	"fmt"

	ds "github.com/deepak-muley/golangexamples/prefix-transform"
)

/*
Input
{
  "1": "ABC",
  "2": "AB",
  "3", "B",
  "4", "ABCD"
}

* Transform into the below
Output:
{
  "1": "2C",
  "2": "AB",
  "3": "B",
  "4": "1D"
}
*/

func createGraph(input map[string][]string) *ds.Graph {
	gr := ds.NewGraph()

	//Create the graph with given input
	for tag, vals := range input {
		var curNode, curRoot *ds.GraphNode

		var v string
		if len(vals) > 0 {
			v = vals[0]
		} else {
			continue
		}
		var gnode *ds.GraphNode
		gnode = gr.GetRootNode(v)
		if gnode == nil {
			gnode = gr.AddRootNode(v)
		}
		curRoot = gnode
		curNode = curRoot

		for _, val := range vals {
			childNode := curNode.GetChild(val)
			if childNode == nil {
				childNode = curNode.AddChild(val)
			}
			curNode = childNode
		}
		curNode.SetTag(tag)
	}

	return gr
}

func createOutputMap(gr *ds.Graph, input map[string][]string) map[string][]string {
	output := make(map[string][]string)
	var curNode, curRoot *ds.GraphNode

	for tag, vals := range input {

		var v string
		if len(vals) > 0 {
			v = vals[0]
		} else {
			continue
		}

		var gnode *ds.GraphNode
		gnode = gr.GetRootNode(v)
		if gnode == nil {
			gnode = gr.AddRootNode(v)
		}
		curRoot = gnode
		curNode = curRoot

		var curTag, prevTag string
		var outputVals []string
		var history []string
		for _, val := range vals {
			childNode := curNode.GetChild(val)
			if childNode == nil {
				outputVals = append(outputVals, val)
				continue
			}
			if childNode.GetTag() != "" {
				curTag = childNode.GetTag()
			} else {
				history = append(history, val)
			}

			//end of vals == input tag
			if curTag == tag {
				if prevTag != "" {
					outputVals = append(outputVals, prevTag)
					outputVals = append(outputVals, val)
				} else {
					outputVals = append(history, val)
				}
			} else {
				curTag = childNode.GetTag()
				prevTag = childNode.GetTag()
			}

			curNode = childNode
		}
		output[tag] = outputVals
	}

	return output
}

func main() {
	input := map[string][]string{
		"1": []string{"A", "B", "C"},
		"2": []string{"A", "B"},
		"3": []string{"B"},
		"4": []string{"A", "B", "C", "D"},
	}
	fmt.Println("Input:")
	for tag, vals := range input {
		fmt.Println(tag, vals)
	}

	//Create graph from input map
	gr := createGraph(input)
	//print the entire graph
	gr.Print()

	//Create Output from graph
	output := createOutputMap(gr, input)
	fmt.Println("Output:")
	for tag, vals := range output {
		fmt.Println(tag, vals)
	}

}
