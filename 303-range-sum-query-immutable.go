import "fmt"

type NumArray struct {
    Nums []int
    hash map[string]int
}

func Constructor(nums []int) NumArray {
    h := make(map[string]int, len(nums))
    return NumArray{
        Nums: nums,
        hash: h,
    }
}

func (this *NumArray) SumRange(i int, j int) int {
    key := fmt.Sprintf("%d:%d", i, j)
    sum, ok := this.hash[key]
    if ok {
        return sum
    }
    for idx := i; idx <= j; idx++ {
        sum += this.Nums[idx]
    }
    this.hash[key] = sum
    return sum
}

type NumArray struct {
    Nums []int
}

func Constructor(nums []int) NumArray {
    dp := make([]int, len(nums)+1, len(nums)+1)
    for i := 0; i < len(nums); i++ {
        dp[i+1] = dp[i] + nums[i]
    }
    return NumArray{Nums: dp}
}

func (this *NumArray) SumRange(i int, j int) int {
    return this.Nums[j+1] - this.Nums[i]
}

