package demo_test

import (
	"demo"
	"testing"
)

func TestFirstCase(t *testing.T) {
	expected := "Hello"
	r, _ := demo.SayHi()
	if r != "Hello" {
		t.Errorf("SayHi() = %s want %s", r, expected)
	}
}