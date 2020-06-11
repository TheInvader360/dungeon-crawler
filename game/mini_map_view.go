package main

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	cellSize     = 5
	viewportSize = 9
)

var (
	offsetX      int
	offsetY      int
	blockedImg   *ebiten.Image
	unblockedImg *ebiten.Image
	playerImg    *ebiten.Image
)

func init() {
	offsetX = screenWidth - cellSize*viewportSize
	offsetY = screenHeight - cellSize*viewportSize
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
	viewportCells := getCells(g.player.x-((viewportSize-1)/2), g.player.y-((viewportSize-1)/2), viewportSize, viewportSize, g.gridMap)
	for y := range viewportCells {
		for x := range viewportCells[y] {
			cellOp := &ebiten.DrawImageOptions{}
			cellOp.GeoM.Translate(float64(x*cellSize+offsetX), float64(y*cellSize+offsetY))
			v.DrawImage(unblockedImg, cellOp)
			if viewportCells[y][x] == 1 {
				v.DrawImage(blockedImg, cellOp)
			}
		}
	}

	playerOp := &ebiten.DrawImageOptions{}
	w, h := playerImg.Size()
	playerOp.GeoM.Translate(-float64(w)/2, -float64(h)/2) //move centre of image to origin before rotating
	playerOp.GeoM.Rotate(float64(g.player.dir*90) * 2 * math.Pi / 360)
	playerOp.GeoM.Translate(float64(w)/2, float64(h)/2) //move centre of image back to the starting point
	playerOp.GeoM.Translate(float64(cellSize*((viewportSize-1)/2)+offsetX), float64(cellSize*((viewportSize-1)/2)+offsetY))
	v.DrawImage(playerImg, playerOp)

	return v
}
