package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

var (
	bgImg    *ebiten.Image
	crackImg *ebiten.Image
	lockImg  *ebiten.Image
	wallImgs []*ebiten.Image
)

func init() {
	bgImg = essentialNewImageFromFile(fmt.Sprintf("../assets/%s/bg.png", testSkin))
	crackImg = essentialNewImageFromFile(fmt.Sprintf("../assets/%s/crack.png", testSkin))
	lockImg = essentialNewImageFromFile(fmt.Sprintf("../assets/%s/lock.png", testSkin))
	for i := 0; i < 10; i++ {
		wallImg := essentialNewImageFromFile(fmt.Sprintf("../assets/%s/%d.png", testSkin, i))
		wallImgs = append(wallImgs, wallImg)
	}
}

func renderFirstPersonView(p player, gm [][]cell, v *ebiten.Image) *ebiten.Image {
	bgOp := &ebiten.DrawImageOptions{}
	v.DrawImage(bgImg, bgOp)

	//order: ffll, ffrr, ffl, ffr, ff, fl, fr, f, l, r
	fovCells := make([]cell, 10)
	switch p.dir {
	case north:
		fovCells[0] = getCell(p.x-2, p.y-2, gm)
		fovCells[1] = getCell(p.x+2, p.y-2, gm)
		fovCells[2] = getCell(p.x-1, p.y-2, gm)
		fovCells[3] = getCell(p.x+1, p.y-2, gm)
		fovCells[4] = getCell(p.x, p.y-2, gm)
		fovCells[5] = getCell(p.x-1, p.y-1, gm)
		fovCells[6] = getCell(p.x+1, p.y-1, gm)
		fovCells[7] = getCell(p.x, p.y-1, gm)
		fovCells[8] = getCell(p.x-1, p.y, gm)
		fovCells[9] = getCell(p.x+1, p.y, gm)
	case east:
		fovCells[0] = getCell(p.x+2, p.y-2, gm)
		fovCells[1] = getCell(p.x+2, p.y+2, gm)
		fovCells[2] = getCell(p.x+2, p.y-1, gm)
		fovCells[3] = getCell(p.x+2, p.y+1, gm)
		fovCells[4] = getCell(p.x+2, p.y, gm)
		fovCells[5] = getCell(p.x+1, p.y-1, gm)
		fovCells[6] = getCell(p.x+1, p.y+1, gm)
		fovCells[7] = getCell(p.x+1, p.y, gm)
		fovCells[8] = getCell(p.x, p.y-1, gm)
		fovCells[9] = getCell(p.x, p.y+1, gm)
	case south:
		fovCells[0] = getCell(p.x+2, p.y+2, gm)
		fovCells[1] = getCell(p.x-2, p.y+2, gm)
		fovCells[2] = getCell(p.x+1, p.y+2, gm)
		fovCells[3] = getCell(p.x-1, p.y+2, gm)
		fovCells[4] = getCell(p.x, p.y+2, gm)
		fovCells[5] = getCell(p.x+1, p.y+1, gm)
		fovCells[6] = getCell(p.x-1, p.y+1, gm)
		fovCells[7] = getCell(p.x, p.y+1, gm)
		fovCells[8] = getCell(p.x+1, p.y, gm)
		fovCells[9] = getCell(p.x-1, p.y, gm)
	case west:
		fovCells[0] = getCell(p.x-2, p.y+2, gm)
		fovCells[1] = getCell(p.x-2, p.y-2, gm)
		fovCells[2] = getCell(p.x-2, p.y+1, gm)
		fovCells[3] = getCell(p.x-2, p.y-1, gm)
		fovCells[4] = getCell(p.x-2, p.y, gm)
		fovCells[5] = getCell(p.x-1, p.y+1, gm)
		fovCells[6] = getCell(p.x-1, p.y-1, gm)
		fovCells[7] = getCell(p.x-1, p.y, gm)
		fovCells[8] = getCell(p.x, p.y+1, gm)
		fovCells[9] = getCell(p.x, p.y-1, gm)
	}

	for i := range fovCells {
		cellOp := &ebiten.DrawImageOptions{}
		if fovCells[i].wall != none {
			v.DrawImage(wallImgs[i], cellOp)
			if i == 7 {
				//only draw cracks/locks if immediately in front of player
				if fovCells[i].wall == breakable {
					v.DrawImage(crackImg, cellOp)
				}
				if fovCells[i].wall == locked {
					v.DrawImage(lockImg, cellOp)
				}
			}
		}
		if fovCells[i].collectible != nil {
			//order: ffl, ffr, ff, fl, fr, f
			if i == 2 {
				cellOp.GeoM.Translate(1, 20)
				v.DrawImage(fovCells[i].collectible.farImg, cellOp)
			}
			if i == 3 {
				cellOp.GeoM.Translate(41, 20)
				v.DrawImage(fovCells[i].collectible.farImg, cellOp)
			}
			if i == 4 {
				cellOp.GeoM.Translate(21, 20)
				v.DrawImage(fovCells[i].collectible.farImg, cellOp)
			}
			if i == 5 {
				cellOp.GeoM.Translate(-13, 15)
				v.DrawImage(fovCells[i].collectible.midImg, cellOp)
			}
			if i == 6 {
				cellOp.GeoM.Translate(47, 15)
				v.DrawImage(fovCells[i].collectible.midImg, cellOp)
			}
			if i == 7 {
				cellOp.GeoM.Translate(17, 15)
				v.DrawImage(fovCells[i].collectible.midImg, cellOp)
			}
		}
		if fovCells[i].enemy != nil {
			//order: ffl, ffr, ff, fl, fr, f
			if i == 2 {
				cellOp.GeoM.Translate(1, 20)
				v.DrawImage(fovCells[i].enemy.farImg, cellOp)
			}
			if i == 3 {
				cellOp.GeoM.Translate(41, 20)
				v.DrawImage(fovCells[i].enemy.farImg, cellOp)
			}
			if i == 4 {
				cellOp.GeoM.Translate(21, 20)
				v.DrawImage(fovCells[i].enemy.farImg, cellOp)
			}
			if i == 5 {
				cellOp.GeoM.Translate(-13, 15)
				v.DrawImage(fovCells[i].enemy.midImg, cellOp)
			}
			if i == 6 {
				cellOp.GeoM.Translate(47, 15)
				v.DrawImage(fovCells[i].enemy.midImg, cellOp)
			}
			if i == 7 {
				cellOp.GeoM.Translate(17, 15)
				v.DrawImage(fovCells[i].enemy.midImg, cellOp)
			}
		}
	}

	return v
}
