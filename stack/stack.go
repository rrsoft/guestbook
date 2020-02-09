// thread safe stack package
package stack

import (
	"sync"
)

type Stack struct {
	sync.Mutex // protects data
	count int
	data  []interface{}
}

func (s *Stack) Push(obj interface{}) {
	s.Lock()
	defer s.Unlock()
	s.data = append(s.data, obj)
	s.count++
}

func (s *Stack) Pop() interface{} {
	s.Lock()
	defer s.Unlock()
	if s.count > 0 {
		s.count--
		ret := s.data[s.count]
		s.data = s.data[0:s.count]
		return ret
	}
	return nil
}

func (s *Stack) Peek() interface{} {
	s.Lock()
	defer s.Unlock()
	if s.count > 0 {
		return s.data[s.count-1]
	}
	return nil
}

func (s *Stack) Clear() {
	s.Lock()
	defer s.Unlock()
	s.data = s.data[0:0]
	s.count = 0
}

func (s *Stack) Count() int {
	return s.count
}
