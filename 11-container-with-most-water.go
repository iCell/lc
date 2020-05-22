package main

import "math"

func maxArea1(height []int) int {
	max := 0
	for i := 0; i < len(height)-1; i++ {
		for j := 1; j < len(height); j++ {
			h := height[i]
			if h > height[j] {
				h = height[j]
			}
			area := (j - i) * h
			if area > max {
				max = area
			}
		}
	}
	return max
}

func maxArea2(height []int) int {
	max := 0

	i, j := 0, len(height)-1
	for {
		if i >= j {
			break
		}

		var h int
		if height[i] > height[j] {
			h = height[j]
			j--
		} else {
			h = height[i]
			i++
		}
		area := (j - i + 1) * h
		if area > max {
			max = area
		}
	}

	return max
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func maxAreaAgain(height []int) int {
	var (
		i = 0
		j = len(height) - 1
	)
	maxArea := 0
	for {
		if i >= j {
			break
		}

		maxArea = max(maxArea, (j-i)*min(height[i], height[j]))
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return maxArea
}
