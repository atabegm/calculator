package stack

import "fmt"

type Stack struct {
	content []int
}

func (s *Stack) Push(x int) {
	s.content = append(s.content, x)
}

func (s *Stack) Pop() (int, error) {
	if s.empty() {
		return 0, fmt.Errorf("empty stack")
	}

	a := s.content[len(s.content)-1]
	s.content = s.content[:len(s.content)-1]

	return a, nil
}

func (s *Stack) Peek() int {
	return s.content[len(s.content)-1]
}

func (s *Stack) empty() bool {
	return len(s.content) == 0
}
