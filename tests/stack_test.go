package tests

import (
	"guia2/stack"
	"testing"
)

func TestPush(t *testing.T) {
	var s stack.Stack

	s.Push(1)
	s.Push(2)
	s.Push(3)

	a, _ := s.Pop()
	b, _ := s.Pop()
	c, _ := s.Pop()

	if a != 1 || b != 2 || c != 3 {
		t.Error("Error en Push")
	}
}

func TestPop(t *testing.T) {
	var s stack.Stack

	s.Push(1)
	s.Push(2)
	s.Push(3)

	v, err := s.Pop()

	if err != nil || v != 3 {
		t.Error("Error en Pop")
	}

	v, err = s.Pop()

	if err != nil || v != 2 {
		t.Error("Error en Pop")
	}

	v, err = s.Pop()

	if err != nil || v != 1 {
		t.Error("Error en Pop")
	}

	_, err = s.Pop()

	if err == nil {
		t.Error("Error en Pop")
	}
}
