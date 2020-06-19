package dungeon

import (
	resfirstperson "github.com/TheInvader360/dungeon-crawler/res/firstperson/rgb" // 1bit/rgb
	"github.com/hajimehoshi/ebiten"
)

var (
	bgImg      *ebiten.Image
	crackImg   *ebiten.Image
	lockFarImg *ebiten.Image
	lockMidImg *ebiten.Image
	exitFarImg *ebiten.Image
	exitMidImg *ebiten.Image
	wallImgs   []*ebiten.Image
)

func init() {
	bgImg = EssentialNewImageFromEncoded(resfirstperson.Bg_png)
	crackImg = EssentialNewImageFromEncoded(resfirstperson.Crack_png)
	lockFarImg = EssentialNewImageFromEncoded(resfirstperson.LockFar_png)
	lockMidImg = EssentialNewImageFromEncoded(resfirstperson.LockMid_png)
	exitFarImg = EssentialNewImageFromEncoded(resfirstperson.ExitFar_png)
	exitMidImg = EssentialNewImageFromEncoded(resfirstperson.ExitMid_png)
	w0 := EssentialNewImageFromEncoded(resfirstperson.Wall0_png)
	w1 := EssentialNewImageFromEncoded(resfirstperson.Wall1_png)
	w2 := EssentialNewImageFromEncoded(resfirstperson.Wall2_png)
	w3 := EssentialNewImageFromEncoded(resfirstperson.Wall3_png)
	w4 := EssentialNewImageFromEncoded(resfirstperson.Wall4_png)
	w5 := EssentialNewImageFromEncoded(resfirstperson.Wall5_png)
	w6 := EssentialNewImageFromEncoded(resfirstperson.Wall6_png)
	w7 := EssentialNewImageFromEncoded(resfirstperson.Wall7_png)
	w8 := EssentialNewImageFromEncoded(resfirstperson.Wall8_png)
	w9 := EssentialNewImageFromEncoded(resfirstperson.Wall9_png)
	wallImgs = []*ebiten.Image{w0, w1, w2, w3, w4, w5, w6, w7, w8, w9}
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
			//cracks/locks/exits will always be in the middle of the fov
			//cracks aren't visible from afar (can appear anywhere on map)
			//locks/exits are visible from afar (restricted to corridors)
			if i == 4 {
				if fovCells[i].wall == locked {
					v.DrawImage(lockFarImg, cellOp)
				}
				if fovCells[i].wall == exit {
					v.DrawImage(exitFarImg, cellOp)
				}
			}
			if i == 7 {
				if fovCells[i].wall == breakable {
					v.DrawImage(crackImg, cellOp)
				}
				if fovCells[i].wall == locked {
					v.DrawImage(lockMidImg, cellOp)
				}
				if fovCells[i].wall == exit {
					v.DrawImage(exitMidImg, cellOp)
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
