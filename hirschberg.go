package sequencing

func rev(s []byte) []byte {
	ret := make([]byte, 0)
	for i := len(s) - 1; i >= 0; i-- {
		ret = append(ret, s[i])
	}
	return ret
}

func revInt(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func insert(a byte) int {
	return -2
}
func del(a byte) int {
	return -2
}
func sub(a byte, b byte) int {
	if a == b {
		return 2
	}
	return -1
}

func nwScore(x []byte, y []byte) []int {
	//fmt.Printf("NWS on %v and %v\n", string(x), string(y))
	row1 := make([]int, len(y)+1)
	row2 := make([]int, len(y)+1)

	for j := 0; j < len(y); j++ {
		row1[j+1] = insert(y[j])
	}

	for i := 1; i < len(x)+1; i++ {
		row1, row2 = row2, row1
		row2[0] = row1[0] + del(x[i-1])

		for j := 1; j < len(y)+1; j++ {
			scoreSub := row1[j-1] + sub(x[i-1], y[j-1])
			scoreDel := row1[j] + del(x[i-1])
			scoreIns := row2[j-1] + insert(y[j-1])
			row2[j] = max3(scoreSub, scoreDel, scoreIns)
		}
	}
	//fmt.Printf("%v\n", score)
	return row2
}

func argmax(a []int, b []int) int {
	if len(a) != len(b) {
		panic("Arrays have different length!")
	}
	index := 0
	max := a[0] + b[0]
	for i := 1; i < len(a); i++ {
		if a[i]+b[i] >= max {
			max = a[i] + b[i]
			index = i
		}
	}
	return index
}

func Hirschberg(x []byte, y []byte) ([]byte, []byte) {
	//fmt.Printf("H on %s and %s\n", string(x), string(y))
	var z, w []byte

	revy := rev(y)
	if len(x) == 0 || len(y) == 0 {
		z = make([]byte, 0)
		w = make([]byte, 0)
	}
	if len(x) == 0 {
		for i := 0; i < len(y); i++ {
			z = append(z, '-')
			w = append(w, y[i])
		}
		return z, w
	}
	if len(y) == 0 {
		for i := 0; i < len(x); i++ {
			z = append(z, x[i])
			w = append(w, '-')
		}
		return z, w
	}
	if len(x) == 1 || len(y) == 1 {
		z, w, _ = NeedlemanWunsch(x, y, -2, sub)
		return z, w
	}

	xlen := len(x)
	xmid := xlen / 2
	ScoreL := nwScore(x[0:xmid], y)
	ScoreR := nwScore(rev(x[xmid:]), revy)
	//fmt.Printf("L %v R %v\n", ScoreL, ScoreR)
	ymid := argmax(ScoreL, revInt(ScoreR))
	//fmt.Printf("y %s\n", string(y))
	//fmt.Printf("Max %d\n", ymid)
	z1, w1 := Hirschberg(x[:xmid], y[:ymid])
	z2, w2 := Hirschberg(x[xmid:], y[ymid:])
	return append(z1, z2...), append(w1, w2...)
}
