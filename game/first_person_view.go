package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	bgImg         *ebiten.Image
	entityFarImg  *ebiten.Image
	entityMidImg  *ebiten.Image
	entityNearImg *ebiten.Image
	wallImgs      []*ebiten.Image
)

func init() {
	bgImg, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("../assets/%s/bg.png", testSkin), ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	entityFarImg, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("../assets/entity/%sFar.png", testEnemy), ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	entityMidImg, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("../assets/entity/%sMid.png", testEnemy), ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	entityNearImg, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("../assets/entity/%sNear.png", testEnemy), ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		path := fmt.Sprintf("../assets/%s/%d.png", testSkin, i)
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

	//order: ffll, ffrr, ffl, ffr, ff, fl, fr, f, l, r, x
	fovCells := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
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
	//TODO: Do not display current cell entity. Near enemy shown in combat mode only. Other entities blocking or pickups.
	fovCells[10] = getCell(g.player.x, g.player.y, g.gridMap)

	for i := range fovCells {
		cellOp := &ebiten.DrawImageOptions{}
		if fovCells[i] == 1 {
			v.DrawImage(wallImgs[i], cellOp)
		} else if fovCells[i] > 1 {
			if i == 2 {
				cellOp.GeoM.Translate(1, 20)
				v.DrawImage(entityFarImg, cellOp)
			}
			if i == 3 {
				cellOp.GeoM.Translate(41, 20)
				v.DrawImage(entityFarImg, cellOp)
			}
			if i == 4 {
				cellOp.GeoM.Translate(21, 20)
				v.DrawImage(entityFarImg, cellOp)
			}
			if i == 5 {
				cellOp.GeoM.Translate(-13, 15)
				v.DrawImage(entityMidImg, cellOp)
			}
			if i == 6 {
				cellOp.GeoM.Translate(47, 15)
				v.DrawImage(entityMidImg, cellOp)
			}
			if i == 7 {
				cellOp.GeoM.Translate(17, 15)
				v.DrawImage(entityMidImg, cellOp)
			}
			//TODO: Remove, see comment above...
			if i == 10 {
				v.DrawImage(entityNearImg, cellOp)
			}
		}
	}

	return v
}
