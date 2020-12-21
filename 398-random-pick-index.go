type Solution struct {
	nums []int
	r *rand.Rand
}

func Constructor(nums []int) Solution {
	solution := Solution{
		nums: nums,
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	return solution
}

func (this *Solution) Pick(target int) int {
	index := 0
	count := 1
	for i := 0; i < len(this.nums); i++ {
		if this.nums[i] == target {
			if this.r.Intn(count)+1 == count {
				index = i
			}
			count++
		}
	}
	return index
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */