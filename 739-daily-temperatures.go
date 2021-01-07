func dailyTemperatures(T []int) []int {
	stack, memo := NewStack(), make(map[int]int)
	
	for index, temperature := range T {
		for !stack.IsEmpty() && temperature > T[stack.Peek()] {
			pop := stack.Pop()
			memo[pop] = index
		}
		stack.Push(index)
	}
	
	var ans []int
	for i := range T {
		l, ok := memo[i]
		if ok {
			ans = append(ans, l - i)   
		} else {
			ans = append(ans, 0)   
		}
	}
	return ans
}

type Stack struct {
	values []int
}

func NewStack() *Stack {
	return &Stack{values: make([]int, 0)}
}

func (s *Stack) Push(v int) {
	s.values = append(s.values, v)
}

func (s *Stack) Pop() int {
	pop, temp := s.values[len(s.values) - 1], s.values[:len(s.values) - 1]
	s.values = temp
	return pop
}

func (s *Stack) Peek() int {
	return s.values[len(s.values) - 1]   
}

func (s *Stack) IsEmpty() bool {
	return len(s.values) == 0
}