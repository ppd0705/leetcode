package main

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	}

	if n < 0 {
		n = -n
		x = 1 / x
	}

	sub := myPow(x, n/2)

	if n%2 == 1 {
		return x * sub * sub
	} else {
		return sub * sub
	}
}

func myPow2(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	}

	if n < 0 {
		n = -n
		x = 1 / x
	}
	ans := 1.0
	for n > 0 {
		if n%2 == 1 {
			ans *= x
		}
		x *= x
		n = n / 2
	}
	return ans
}
