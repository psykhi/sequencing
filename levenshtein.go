package sequencing

// https://en.wikipedia.org/wiki/Levenshtein_distance?section=9#Iterative_with_two_matrix_rows
func LevenshteinDistance(a []byte, b []byte, v0 []int, v1 []int) int {
	m := len(a)
	n := len(b)

	if v0 == nil {
		v0 = make([]int, n+1, n+1)
	}
	if v1 == nil {
		v1 = make([]int, n+1, n+1)
	}

	for i := 0; i < n+1; i++ {
		v0[i] = i
	}

	for i := 0; i < m; i++ {
		v1[0] = i + 1

		for j := 0; j < n; j++ {
			substitutionCost := 0
			if a[i] == b[j] {
				substitutionCost = 0
			} else {
				substitutionCost = 1
			}
			v1[j+1] = min3(v1[j]+1, v0[j+1]+1, v0[j]+substitutionCost)
		}
		temp := v0
		v0 = v1
		v1 = temp
	}

	return v0[n]
}

func LevenshteinDistanceStrings(a []string, b []string, v0 []int, v1 []int) int {
	m := len(a)
	n := len(b)

	if v0 == nil {
		v0 = make([]int, n+1, n+1)
	}
	if v1 == nil {
		v1 = make([]int, n+1, n+1)
	}

	for i := 0; i < n+1; i++ {
		v0[i] = i
	}

	for i := 0; i < m; i++ {
		v1[0] = i + 1

		for j := 0; j < n; j++ {
			substitutionCost := 0
			if a[i] == b[j] {
				substitutionCost = 0
			} else {
				substitutionCost = 1
			}
			v1[j+1] = min3(v1[j]+1, v0[j+1]+1, v0[j]+substitutionCost)
		}
		temp := v0
		v0 = v1
		v1 = temp
	}

	return v0[n]
}
