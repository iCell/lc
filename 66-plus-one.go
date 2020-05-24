package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{9, 9, 9}))
}

type Stack struct {
	values []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) isEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack) push(v int) {
	s.values = append(s.values, v)
}

func (s *Stack) pop() int {
	if len(s.values) == 0 {
		return -1
	}
	x, v := s.values[len(s.values)-1], s.values[:len(s.values)-1]
	s.values = v
	return x
}

func plusOne(digits []int) []int {
	st := NewStack()

	plus := true

	for i := len(digits) - 1; i >= 0; i-- {
		digit := digits[i]
		if plus {
			digit++
		}
		if digit == 10 {
			plus = true
			st.push(0)
			if i == 0 {
				st.push(1)
			}
		} else {
			plus = false
			st.push(digit)
		}
	}

	var result []int
	for {
		p := st.pop()

		if p == -1 {
			return result
		}
		result = append(result, p)
	}
}
