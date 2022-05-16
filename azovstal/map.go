package azovstal

import (
	"errors"
	"fmt"
)

type Map struct {
	mm   []byte
	size int
}

func (g *Map) Print() {
	for i := 0; i < g.size; i++ {
		for j := 0; j < g.size; j++ {
			c := g.Get(i, j)
			if c > 0 {
				fmt.Print(c)
				fmt.Print(" ")
			} else {
				fmt.Print("◻ ")
			}
		}
		fmt.Println()
	}
}
func NewMap(s int) (g *Map) {
	g = &Map{
		size: s,
		mm:   make([]byte, s*s),
	}
	return
}

func (g *Map) Get(x, y int) byte {
	return g.mm[y*g.size+x]
}

func (g *Map) Set(x, y int, b byte) (err error) {
	if x < g.size && y < g.size {
		g.mm[y*g.size+x] = b
	} else {
		return errors.New("location out of range")
	}
	return
}
