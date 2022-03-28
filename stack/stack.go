package stack

import (
	"errors"
	"sync"
)

type stack struct {
	s          []interface{}
	len        int
	lock       *sync.Mutex
	enableLock bool
}

// if the stack use as a global storage, enabling lock is must
func NewSatck(enableLock bool) Stack {

	s := &stack{
		s:          make([]interface{}, 0),
		enableLock: enableLock,
		lock:       new(sync.Mutex),
	}

	return s
}

func (s *stack) Push(value interface{}) {

	if s.enableLock {
		s.lock.Lock()
		defer s.lock.Unlock()
	}

	s.s = append(s.s, value)
	s.len += 1
}

func (s *stack) Pop() (interface{}, error) {

	if s.enableLock {
		s.lock.Lock()
		defer s.lock.Unlock()
	}

	if len(s.s) == 0 {
		return nil, errors.New("stack is empty")
	}

	value := s.s[len(s.s)-1]

	s.s = s.s[:len(s.s)-1]

	s.len -= 1

	return value, nil
}

func (s *stack) Len() int {

	if s.enableLock {
		s.lock.Lock()
		defer s.lock.Unlock()
	}

	len := s.len

	return len
}
