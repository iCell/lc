package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isValid("()"))
}

type Stack struct {
	values []rune
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) push(v rune) {
	s.values = append(s.values, v)
}

func (s *Stack) pop() rune {
	if len(s.values) == 0 {
		return 0
	}
	x, v := s.values[len(s.values)-1], s.values[:len(s.values)-1]
	s.values = v
	return x
}

const (
	left1  = rune('(')
	right1 = rune(')')
	left2  = rune('[')
	right2 = rune(']')
	left3  = rune('{')
	right3 = rune('}')
)

var left = string([]rune{left1, left2, left3})
var right = string([]rune{right1, right2, right3})

func isValid(s string) bool {
	if len(s) == 1 {
		return false
	}

	st := NewStack()

	for _, r := range []rune(s) {
		switch {
		case strings.ContainsRune(left, r):
			st.push(r)
		case strings.ContainsRune(right, r):
			l := st.pop()
			if l == 0 {
				return false
			}
			diff := r - l
			if diff != 1 && diff != 2 {
				return false
			}
		}
	}
	return len(st.values) == 0
}
