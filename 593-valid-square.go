func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	dists := make(map[int]bool)
	dists[distance(p1, p2)] = true
	dists[distance(p1, p3)] = true
	dists[distance(p1, p4)] = true
	dists[distance(p2, p3)] = true
	dists[distance(p2, p4)] = true
	dists[distance(p3, p4)] = true
	if dists[-1] {
		return false
	}
	return len(dists) == 2
}

func distance(p1, p2 []int) int {
	if p1[0] == p2[0] && p1[1] == p2[1] {
		return -1
	}
	return (p1[0] - p2[0]) * (p1[0] - p2[0]) + (p1[1] - p2[1]) * (p1[1] - p2[1])
}
