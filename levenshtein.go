package sequencing

// https://en.wikipedia.org/wiki/Levenshtein_distance?section=9#Iterative_with_two_matrix_rows
func LevenshteinDistance(a []byte, b []byte) int {

	m := len(a)
	n := len(b)

	v0 := make([]int, n+1, n+1)
	v1 := make([]int, n+1, n+1)

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
			//min1 := math.Min(float64(v1[j]+1), float64(v0[j+1]+1))

			//v1[j+1] = int(math.Min(min1, float64(v0[j]+substitutionCost)))
			v1[j+1] = min3(v1[j]+1, v0[j+1]+1, v0[j]+substitutionCost)
		}
		temp := v0
		v0 = v1
		v1 = temp
	}

	return v0[n]
}
