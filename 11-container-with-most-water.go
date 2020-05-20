package main

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
