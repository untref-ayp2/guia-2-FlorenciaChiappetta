package queue

import (
	"errors"
	"guia2/stack"
)

// Queue implementa una cola genérica sobre un arreglo dinámico.
type Queue struct {
	cola []any
}

// Enqueue agrega un elemento a la cola. O(1)
func (q *Queue) Enqueue(v any) {
	(*q).cola = append((*q).cola, v)
}

// Dequeue saca un elemento de la cola. O(1)
func (q *Queue) Dequeue() (any, error) {

	if len((*q).cola) == 0 {
		return nil, errors.New("queue is empty")
	}
	head := ((*q).cola)[0]
	(*q).cola = (*q).cola[1:]
	return head, nil
}

// Front devuelve el elemento del frente de la cola. O(1)
func (q *Queue) Front() (any, error) {
	if len((*q).cola) == 0 {
		return nil, errors.New("queue is empty")
	}
	return ((*q).cola)[0], nil
}

// IsEmpty verifica si la cola esta vacia. O(1)
func (q *Queue) IsEmpty() bool {
	return len((*q).cola) == 0
}

type QueueS struct {
	stack1 stack.Stack
	stack2 stack.Stack
}

//valida que ambas pilas esten vacias

func (q *QueueS) IsEmpty() bool {
	return q.stack1.IsEmpty() && q.stack2.IsEmpty()
}

//si ambas pilas estan vacias, agrega un elemento a la pila1
//si la pila 2 no esta vacia, elimina el elemento y lo guarda en la pila1.
// si la pila2 esta vacia, lo almacena en la pila1.

func (q *QueueS) Enqueue(v any) {
	if q.IsEmpty() {
		q.stack1.Push(v)
	} else {
		for !q.stack2.IsEmpty() {
			aux, _ := q.stack2.Pop()
			q.stack1.Push(aux)
		}
		q.stack1.Push(v)
	}

}

// Devuelve el elemento del frente de la cola y lo elimina
// valida si la cola esta vacia, si la pila1 no esta vacia entonces agrega el elemento en la pila2 y lo elimina de la pila1.
// si la pila1 esta vacia,
func (q *QueueS) Dequeue() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	for !q.stack1.IsEmpty() {
		aux, _ := q.stack1.Pop()
		q.stack2.Push(aux)
	}
	head, _ := q.stack2.Pop()
	return head, nil
}

// Devuelve el elemento del frente de la cola, pero no lo elimina
func (q *QueueS) Front() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	for !q.stack1.IsEmpty() {
		aux, _ := q.stack1.Pop()
		q.stack2.Push(aux)
	}
	head, _ := q.stack2.Top()

	return head, nil
}
func UnirColas(q1, q2 Queue) (q Queue) {
	if q1.IsEmpty() {
		return q2
	}
	if q2.IsEmpty() {
		return q1
	}
	q1.Enqueue(q2)

	return q1

}

// Stack implementa una pila genérica sobre un arreglo dinámico.
type Stack struct {
	stack []any
}

// Push agrega un elemento a la pila. O(1)
func (s *Stack) Push(v any) {
	s.stack = append(s.stack, v)
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

func TransformarAPila(q1 *Queue) (s1 *Stack) {

	if q1.IsEmpty() {
		return s1
	}

	elemento, _ := q1.Dequeue()
	s1.Push(elemento)
	TransformarAPila(q1)
	return s1
}
