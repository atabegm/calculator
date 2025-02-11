package stack

import "fmt"

type Stack struct {
	content []string
}

func New() Stack {
	return Stack{
		content: make([]string, 0),
	}
}

func (s *Stack) Push(x string) {
	s.content = append(s.content, x)
}

func (s *Stack) Pop() (string, error) {
	if s.Empty() {
		return "", fmt.Errorf("empty stack")
	}

	a := s.content[len(s.content)-1]
	s.content = s.content[:len(s.content)-1]

	return a, nil
}

func (s *Stack) Peek() (string, error) {
	if s.Empty() {
		return "", fmt.Errorf("empty stack")
	}

	return s.content[len(s.content)-1], nil
}

func (s *Stack) Empty() bool {
	return len(s.content) == 0
}
