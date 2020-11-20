func decodeString(s string) string {
	var result []rune
	
	stack := make([]rune, 0)
	for _, c := range s {
		if isDigit(c) || isLetter(c) || c == '[' {
			stack = append(stack, c)
		} else if c == ']' {
			
			var inner []rune
			
			for len(stack) != 0 && stack[len(stack) - 1] != '[' {
				last, temp := stack[len(stack)-1], stack[:len(stack)-1]
				stack = temp
				
				inner = append(inner, last)
			}
			
			stack = stack[:len(stack)-1]
			
			var cnt []rune
			for len(stack) != 0 && isDigit(stack[len(stack) - 1]) {
				last, temp := stack[len(stack)-1], stack[:len(stack)-1]
				stack = temp
				
				cnt = append(cnt, last)
			}
			
			cntStr := string(reverse(cnt))
			count, _ := strconv.Atoi(cntStr)

			multipled := multiple(inner, count)
			reversed := reverse(multipled)

			stack = append(stack, reversed...)
		}
	}
	
	for len(stack) != 0 {
		last, temp := stack[len(stack)-1], stack[:len(stack)-1]
		stack = temp
		
		result = append(result, last)
	}
	
	reverse(result)
	
	return string(result)
}

func isDigit(b rune) bool {
	return b >= '0' && b <= '9'
}

func isLetter(b rune) bool {
	return b >= 'a' && b <= 'z' 
}

func multiple(inner []rune, cnt int) []rune {
	length := len(inner) * cnt
	result := make([]rune, 0, length)
	for i := 0; i < length; i++ {
		result = append(result, inner[i%len(inner)])
	}
	return result
}

func reverse(bs []rune) []rune {
	i, j := 0, len(bs) - 1
	for i < j {
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
	return bs
}