package main

// input 1:
// 1 step
// input 2:
//	1 step + 1 step
//	2 steps
// input 3:
// from 3 - 1: input 1
// from 3 - 2: input 2
// input 4:
// from 4 - 1: input 3
// from 4 - 2: input 2
// ...
// input n:
// from n - 1: input n-1
// from n - 2: input n-2
// f(n) = f(n-1) + f(n-2)

// stupid solution
func climbStairs2(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 2
	default:
		return climbStairs2(n-1) + climbStairs2(n-2)
	}
}

func climbStairs(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 2
	}

	var cache = make(map[int]int, n)
	cache[1] = 1
	cache[2] = 2

	for i := 3; i <= n; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}

	return cache[n]
}

// f[n] = f[n-1] + f[n-2]
func climbStairs3(n int) int {
	cache := make(map[int]int, n)
	return helper(n, cache)
}

func helper(n int, cache map[int]int) int {
	if n <= 2 {
		return n
	}
	cached, ok := cache[n]
	if ok {
		return cached
	}
	result := helper(n-1, cache) + helper(n-2, cache)
	cache[n] = result
	return result
}
