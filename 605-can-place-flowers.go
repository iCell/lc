func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		left := (i == 0 || flowerbed[i-1] == 0)
		right := (i == len(flowerbed) - 1 || flowerbed[i+1] == 0)
		if flowerbed[i] == 0 && left && right {
			flowerbed[i] = 1
			n -= 1
		}
		if n <= 0 {
			return true
		}
	}
	return false
}