package queue

import "testing"

func TestQueue(t *testing.T) {

	q := NewQueue(true)

	_, err := q.Dequeue()

	if err.Error() != "queue is empty" {
		t.Error("unexpected output")
	}

	q.Enqueue(6)
	q.Enqueue(7)

	val, _ := q.Dequeue()

	num, _ := val.(int)
	if num != 6 {
		t.Error("unexpected output")
	}

	q.Enqueue(8)

	if q.Len() != 2 {
		t.Error("unexpected output")
	}

	val, _ = q.Dequeue()

	num, _ = val.(int)
	if num != 7 {
		t.Error("unexpected output")
	}

	val, _ = q.Dequeue()

	num, _ = val.(int)
	if num != 8 {
		t.Error("unexpected output")
	}
}
