package main

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	cellSize = 5
)

var (
	offsetX      int
	offsetY      int
	blockedImg   *ebiten.Image
	unblockedImg *ebiten.Image
	playerImg    *ebiten.Image
)

func init() {
	offsetX = screenWidth - cellSize*7
	offsetY = 0
	blockedImg, _, err = ebitenutil.NewImageFromFile("../assets/blocked.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	unblockedImg, _, err = ebitenutil.NewImageFromFile("../assets/unblocked.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	playerImg, _, err = ebitenutil.NewImageFromFile("../assets/player.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

func renderMiniMapView(g *Game, v *ebiten.Image) *ebiten.Image {
	for y := range g.gridMap {
		for x := range g.gridMap[y] {
			cellOp := &ebiten.DrawImageOptions{}
			cellOp.GeoM.Translate(float64(offsetX+x*cellSize), float64(offsetY+y*cellSize))
			if getCell(x, y, g.gridMap) == 0 {
				v.DrawImage(unblockedImg, cellOp)
			}
			if getCell(x, y, g.gridMap) == 1 {
				v.DrawImage(blockedImg, cellOp)
			}
		}
	}

	playerOp := &ebiten.DrawImageOptions{}
	w, h := playerImg.Size()
	playerOp.GeoM.Translate(-float64(w)/2, -float64(h)/2) //move centre of image to origin before rotating
	playerOp.GeoM.Rotate(float64(g.player.dir*90) * 2 * math.Pi / 360)
	playerOp.GeoM.Translate(float64(offsetX+g.player.x*cellSize+w/2), float64(offsetY+g.player.y*cellSize+h/2))
	v.DrawImage(playerImg, playerOp)

	return v
}
