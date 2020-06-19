package dungeon

import (
	"math"

	resminimap "github.com/TheInvader360/dungeon-crawler/res/minimap"
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
	exitImg      *ebiten.Image
	keyImg       *ebiten.Image
	goldImg      *ebiten.Image
	potionImg    *ebiten.Image
	enemyImg     *ebiten.Image
	playerImg    *ebiten.Image
)

func init() {
	offsetX = ScreenWidth - cellSize*viewportSize
	offsetY = ScreenHeight - cellSize*viewportSize
	noneImg = EssentialNewImageFromEncoded(resminimap.None_png)
	solidImg = EssentialNewImageFromEncoded(resminimap.Solid_png)
	breakableImg = EssentialNewImageFromEncoded(resminimap.Breakable_png)
	lockedImg = EssentialNewImageFromEncoded(resminimap.Locked_png)
	exitImg = EssentialNewImageFromEncoded(resminimap.Exit_png)
	keyImg = EssentialNewImageFromEncoded(resminimap.Key_png)
	goldImg = EssentialNewImageFromEncoded(resminimap.Gold_png)
	potionImg = EssentialNewImageFromEncoded(resminimap.Potion_png)
	enemyImg = EssentialNewImageFromEncoded(resminimap.Enemy_png)
	playerImg = EssentialNewImageFromEncoded(resminimap.Player_png)
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
			if viewportCells[y][x].wall == exit {
				v.DrawImage(exitImg, cellOp)
			}
			if viewportCells[y][x].collectible == &key {
				v.DrawImage(keyImg, cellOp)
			}
			if viewportCells[y][x].collectible == &gold {
				v.DrawImage(goldImg, cellOp)
			}
			if viewportCells[y][x].collectible == &potion {
				v.DrawImage(potionImg, cellOp)
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
