package sequencing

func NeedlemanWunsch(x []byte, y []byte, gap int, similarity func(byte, byte) int, f [][]int) ([]byte, []byte) {
	// Build grid
	if f == nil {
		f = make([][]int, len(x))
		for i := 0; i < len(x); i++ {
			f[i] = make([]int, len(y))
		}
	}

	for i := 0; i < len(x); i++ {
		f[i][0] = gap * i
	}
	for j := 0; j < len(y); j++ {
		f[0][j] = gap * j
	}
	for i := 1; i < len(x); i++ {
		for j := 1; j < len(y); j++ {
			match := f[i-1][j-1] + similarity(x[i], y[j])
			del := f[i-1][j] + gap
			ins := f[i][j-1] + gap
			f[i][j] = max3(match, del, ins)
		}
	}

	// Align
	z := make([]byte, 0)
	w := make([]byte, 0)

	i := len(x) - 1
	j := len(y) - 1

	for i > 0 || j > 0 {
		if i > 0 && j > 0 && f[i][j] == f[i-1][j-1]+similarity(x[i], y[j]) { // match
			z = append([]byte{x[i]}, z...)
			w = append([]byte{y[j]}, w...)
			i--
			j--
		} else if i > 0 && f[i][j] == f[i-1][j]+gap { // del
			z = append([]byte{x[i]}, z...)
			w = append([]byte{'-'}, w...)
			i--
		} else { // ins
			z = append([]byte{'-'}, z...)
			w = append([]byte{y[j]}, w...)
			j--
		}
	}

	if i == 0 && j == 0 {
		z = append([]byte{x[i]}, z...)
		w = append([]byte{y[j]}, w...)
	}

	return z, w
}

func NeedlemanWunschReuseWords(x []string, y []string, gap int, similarity func(string, string) int) ([]string, []string, int) {
	score := 0
	// Build grid
	if f == nil {
		f = make([][]int, maxSize)
		for i := 0; i < maxSize; i++ {
			f[i] = make([]int, maxSize)
		}
	}

	for i := 0; i < len(x); i++ {
		f[i][0] = gap * i
	}
	for j := 0; j < len(y); j++ {
		f[0][j] = gap * j
	}
	for i := 1; i < len(x); i++ {
		for j := 1; j < len(y); j++ {
			match := f[i-1][j-1] + similarity(x[i], y[j])
			del := f[i-1][j] + gap
			ins := f[i][j-1] + gap
			f[i][j] = max3(match, del, ins)
		}
	}

	// Align
	z := make([]string, 0)
	w := make([]string, 0)

	i := len(x) - 1
	j := len(y) - 1

	for i > 0 || j > 0 {
		if i > 0 && j > 0 && f[i][j] == f[i-1][j-1]+similarity(x[i], y[j]) { // match
			score += similarity(x[i], y[j])

			//z = append([]string{x[i]}, z...)
			//w = append([]string{y[j]}, w...)
			i--
			j--
		} else if i > 0 && f[i][j] == f[i-1][j]+gap { // del
			score += gap
			//z = append([]string{x[i]}, z...)
			//w = append([]string{"-"}, w...)
			i--
		} else { // ins
			score += gap
			//z = append([]string{"-"}, z...)
			//w = append([]string{y[j]}, w...)
			j--
		}
	}

	if i == 0 && j == 0 {
		score += similarity(x[i], y[j])
		//z = append([]string{x[i]}, z...)
		//w = append([]string{y[j]}, w...)
	}

	return z, w, score
}
