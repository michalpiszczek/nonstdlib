package worklist

import (
	"testing"
)

func TestWorklist(t *testing.T) {
	q := NewQueue()
	s := NewStack()
	var _ WorkList = q
	var _ WorkList = s
}
