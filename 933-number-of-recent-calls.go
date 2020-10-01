type RecentCounter struct {
	Values []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		Values: make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	for len(this.Values) > 0 && this.Values[0] < t-3000 {
		this.Values = this.Values[1:]
	}
	this.Values = append(this.Values, t)
	return len(this.Values)
}