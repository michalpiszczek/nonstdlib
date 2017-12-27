package set

import (
	"testing"
)

func TestSet(t *testing.T) {
	s := NewHashSet()
	var _ Set = s

    ts := NewTreeSet()
    var _ Set = ts
}
