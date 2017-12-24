package dictionary

import (
	"testing"
)

func TestDictionary(t *testing.T) {
	h := NewHashMap()
	var _ Dictionary = h

	tm := NewTreeMap()
	var _ Dictionary = tm
}
