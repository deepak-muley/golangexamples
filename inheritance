package main

import (
    "fmt"
)

// An interface
type FI interface {
    F()
}

// A parent with a default implementation of FI
type parent struct {
	FI
}
func (s *parent) F() {
	fmt.Printf("parent.F()\n");
}
func (s *parent) doit() {
	s.FI.F()
}
func NewParent() (rv * parent) {
	rv = new(parent)
	rv.FI = rv
	return
}

// A child1 with a different FI 
type child1 struct {
	parent // Anonymous parent struct to get default implementations
}
func (s *child1) F() {
	fmt.Printf("child1.F()\n");
}
func NewChild1() * child1 {
	c := new(child1)
	c.FI = c
	return c
}

// A child2 with default FI
type child2 struct {
	parent // Anonymous parent struct to get default implementations
	j int
}
func NewChild2() * child2 {
	c := new(child2)
	c.FI = c
	c.j = 3
	return c
}

func main() {
	p := NewParent()
	c1 := NewChild1()
	c2 := NewChild2()
	p.doit()
	c1.doit()
	c2.doit()
}

/*
Output:
parent.F()
child1.F()
parent.F()
*/
