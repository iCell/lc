import "math/rand"

type RandomizedSet struct {
    valueToIndex map[int]int
    values       []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
    return RandomizedSet{
        valueToIndex: make(map[int]int),
        values:       make([]int, 0),
    }
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
    _, exist := this.valueToIndex[val]
    if exist {
        return false
    }
    this.values = append(this.values, val)
    this.valueToIndex[val] = len(this.values) - 1
    return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
    index, exist := this.valueToIndex[val]
    if !exist {
        return false
    }

    lastIdx, last := len(this.values)-1, this.values[len(this.values)-1]
    this.values[index], this.values[lastIdx] = this.values[lastIdx], this.values[index]

    this.values = this.values[:lastIdx]
    this.valueToIndex[last] = index
    delete(this.valueToIndex, val)

    return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
    index := rand.Intn(len(this.values))
    return this.values[index]
}

