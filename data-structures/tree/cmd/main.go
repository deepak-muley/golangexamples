package main

import (
	"fmt"

	"github.com/deepak-muley/golangexamples/data-structures/tree"
)

func main() {
	// (3 + 4) * 6
	//      *
	//  +      6
	//3   4
	root := new(tree.TreeNode)
	root.SetValue("*")
	root.AddRightChild("6")
	plusNode := root.AddLeftChild("+")
	plusNode.AddLeftChild("3")
	plusNode.AddRightChild("4")

	fmt.Println("Printing Tree")
	root.PrintNode()
	plusNode.PrintNode()

	fmt.Println("Printing Inorder")
	tree.PrintInorder(root)
	fmt.Println()

	fmt.Println("Printing Preorder")
	tree.PrintPreorder(root)
	fmt.Println()

	fmt.Println("Printing Postorder")
	tree.PrintPostorder(root)
	fmt.Println()
}
