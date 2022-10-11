package main

import (
	"fmt"
	"io/ioutil"
	"runtime/debug"
)


func main() {
	defer panicHandler()

    // Read data from file
    b, err := ioutil.ReadFile("try_panic.go")
    if err != nil {
		// log.Fatalf("Error %v", err);        
		panic("some error")
    }

    fmt.Println(string(b))
}

func panicHandler() {
    err := recover()
    if err == "some error" {
        fmt.Println("Try to recover from panic")
        // TODO
        debug.PrintStack()
    }
}