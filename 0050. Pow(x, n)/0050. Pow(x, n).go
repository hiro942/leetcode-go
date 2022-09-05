package leetcode

func myPow(x float64, n int) float64 {
	isNeg := false
	if n < 0 {
		n = -n
		isNeg = true
	}

	res := 1.0

	for n != 0 {
		if n&1 == 1 {
			res = res * x
		}
		n >>= 1
		x = x * x
	}

	if isNeg {
		res = 1 / res
	}

	return res
}
