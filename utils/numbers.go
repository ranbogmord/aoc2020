package utils

func Max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func AbsInt(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}
