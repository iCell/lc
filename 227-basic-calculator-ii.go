func calculate(s string) int {
	stack, op, num := make([]int, 0), '+', 0
	for i, c := range s {
		digit := isDigit(c)
		
		if digit {
			num = num * 10 + int(c - '0')
		}
		if (!digit && c != ' ') || i == len(s) - 1 {
			switch op {
				case '+':
					stack = append(stack, num)
				case '-':
					stack = append(stack, -num)
				case '*':
					pop, temp := stack[len(stack)-1], stack[:len(stack)-1]
					stack = temp
					stack = append(stack, num * pop)
				case '/':
					pop, temp := stack[len(stack)-1], stack[:len(stack)-1]
					stack = temp
					stack = append(stack, pop / num)
			}
			op, num = c, 0
		}
	}

	var result int
	for len(stack) != 0 {
		pop, temp := stack[len(stack)-1], stack[:len(stack)-1]
		stack = temp
		result += pop
	}

	return result
}

func isDigit(r rune) bool {
	return r - '0' >= 0 && r - '9' <= 9
}