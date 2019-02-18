package tree

import "fmt"

type TreeNode struct {
	value string
	left  *TreeNode
	right *TreeNode
}

func (tn TreeNode) PrintNodeValue() {
	fmt.Print(" ", tn.value)
}

func (tn TreeNode) PrintNode() {
	fmt.Print("Value: ", tn.value)
	if tn.left != nil {
		fmt.Print(" Left: ", tn.left.value)
	}
	if tn.right != nil {
		fmt.Print(" Right: ", tn.right.value)
	}
	fmt.Println()
}

func (tn *TreeNode) SetValue(value string) {
	tn.value = value
}

func (tn *TreeNode) AddRightChild(value string) *TreeNode {
	childNode := new(TreeNode)
	childNode.SetValue(value)
	tn.right = childNode
	return childNode
}

func (tn *TreeNode) AddLeftChild(value string) *TreeNode {
	childNode := new(TreeNode)
	childNode.SetValue(value)
	tn.left = childNode
	return childNode
}

func PrintInorder(tn *TreeNode) {
	if tn == nil {
		return
	}
	PrintInorder(tn.left)
	tn.PrintNodeValue()
	PrintInorder(tn.right)
}

func PrintPreorder(tn *TreeNode) {
	if tn == nil {
		return
	}
	tn.PrintNodeValue()
	PrintInorder(tn.left)
	PrintInorder(tn.right)
}

func PrintPostorder(tn *TreeNode) {
	if tn == nil {
		return
	}

	PrintInorder(tn.left)
	PrintInorder(tn.right)
	tn.PrintNodeValue()
}
