package sequencing

func LCSStrings(a []string, b []string, v0 []int, v1 []int) int {
	m := len(a)
	n := len(b)

	if v0 == nil {
		v0 = make([]int, n+1, n+1)
	}
	if v1 == nil {
		v1 = make([]int, n+1, n+1)
	}

	for i := 0; i < m; i++ {
		v1[0] = 0

		for j := 0; j < n; j++ {
			if a[i] == b[j] {
				v1[j+1] = v0[j] + 1
			} else {
				v1[j+1] = max(v1[j], v0[j+1])
			}
		}
		temp := v0
		v0 = v1
		v1 = temp
	}

	return v0[n]
}

func LCS(a []byte, b []byte, v0 []int, v1 []int) int {
	m := len(a)
	n := len(b)

	if v0 == nil {
		v0 = make([]int, n+1, n+1)
	}
	if v1 == nil {
		v1 = make([]int, n+1, n+1)
	}

	for i := 0; i < m; i++ {
		v1[0] = 0

		for j := 0; j < n; j++ {
			if a[i] == b[j] {
				v1[j+1] = v0[j] + 1
			} else {
				v1[j+1] = max(v1[j], v0[j+1])
			}
		}
		temp := v0
		v0 = v1
		v1 = temp
	}

	return v0[n]
}
