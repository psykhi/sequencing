package sequencing

import "math"

// Compute Levenshtein distance on byte arrays
// https://en.wikipedia.org/wiki/Levenshtein_distance?section=9#Iterative_with_two_matrix_rows
// v0 and v1 buffers can be provided in order to reuse them and avoid allocations
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

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func LevenshteinDistanceK(a []byte, b []byte, v0 []int, v1 []int, k int) int {
	m := len(a)
	n := len(b)

	if n == 0 {
		if m <= k {
			return m
		}
		return -1
	} else if m == 0 {
		if n <= k {
			return n
		}
	}

	if n > m {
		n, m = m, n
		a, b = b, a
	}

	if v0 == nil {
		v0 = make([]int, n+1, n+1)
	}
	if v1 == nil {
		v1 = make([]int, n+1, n+1)
	}

	boundary := min(n, k) + 1
	for i := 0; i < boundary; i++ {
		v0[i] = i
	}
	for i := boundary; i < n+1; i++ {
		v0[i] = math.MaxInt32
	}
	for i := 0; i < n+1; i++ {
		v1[i] = math.MaxInt32
	}

	for i := 1; i <= m; i++ {
		v1[0] = i + 1

		minStripe := max(1, i-k)

		max := min(n, i+k)

		if i > math.MaxInt32-k {
			max = n
		}

		if minStripe > max {
			return -1
		}
		if minStripe > 1 {
			v1[minStripe-1] = math.MaxInt32
		}

		for j := minStripe; j <= max; j++ {
			if b[j-1] == a[i-1] {
				v1[j] = v0[j-1]
			} else {
				v1[j] = 1 + min3(v1[j-1], v0[j], v0[j-1])
			}

		}
		v0, v1 = v1, v0
	}

	if v0[n] <= k {
		return v0[n]
	}
	return -1
}

func LevenshteinDistanceKStrings(a []string, b []string, v0 []int, v1 []int, k int) int {
	m := len(a)
	n := len(b)

	if n == 0 {
		if m <= k {
			return m
		}
		return -1
	} else if m == 0 {
		if n <= k {
			return n
		}
	}

	if n > m {
		n, m = m, n
		a, b = b, a
	}

	if v0 == nil {
		v0 = make([]int, n+1, n+1)
	}
	if v1 == nil {
		v1 = make([]int, n+1, n+1)
	}

	boundary := min(n, k) + 1
	for i := 0; i < boundary; i++ {
		v0[i] = i
	}
	for i := boundary; i < n+1; i++ {
		v0[i] = math.MaxInt32
	}
	for i := 0; i < n+1; i++ {
		v1[i] = math.MaxInt32
	}

	for i := 1; i <= m; i++ {
		v1[0] = i + 1

		minStripe := max(1, i-k)

		max := min(n, i+k)

		if i > math.MaxInt32-k {
			max = n
		}

		if minStripe > max {
			return -1
		}
		if minStripe > 1 {
			v1[minStripe-1] = math.MaxInt32
		}

		for j := minStripe; j <= max; j++ {
			if b[j-1] == a[i-1] {
				v1[j] = v0[j-1]
			} else {
				v1[j] = 1 + min3(v1[j-1], v0[j], v0[j-1])
			}

		}
		v0, v1 = v1, v0
	}

	if v0[n] <= k {
		return v0[n]
	}
	return -1
}

// Compute Levenshtein distance on strings sequences
// https://en.wikipedia.org/wiki/Levenshtein_distance?section=9#Iterative_with_two_matrix_rows
// v0 and v1 buffers can be provided in order to reuse them and avoid allocations
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

func LevenshteinDistanceInterface(a Sequence, b Sequence, v0 []int, v1 []int) int {
	m := a.Len()
	n := b.Len()

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
			if a.Val(i) == b.Val(j) {
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
