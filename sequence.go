package sequencing

type Comparable interface {
	Equal(comparable Comparable) bool
}

type Sequence interface {
	Val(int) Comparable
	Len() int
}
