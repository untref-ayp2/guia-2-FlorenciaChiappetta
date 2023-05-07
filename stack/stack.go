package stack

import (
	"errors"
)

// Stack implementa una pila genérica sobre un arreglo dinámico.
type Stack struct {
	stack []any
}

// Push agrega un elemento a la pila. O(1)
func (s *Stack) Push(v any) {
	(*s).stack = append((*s).stack, v)
}

// Pop saca un elemento de la pila. O(1)
func (s *Stack) Pop() (any, error) {
	if (*s).IsEmpty() {
		return nil, errors.New("la pila esta vacia")
	}
	v := (*s).stack[len((*s).stack)-1]
	(*s).stack = ((*s).stack)[:len((*s).stack)-1]
	return v, nil
}

// Top devuelve el elemento del tope de la pila. O(1)
func (s *Stack) Top() (any, error) {
	if (*s).IsEmpty() {
		return 0, errors.New("la pila esta vacia")
	}
	v := (*s).stack[len((*s).stack)-1]
	return v, nil
}

// IsEmpty verifica si la pila esta vacia. O(1)
func (s *Stack) IsEmpty() bool {
	return len((*s).stack) == 0
}

func (s *Stack) Size() int {
	return len((*s).stack)
}
