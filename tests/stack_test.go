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

	aux, _ := s.Pop()
	if aux != 3 {
		t.Error("Error en Push")
	}
	aux, _ = s.Pop()
	if aux != 2 {
		t.Error("Error en Push")
	}
	aux, _ = s.Pop()
	if aux != 1 {
		t.Error("Error en Push")
	}

	// if s[0] != 1 || s[1] != 2 || s[2] != 3 {
	// 	t.Error("Error en Push")
	// }
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
