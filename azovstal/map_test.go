package azovstal

import "testing"

func TestMap_Print(t *testing.T) {
	m := NewMap(4)

	m.Set(0, 3, 4)
	m.Set(1, 3, 2)

	m.Print()
}
