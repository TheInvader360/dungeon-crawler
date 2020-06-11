package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	bgImg    *ebiten.Image
	wallImgs []*ebiten.Image
)

func init() {
	bgImg, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("../assets/%s/bg.png", skin), ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		path := fmt.Sprintf("../assets/%s/%d.png", skin, i)
		wallImg, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
		wallImgs = append(wallImgs, wallImg)
	}
}

func renderFirstPersonView(g *Game, v *ebiten.Image) *ebiten.Image {
	bgOp := &ebiten.DrawImageOptions{}
	v.DrawImage(bgImg, bgOp)

	//order: ffll, ffrr, ffl, ffr, ff, fl, fr, f, l, r
	fovCells := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	switch g.player.dir {
	case north:
		fovCells[0] = getCell(g.player.x-2, g.player.y-2, g.gridMap)
		fovCells[1] = getCell(g.player.x+2, g.player.y-2, g.gridMap)
		fovCells[2] = getCell(g.player.x-1, g.player.y-2, g.gridMap)
		fovCells[3] = getCell(g.player.x+1, g.player.y-2, g.gridMap)
		fovCells[4] = getCell(g.player.x, g.player.y-2, g.gridMap)
		fovCells[5] = getCell(g.player.x-1, g.player.y-1, g.gridMap)
		fovCells[6] = getCell(g.player.x+1, g.player.y-1, g.gridMap)
		fovCells[7] = getCell(g.player.x, g.player.y-1, g.gridMap)
		fovCells[8] = getCell(g.player.x-1, g.player.y, g.gridMap)
		fovCells[9] = getCell(g.player.x+1, g.player.y, g.gridMap)
	case east:
		fovCells[0] = getCell(g.player.x+2, g.player.y-2, g.gridMap)
		fovCells[1] = getCell(g.player.x+2, g.player.y+2, g.gridMap)
		fovCells[2] = getCell(g.player.x+2, g.player.y-1, g.gridMap)
		fovCells[3] = getCell(g.player.x+2, g.player.y+1, g.gridMap)
		fovCells[4] = getCell(g.player.x+2, g.player.y, g.gridMap)
		fovCells[5] = getCell(g.player.x+1, g.player.y-1, g.gridMap)
		fovCells[6] = getCell(g.player.x+1, g.player.y+1, g.gridMap)
		fovCells[7] = getCell(g.player.x+1, g.player.y, g.gridMap)
		fovCells[8] = getCell(g.player.x, g.player.y-1, g.gridMap)
		fovCells[9] = getCell(g.player.x, g.player.y+1, g.gridMap)
	case south:
		fovCells[0] = getCell(g.player.x+2, g.player.y+2, g.gridMap)
		fovCells[1] = getCell(g.player.x-2, g.player.y+2, g.gridMap)
		fovCells[2] = getCell(g.player.x+1, g.player.y+2, g.gridMap)
		fovCells[3] = getCell(g.player.x-1, g.player.y+2, g.gridMap)
		fovCells[4] = getCell(g.player.x, g.player.y+2, g.gridMap)
		fovCells[5] = getCell(g.player.x+1, g.player.y+1, g.gridMap)
		fovCells[6] = getCell(g.player.x-1, g.player.y+1, g.gridMap)
		fovCells[7] = getCell(g.player.x, g.player.y+1, g.gridMap)
		fovCells[8] = getCell(g.player.x+1, g.player.y, g.gridMap)
		fovCells[9] = getCell(g.player.x-1, g.player.y, g.gridMap)
	case west:
		fovCells[0] = getCell(g.player.x-2, g.player.y+2, g.gridMap)
		fovCells[1] = getCell(g.player.x-2, g.player.y-2, g.gridMap)
		fovCells[2] = getCell(g.player.x-2, g.player.y+1, g.gridMap)
		fovCells[3] = getCell(g.player.x-2, g.player.y-1, g.gridMap)
		fovCells[4] = getCell(g.player.x-2, g.player.y, g.gridMap)
		fovCells[5] = getCell(g.player.x-1, g.player.y+1, g.gridMap)
		fovCells[6] = getCell(g.player.x-1, g.player.y-1, g.gridMap)
		fovCells[7] = getCell(g.player.x-1, g.player.y, g.gridMap)
		fovCells[8] = getCell(g.player.x, g.player.y+1, g.gridMap)
		fovCells[9] = getCell(g.player.x, g.player.y-1, g.gridMap)
	}

	for i := range fovCells {
		if fovCells[i] == 1 {
			wiOp := &ebiten.DrawImageOptions{}
			v.DrawImage(wallImgs[i], wiOp)
		}
	}

	return v
}
