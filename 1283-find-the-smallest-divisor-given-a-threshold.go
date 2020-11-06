func smallestDivisor(nums []int, threshold int) int {
	max := math.MinInt64
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	
	left, right := 1, max
	for left < right {
		mid := (left + right) / 2
		
		var sum int
		for _, num := range nums {
			sum += num / mid
			if num % mid != 0 {
				sum++
			}
		}
		
		if sum > threshold {
			left = mid + 1
		} else {
			right = mid
		}
	}
	
	return left
}
