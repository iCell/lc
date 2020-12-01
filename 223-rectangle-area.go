func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	area1, area2 := (C - A) * (D - B), (G - E) * (H - F)
	if E > C || F > D || G < A || H < B {
		return area1 + area2
	}
	left, right := max(A, E), min(C, G)
	top, bottom := min(H, D), max(B, F)
	return area1 + area2 - (right - left) * (top - bottom);
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}