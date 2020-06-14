package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

type enemy struct {
	farImg  *ebiten.Image
	midImg  *ebiten.Image
	nearImg *ebiten.Image
}

var enemies []enemy

func init() {
	for i := 0; i < 7; i++ {
		f := essentialNewImageFromFile(fmt.Sprintf("../assets/entity/enemy%dFar.png", i))
		m := essentialNewImageFromFile(fmt.Sprintf("../assets/entity/enemy%dMid.png", i))
		n := essentialNewImageFromFile(fmt.Sprintf("../assets/entity/enemy%dNear.png", i))
		e := enemy{farImg: f, midImg: m, nearImg: n}
		enemies = append(enemies, e)
	}
}
