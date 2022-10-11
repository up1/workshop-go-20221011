package main

import "fmt"

// class
//  + properties
//  + behaviors

type B interface {
	process() string
}

type Data struct {
	Id int
	Name string
}

func (d Data) process() string {
	return fmt.Sprintf("Call process with id=%d", d.Id)
}

func (d Data) String() string {
	return fmt.Sprintf("[Id=%d, Name=%s]", d.Id, d.Name);
}

func doSth(b B) {
	fmt.Println(b.process())
}

func main() {
	d2 := Data{Id: 100, Name: "Original name"}
	fmt.Println(d2)
	doSth(d2)
}