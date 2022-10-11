package main

import "fmt"


type A1 struct {
	id int
}

// Composition over inheritance
type AService struct {
	name string
	A1
}

//  Builder patterns
func NewAService(id int) AService {
	a1 := A1{999}
	return AService{"", a1}
}

func main() {
	// Create object/instance
	a2 := NewAService(100)
	fmt.Println(a2.id)
}