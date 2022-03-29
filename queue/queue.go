package queue

import (
	"errors"
	"sync"
)

type queue struct {
	q          []interface{}
	lock       *sync.Mutex
	enableLock bool
	len        int
}

func NewQueue(enableLock bool) Queue {

	q := &queue{
		q:          make([]interface{}, 0),
		lock:       new(sync.Mutex),
		enableLock: enableLock,
	}

	return q
}

func (q *queue) Enqueue(value interface{}) {

	if q.enableLock {
		q.lock.Lock()
		defer q.lock.Unlock()
	}

	q.q = append(q.q, value)

	q.len += 1
}

func (q *queue) Dequeue() (interface{}, error) {

	if q.enableLock {
		q.lock.Lock()
		defer q.lock.Unlock()
	}

	if q.len == 0 {
		return nil, errors.New("queue is empty")
	}

	value := q.q[0]
	q.q = q.q[1:]

	q.len -= 1

	return value, nil
}

func (q *queue) Len() int {

	if q.enableLock {
		q.lock.Lock()
		defer q.lock.Unlock()
	}

	len := q.len

	return len
}
