package azovstal

import (
	"fmt"
	"strconv"
	"testing"
)

func TestPiece2byte(t *testing.T) {
	x := Piece2uint16([3][3]bool{
		{true, true, false},
		{true, true, false},
		{false, true, false}})
	fmt.Println(strconv.FormatInt(int64(uint8(x)), 2))
}

func TestByte2piece(t *testing.T) {
	b, _ := strconv.ParseInt("01010110", 2, 8)
	x := Uint162piece(uint16(b))
	fmt.Println(x)
}

func TestCheckbyte(t *testing.T) {

	for _, v := range defaultPieceMap {
		g := NewGame(3, 1)
		fmt.Println(v)
		g.Set(1, 0, 0, Piece2uint16(v))
		g.m.Print()
		fmt.Println()
	}

}
