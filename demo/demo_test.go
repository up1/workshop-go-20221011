package demo_test

import (
	"demo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstCase(t *testing.T) {
	expected := "Hello"
	r, _ := demo.SayHi()
	if r != "Hello" {
		t.Errorf("SayHi() = %s want %s", r, expected)
	}
	assert.Equal(t, expected, r)
}
