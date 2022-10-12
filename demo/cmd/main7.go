package main

import "fmt"

type BusinessError struct {
	code int
	message string
}

func (be BusinessError) Error() string {
	return fmt.Sprintf("Business error with code=%d", be.code)
}

func main() {
	r, err := doWithError()
	fmt.Println(r, err)
}

func doWithError() (int, error) {
	return 1, BusinessError{code: 404}
}
