package azovstal

import "testing"

func TestGame_Set(t *testing.T) {
	g := NewGame(9, 2)
	g.Set(1, 0, 0, Piece2uint16([3][3]bool{
		{true, true, false},
		{false, false, false},
		{false, false, false},
	}))
	g.m.Print()
}
