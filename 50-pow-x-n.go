package main

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	an := n
	if an < 0 {
		an = -an
	}
	r := pow(x, an)
	if n >= 0 {
		return r
	}
	return 1 / r
}

func pow(x float64, n int) float64 {
	if n == 1 {
		return x
	}

	r := pow(x, n/2)
	if n%2 == 1 {
		return r * r * x
	}
	return r * r
}
