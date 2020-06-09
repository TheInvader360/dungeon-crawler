package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	bgImg *ebiten.Image
)

func init() {
	bgImg, _, err = ebitenutil.NewImageFromFile("../assets/rgb/bg.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

func renderFirstPersonView(g *Game, v *ebiten.Image) *ebiten.Image {
	bgOp := &ebiten.DrawImageOptions{}
	v.DrawImage(bgImg, bgOp)

	return v
}
