package sequencing

func max3(a, b, c int) int {
	if b > c {
		if b > a {
			return b
		}
		return a
	} else {
		if c > a {
			return c
		}
		return a
	}
}

func min3(a, b, c int) int {
	if b < c {
		if b < a {
			return b
		}
		return a
	} else {
		if c < a {
			return c
		}
		return a
	}
}
