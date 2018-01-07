package alignment

var GapPenalty = -1

func similarity(a byte, b byte) int {
	if a == b {
		return 1
	}
	return -1
}

func NeedlemanWunsch(x []byte, y []byte) ([]byte, []byte) {
	sub := 0
	ins := 0
	del := 0
	same := 0
	// Build grid
	f := make([][]int, len(x))
	for i := 0; i < len(x); i++ {
		f[i] = make([]int, len(y))
		f[i][0] = GapPenalty * i
	}
	for j := 0; j < len(y); j++ {
		f[0][j] = GapPenalty * j
	}
	for i := 1; i < len(x); i++ {
		for j := 1; j < len(y); j++ {
			match := f[i-1][j-1] + similarity(x[i], y[j])
			del := f[i-1][j] + GapPenalty
			ins := f[i][j-1] + GapPenalty
			f[i][j] = max(match, del, ins)
		}
	}

	// Align
	z := make([]byte, 0)
	w := make([]byte, 0)

	i := len(x) - 1
	j := len(y) - 1
	for i > 0 || j > 0 {
		if i > 0 && j > 0 && f[i][j] == f[i-1][j-1]+similarity(x[i], y[j]) {
			if similarity(x[i], y[j]) == -1 {
				sub++
			} else {
				same++
			}
			z = append([]byte{x[i]}, z...)
			w = append([]byte{y[j]}, w...)
			i--
			j--
		} else if i > 0 && f[i][j] == f[i-1][j]+GapPenalty {
			del++
			z = append([]byte{x[i]}, z...)
			w = append([]byte{'-'}, w...)
			i--
		} else {
			ins++
			z = append([]byte{'-'}, z...)
			w = append([]byte{y[j]}, w...)
			j--
		}
	}
	if i == 0 && j == 0 {
		z = append([]byte{x[i]}, z...)
		w = append([]byte{y[j]}, w...)
	}

	//fmt.Printf("LEN: %d|%d, SAME: %d, SUB: %d, DEL: %d, INSERT: %d\n", len(x), len(y), same, sub, del, ins)
	return z, w
}
