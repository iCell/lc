package main

func lemonadeChange(bills []int) bool {
	maps := make(map[int]int)
	for _, num := range bills {
		maps[num] += 1
		change := num - 5
		if change == 0 {
			continue
		}
		switch change {
		case 5:
			if maps[5] <= 0 {
				return false
			}
			maps[5] -= 1
		case 15:
			if maps[10] > 0 && maps[5] > 0 {
				maps[5] -= 1
				maps[10] -= 1
			} else if maps[10] == 0 && maps[5] >= 3 {
				maps[5] -= 3
			} else {
				return false
			}
		}
	}
	return true
}
