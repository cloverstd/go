package jsoniter

import "fmt"

var MaxDepth = 1024

type MaxDepthError struct {
	depth int
}

func (e MaxDepthError) Error() string {
	return fmt.Sprintf("exceeding maximum depth %d", e.depth)
}

func IsMaxDepthError(err error) bool {
	_, ok := err.(MaxDepthError)
	return ok
}

func newMaxDepthError(depth int) error {
	return MaxDepthError{
		depth,
	}
}
