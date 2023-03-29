package ejercicios

import (
	"fmt"
	"guia2/queue"
	"guia2/stack"
	"strings"
)

func InvertirCadena(z string) string { //O(n)

	stack2 := stack.Stack{}
	stack := stack.Stack{}
	resultado := ""

	splits := strings.Split(z, "")

	for i := 0; i < len(splits); i++ {
		stack.Push(splits[i])

	}
	for i := 0; i < len(splits); i++ {
		aux, _ := stack.Pop()
		stack2.Push(aux)
		resultado += fmt.Sprint(aux)
	}

	return resultado
}

func Palindromo(z string) bool {
	stack_aux := stack.Stack{}
	stack2 := stack.Stack{}
	stack := stack.Stack{}

	resultado1 := ""
	resultado2 := ""

	splits := strings.Split(z, "")

	for i := 0; i < len(splits)/2; i++ {
		stack.Push(splits[i])
		resultado1 += fmt.Sprint(splits[i])
	}

	for !stack.IsEmpty() {
		aux, _ := stack.Pop()
		stack_aux.Push(aux)
	}
	for !stack_aux.IsEmpty() {
		aux, _ := stack_aux.Pop()
		stack2.Push(aux)
		resultado2 += fmt.Sprint(aux)
	}
	return resultado1 == resultado2
}

func Balanceada(x string) bool {
	stack := stack.Stack{}
	for i := 0; i < len(x); i++ {

		string_ := rune(x[i])
		if string_ == '(' || string_ == '[' || string_ == '{' {
			stack.Push(string_)
		} else if string_ == ')' {
			aux, _ := stack.Pop()
			if aux != '(' {
				return false
			}
		} else if string_ == ']' {
			aux, _ := stack.Pop()
			if aux != '[' {
				return false
			}
		} else if string_ == '}' {
			aux, _ := stack.Pop()
			if aux != '{' {
				return false
			}
		}
	}
	return stack.IsEmpty()
}

func UnirColas(q1, q2 queue.Queue) queue.Queue {
	for !q2.IsEmpty() {
		aux, _ := q2.Dequeue()
		q1.Enqueue(aux)
	}
	return q1
}
