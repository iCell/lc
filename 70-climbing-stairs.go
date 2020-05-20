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
