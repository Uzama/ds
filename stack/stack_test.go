package stack

import "testing"

func TestStack(t *testing.T) {

	s := NewSatck(true)

	_, err := s.Pop()

	if err.Error() != "stack is empty" {
		t.Error("unexpected output")
	}

	s.Push(6)
	s.Push(7)

	val, _ := s.Pop()

	num, _ := val.(int)
	if num != 7 {
		t.Error("unexpected output")
	}

	s.Push(8)

	if s.Len() != 2 {
		t.Error("unexpected output")
	}

	val, _ = s.Pop()

	num, _ = val.(int)
	if num != 8 {
		t.Error("unexpected output")
	}

	val, _ = s.Pop()

	num, _ = val.(int)
	if num != 6 {
		t.Error("unexpected output")
	}
}
