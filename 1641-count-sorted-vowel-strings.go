func countVowelStrings(n int) int {
	return dfs(0, n)
}

func dfs(start, length int) int {
	if length == 0 {
		return 1
	}
	var ans int
	// i = 0，代表从 a 开始，那么这个位置有 5 种可能性，循环 5 次
	// i = 1，代表从 e 开始，那么这个位置有 4 种可能性，循环 4 次
	// 以此类推
	for i := start; i < 5; i++ {
		ans += dfs(i, length - 1)
	}
	return ans
}

func countVowelStrings(n int) int {
	memo := make([][]int, 5)
	for i := range memo {
		memo[i] = make([]int, n + 1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return dfs(0, n, memo)
}

func dfs(start, length int, memo [][]int) int {
	if length == 0 {
		return 1
	}
	if memo[start][length] != -1 {
		return memo[start][length]
	}
	var ans int
	// i = 0，代表从 a 开始，那么这个位置有 5 种可能性，循环 5 次
	// i = 1，代表从 e 开始，那么这个位置有 4 种可能性，循环 4 次
	// 以此类推
	for i := start; i < 5; i++ {
		cnt := dfs(i, length - 1, memo)
		memo[i][length-1] = cnt
		ans += cnt
	}
	return ans
}