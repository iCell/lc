func nextGreaterElements(nums []int) []int {
	ans := make([]int, len(nums))
	for i := range ans {
		ans[i] = -1
	}
	
	stack := NewStack()
	for i := 0; i < len(nums) * 2; i++ {
		num := nums[i%len(nums)]
		for stack.IsEmpty() == false && nums[stack.Peek()] < num {
			pop := stack.Pop()
			ans[pop] = num
		}
		stack.Push(i%len(nums))
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