func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	ans, light, heavy := 0, 0, len(people) - 1
	for light <= heavy {
		if people[light] + people[heavy] <= limit {
			light++
		}
		heavy--
		ans++ 
	}
	return ans
}