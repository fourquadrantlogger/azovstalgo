package azovstal

var defaultPieceMap = [][3][3]bool{
	//1   o
	[3][3]bool{
		{true, false, false},
		{false, false, false},
		{false, false, false},
	},

	//2 oo
	[3][3]bool{
		{true, true, false},
		{false, false, false},
		{false, false, false},
	}, //oo

	//2 8
	[3][3]bool{
		{true, false, false},
		{true, false, false},
		{false, false, false},
	},

	//3 ooo
	[3][3]bool{
		{true, true, true},
		{false, false, false},
		{false, false, false},
	},

	//3 ミ
	[3][3]bool{
		{true, false, false},
		{true, false, false},
		{true, false, false},
	},

	//3 ∟
	[3][3]bool{
		{true, true, false},
		{true, false, false},
		{false, false, false},
	},
	[3][3]bool{
		{true, true, false},
		{false, true, false},
		{false, false, false},
	},
	[3][3]bool{
		{true, false, false},
		{true, true, false},
		{false, false, false},
	},
	[3][3]bool{
		{false, true, false},
		{true, true, false},
		{false, false, false},
	},

	//4
	[3][3]bool{
		{true, true, false},
		{true, true, false},
		{false, false, false},
	}, //L
	[3][3]bool{
		{true, true, false},
		{true, false, false},
		{true, false, false},
	}, //L
	[3][3]bool{
		{true, false, false},
		{true, true, false},
		{true, false, false},
	}, //L
	[3][3]bool{
		{true, false, false},
		{true, false, false},
		{true, true, false},
	}, //L

	[3][3]bool{
		{true, true, true},
		{true, false, false},
		{false, false, false},
	}, //L
	[3][3]bool{
		{true, true, true},
		{false, true, false},
		{false, false, false},
	}, //L
	[3][3]bool{
		{true, true, true},
		{false, false, true},
		{false, false, false},
	}, //L
}

func Uint162piece(piece uint16) (vs [3][3]bool) {
	for i := 0; i < 9; i++ {
		var x uint16

		x = uint16(1 << i)

		if x&piece == x {
			vs[i/3][i%3] = true
		}
	}
	return
}
func Piece2uint16(vs [3][3]bool) (b uint16) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			idx := i*3 + j

			var x uint16 = 1 << idx
			if vs[i][j] {
				b = b | x
			}
		}
	}
	return
}
