func partition(s string) [][]string {
	result := make([][]string, 0)
	if len(s) == 0 {
		return result
	}
	backtrack(s, 0, []string{}, &result)
	return result
}

func backtrack(s string, index int, path []string, result *[][]string) {
	if index >= len(s) {
		*result = append(*result, append([]string{}, path...))
		return
	}
	
	for i := index; i < len(s); i++ {
		pre := s[index:i+1]
		if !isPalindrome(pre) {
			continue
		}
		path = append(path, pre)
		backtrack(s, i + 1, path, result)
		path = path[:len(path)-1]
	}
}

func isPalindrome(s string) bool {
	left, right := 0, len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left += 1
		right -= 1
	}
	return true
}