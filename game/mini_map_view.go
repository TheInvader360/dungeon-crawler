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
	noneImg      *ebiten.Image
	solidImg     *ebiten.Image
	breakableImg *ebiten.Image
	lockedImg    *ebiten.Image
	enemyImg     *ebiten.Image
	playerImg    *ebiten.Image
)

func init() {
	offsetX = screenWidth - cellSize*viewportSize
	offsetY = screenHeight - cellSize*viewportSize
	noneImg = essentialNewImageFromFile("../assets/mini_map/none.png")
	solidImg = essentialNewImageFromFile("../assets/mini_map/solid.png")
	breakableImg = essentialNewImageFromFile("../assets/mini_map/breakable.png")
	lockedImg = essentialNewImageFromFile("../assets/mini_map/locked.png")
	enemyImg = essentialNewImageFromFile("../assets/mini_map/enemy.png")
	playerImg = essentialNewImageFromFile("../assets/mini_map/player.png")
}

func renderMiniMapView(p player, gm [][]cell, v *ebiten.Image) *ebiten.Image {
	viewportCells := getCells(p.x-((viewportSize-1)/2), p.y-((viewportSize-1)/2), viewportSize, viewportSize, gm)
	for y := range viewportCells {
		for x := range viewportCells[y] {
			cellOp := &ebiten.DrawImageOptions{}
			cellOp.GeoM.Translate(float64(x*cellSize+offsetX), float64(y*cellSize+offsetY))
			v.DrawImage(noneImg, cellOp)
			if viewportCells[y][x].wall == solid {
				v.DrawImage(solidImg, cellOp)
			}
			if viewportCells[y][x].wall == breakable {
				v.DrawImage(breakableImg, cellOp)
			}
			if viewportCells[y][x].wall == locked {
				v.DrawImage(lockedImg, cellOp)
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
