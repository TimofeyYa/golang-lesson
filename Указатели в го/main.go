package main

import "fmt"

// Parent is a parent struct.
type Parent struct {
	Name     string
	Children []Child
}

// Child is a child struct.
type Child struct {
	Name string
	Age  int
}

func main() {
	par := Parent{}
	par.Name = "jojo"

	cpPar := CopyParent(&par)
	cpPar.Name = "Pigls"

	fmt.Println(par, cpPar)
}

func CopyParent(p *Parent) Parent {
	if p == nil {
		return Parent{}
	}

	cp := *p

	cp.Children = make([]Child, len(p.Children))
	copy(cp.Children, p.Children)

	return cp
}
