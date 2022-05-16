package azovstal

import "errors"

type Game struct {
	m       Map
	players uint8
}

func NewGame(s int, playerCnt uint8) (g *Game) {
	g = &Game{
		m:       *NewMap(s),
		players: playerCnt,
	}
	return
}
func (g *Game) Set(playerIdx uint8, x, y int, piece uint16) (err error) {

	vs := Uint162piece(piece)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			m, n := x+i, y+j
			if g.m.Get(m, n) > 0 {
				return errors.New("location already exist")
			}
			if vs[i][j] {
				err = g.m.Set(m, n, playerIdx)
				g.m.Set(m, n, 0)
				if err != nil {
					return err
				}
			}
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			m, n := x+i, y+j
			if vs[i][j] {
				g.m.Set(m, n, playerIdx)
			}
		}
	}
	return
}
