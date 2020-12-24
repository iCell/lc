func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, 0, len(nums1))
	for _, num1 := range nums1 {
		start, n := false, num1
		for i, num2 := range nums2 {
			if num2 == num1 {
				start = true
			}
			if start {
				if num2 > num1 {
					n = num2
					break
				} else if i == len(nums2) - 1 {
					n = -1
				}
			}
		}
		ans = append(ans, n)
	}
	return ans
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack, memo := NewStack(), make(map[int]int)
	for _, num := range nums2 {
		for stack.IsEmpty() == false && stack.Peek() < num {
			pop := stack.Pop()
			memo[pop] = num
		}
		stack.Push(num)
	}
	for stack.IsEmpty() == false {
		pop := stack.Pop()
		memo[pop] = -1
	}
	ans := make([]int, 0, len(nums1))
	for _, num := range nums1 {
		ans = append(ans, memo[num])
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