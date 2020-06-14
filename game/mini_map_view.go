package main

import (
	"math"

	"github.com/hajimehoshi/ebiten"
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
	enemyImg     *ebiten.Image
)

func init() {
	offsetX = screenWidth - cellSize*viewportSize
	offsetY = screenHeight - cellSize*viewportSize
	blockedImg = essentialNewImageFromFile("../assets/mini_map/blocked.png")
	unblockedImg = essentialNewImageFromFile("../assets/mini_map/unblocked.png")
	playerImg = essentialNewImageFromFile("../assets/mini_map/player.png")
	enemyImg = essentialNewImageFromFile("../assets/mini_map/enemy.png")
}

func renderMiniMapView(p player, gm [][]cell, v *ebiten.Image) *ebiten.Image {
	viewportCells := getCells(p.x-((viewportSize-1)/2), p.y-((viewportSize-1)/2), viewportSize, viewportSize, gm)
	for y := range viewportCells {
		for x := range viewportCells[y] {
			cellOp := &ebiten.DrawImageOptions{}
			cellOp.GeoM.Translate(float64(x*cellSize+offsetX), float64(y*cellSize+offsetY))
			v.DrawImage(unblockedImg, cellOp)
			if viewportCells[y][x].wall {
				v.DrawImage(blockedImg, cellOp)
			}
			if viewportCells[y][x].enemy != nil {
				v.DrawImage(enemyImg, cellOp)
			}
		}
	}

	playerOp := &ebiten.DrawImageOptions{}
	w, h := playerImg.Size()
	playerOp.GeoM.Translate(-float64(w)/2, -float64(h)/2) //move centre of image to origin before rotating
	playerOp.GeoM.Rotate(float64(p.dir*90) * 2 * math.Pi / 360)
	playerOp.GeoM.Translate(float64(w)/2, float64(h)/2) //move centre of image back to the starting point
	playerOp.GeoM.Translate(float64(cellSize*((viewportSize-1)/2)+offsetX), float64(cellSize*((viewportSize-1)/2)+offsetY))
	v.DrawImage(playerImg, playerOp)

	return v
}
