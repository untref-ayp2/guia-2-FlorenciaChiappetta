package ejercicios

import (
	"guia2/queue"
	"guia2/stack"
)

// Escribir una función que que reciba una cadena de caracteres y devuelva la cadena invertida. Analizar el orden.
func InvertirCadena(value string) string {
	var stack stack.Stack
	var valorRetorno string
	for _, valor := range value {
		stack.Push(string(valor))
	}

	for i := 0; i < len(value); i++ {
		caracter, _ := stack.Pop()
		valorRetorno += caracter.(string)
	}
	return valorRetorno
}

// Escribir una función que verifique si una palabra es palíndromo
// (es decir que una cadena es igual a su inversa. Por ejemplo: las cadenas "1456541" y "145541" son palíndromos). Analizar el orden.
func Palindromo(value string) bool {
	Palabra := InvertirCadena(value)
	return Palabra == value
}

// Escribir una función que evalúe si una cadena de paréntesis, corchetes y llaves está bien balanceada y
// devuelve `true` si está bien balanceada y `false` si está mal balanceada.
// Por ejemplo: `[()]{}{[()()]()}` debe devolver `true`, mientras que `[(])` debe devolver `false`. Analizar el orden.
func Balanceada(string) bool {

}

//Escribir una función, tal que, dadas dos colas,
//construya una cola con el resultado de poner una a continuación de la otra.
//Por ejemplo: si `q1 = [1,2,3]` (con 1 al frente y 3 al final) y `q2 = [5,7]`,
//el resultado es `[1,2,3,5,7]` (con 1 al frente y 7 al final). Analizar el orden.

func UnirColas(cola1, cola2 queue.Queue) queue.Queue {

	for !cola2.IsEmpty() {
		value, _ := cola2.Dequeue()
		cola1.Enqueue(value)
	}
	return cola1

}
