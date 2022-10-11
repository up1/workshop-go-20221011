package main

import (
	"demo"
	"fmt"
)

func main() {
	// Loop :: for
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Loop while
	i := 1
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// Infinity loop
	// for {}

	m, err := demo.SayHi()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
}
