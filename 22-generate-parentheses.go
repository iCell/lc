package main

func main() {
	generateParenthesis(3)
}

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	generate(&result, n, n, "")
	return result
}

func generate(res *[]string, left int, right int, str string) {
	if left == 0 && right == 0 {
		*res = append(*res, str)
		return
	}
	if left > 0 {
		generate(res, left-1, right, str+"(")
	}
	if left < right {
		generate(res, left, right-1, str+")")
	}
}

func generateParenthesis2(n int) []string {
	result := make([]string, 0)
	helper(&result, n, n, "")
	return result
}

func helper(res *[]string, left, right int, result string) {
	if left == 0 && right == 0 {
		*res = append(*res, result)
		return
	}

	if left > 0 {
		result = result + "("
		helper(res, left-1, right, result)
		result = result[:len(result)-1]
	}
	if left < right {
		result = result + ")"
		helper(res, left, right-1, result)
		result = result[:len(result)-1]
	}
}
