package internal

type Number interface {
	int8 | byte | int16 | uint16 | int | rune | uint32 | int64 | uint64 | float32 | float64
}

type EulerScalar[T Number, R any] struct {
	Input    T
	Expected R
}

type EulerGraph[T Number, R any] struct {
	Input    []T
	Expected R
}
