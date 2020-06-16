package main

import (
	"github.com/hajimehoshi/ebiten"
)

type collectible struct {
	farImg *ebiten.Image
	midImg *ebiten.Image
}

var key collectible
var gold collectible
var potion collectible

func init() {
	kf := essentialNewImageFromFile("../assets/entity/keyFar.png")
	km := essentialNewImageFromFile("../assets/entity/keyMid.png")
	key = collectible{farImg: kf, midImg: km}

	gf := essentialNewImageFromFile("../assets/entity/goldFar.png")
	gm := essentialNewImageFromFile("../assets/entity/goldMid.png")
	gold = collectible{farImg: gf, midImg: gm}

	pf := essentialNewImageFromFile("../assets/entity/potionFar.png")
	pm := essentialNewImageFromFile("../assets/entity/potionMid.png")
	potion = collectible{farImg: pf, midImg: pm}
}
