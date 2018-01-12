package sequencing

type Sequence interface {
	Val(int) interface{}
	Len() int
}
