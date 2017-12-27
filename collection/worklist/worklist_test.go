package worklist

import (
	"testing"
)

func TestWorklist(t *testing.T) {
	q := NewQueue()
	var _ WorkList = q

    s := NewStack()
	var _ WorkList = s
}
