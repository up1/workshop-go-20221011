package main


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
	return AService{}
}

func main() {
	// Create object/instance
	a2 := NewAService(100)
	a2.id = 1
}