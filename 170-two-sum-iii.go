type TwoSum struct {
    memo map[int]int
}

/** Initialize your data structure here. */
func Constructor() TwoSum {
    return TwoSum{
        memo: make(map[int]int),
    }
}

/** Add the number to an internal data structure.. */
func (this *TwoSum) Add(number int) {
    this.memo[number] += 1
}

/** Find if there exists any pair of numbers which sum is equal to the value. */
func (this *TwoSum) Find(value int) bool {
    for v, c := range this.memo {
        target := value - v
        if target != v {
            _, exist := this.memo[target]
            if exist {
                return true
            }
        } else {
            if c > 1 {
                return true
            }
        }
    }
    return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */