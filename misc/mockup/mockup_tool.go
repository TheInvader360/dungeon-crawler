package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/inpututil"

	"golang.org/x/image/colornames"

	"github.com/TheInvader360/dungeon-crawler/dungeon"
	resentity "github.com/TheInvader360/dungeon-crawler/res/entity"
	resfirstperson "github.com/TheInvader360/dungeon-crawler/res/firstperson/rgb" // 1bit/rgb
	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 60
	screenHeight = 96
	cellSize     = 12
)

var (
	bg        *ebiten.Image
	entityFar *ebiten.Image
	entityMid *ebiten.Image
	cells     []cell
	err       error
)

type cell struct {
	blocked bool
	entity  bool
	btnX    float64
	btnY    float64
	view    *ebiten.Image
}

func initGame() func(*ebiten.Image) error {
	bg := dungeon.EssentialNewImageFromEncoded(resfirstperson.Bg_png)
	//entityFar := dungeon.EssentialNewImageFromFile("../../res/entity/potionFar.png")
	//entityMid := dungeon.EssentialNewImageFromFile("../../res/entity/potionMid.png")
	entityFar := dungeon.EssentialNewImageFromEncoded(resentity.EntityFar_png)
	entityMid := dungeon.EssentialNewImageFromEncoded(resentity.EntityMid_png)
	w0 := dungeon.EssentialNewImageFromEncoded(resfirstperson.Wall0_png)
	w1 := dungeon.EssentialNewImageFromEncoded(resfirstperson.Wall1_png)
	w2 := dungeon.EssentialNewImageFromEncoded(resfirstperson.Wall2_png)
	w3 := dungeon.EssentialNewImageFromEncoded(resfirstperson.Wall3_png)
	w4 := dungeon.EssentialNewImageFromEncoded(resfirstperson.Wall4_png)
	w5 := dungeon.EssentialNewImageFromEncoded(resfirstperson.Wall5_png)
	w6 := dungeon.EssentialNewImageFromEncoded(resfirstperson.Wall6_png)
	w7 := dungeon.EssentialNewImageFromEncoded(resfirstperson.Wall7_png)
	w8 := dungeon.EssentialNewImageFromEncoded(resfirstperson.Wall8_png)
	w9 := dungeon.EssentialNewImageFromEncoded(resfirstperson.Wall9_png)
	wallImgs := []*ebiten.Image{w0, w1, w2, w3, w4, w5, w6, w7, w8, w9}

	for i := 0; i < 10; i++ {
		cell := cell{}
		if i == 2 || i == 3 || i == 5 || i == 6 || i == 8 || i == 9 {
			cell.blocked = true
		}
		if i > 1 && i < 8 {
			cell.entity = true
		}
		if i < 5 {
			cell.btnY = screenHeight - 3*cellSize
			if i == 0 {
				cell.btnX = 0
			} else if i == 1 {
				cell.btnX = 4 * cellSize
			} else if i == 2 {
				cell.btnX = cellSize
			} else if i == 3 {
				cell.btnX = 3 * cellSize
			} else {
				cell.btnX = 2 * cellSize
			}
		} else if i < 8 {
			cell.btnY = screenHeight - 2*cellSize
			if i == 5 {
				cell.btnX = cellSize
			} else if i == 6 {
				cell.btnX = 3 * cellSize
			} else {
				cell.btnX = 2 * cellSize
			}
		} else {
			cell.btnY = screenHeight - cellSize
			if i == 8 {
				cell.btnX = cellSize
			} else {
				cell.btnX = 3 * cellSize
			}
		}
		cell.view = wallImgs[i]
		cells = append(cells, cell)
	}

	cellBlockedImage, _ := ebiten.NewImage(cellSize, cellSize, ebiten.FilterDefault)
	cellBlockedImage.Fill(colornames.Sienna)
	cellUnblockedImage, _ := ebiten.NewImage(cellSize, cellSize, ebiten.FilterDefault)
	cellUnblockedImage.Fill(colornames.Antiquewhite)
	cellEntityImage, _ := ebiten.NewImage(cellSize/2, cellSize/2, ebiten.FilterDefault)
	cellEntityImage.Fill(colornames.Cornflowerblue)

	return func(screen *ebiten.Image) error {
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			for i, cell := range cells {
				if float64(x) > cell.btnX && float64(x) < cell.btnX+cellSize && float64(y) > cell.btnY && float64(y) < cell.btnY+cellSize {
					cells[i].blocked = !cells[i].blocked
				}
			}
		}
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonRight) {
			x, y := ebiten.CursorPosition()
			for i, cell := range cells {
				if float64(x) > cell.btnX && float64(x) < cell.btnX+cellSize && float64(y) > cell.btnY && float64(y) < cell.btnY+cellSize {
					cells[i].entity = !cells[i].entity
				}
			}
		}
		if ebiten.IsDrawingSkipped() {
			return nil
		}
		screen.Fill(colornames.Grey)
		bgOp := &ebiten.DrawImageOptions{}
		screen.DrawImage(bg, bgOp)
		for i, cell := range cells {
			if cell.entity {
				entityOp := &ebiten.DrawImageOptions{}
				//FAR, L-R
				if i == 2 {
					entityOp.GeoM.Translate(1, 20)
					screen.DrawImage(entityFar, entityOp)
				}
				if i == 4 {
					entityOp.GeoM.Translate(21, 20)
					screen.DrawImage(entityFar, entityOp)
				}
				if i == 3 {
					entityOp.GeoM.Translate(41, 20)
					screen.DrawImage(entityFar, entityOp)
				}
				//MID, L-R
				if i == 5 {
					entityOp.GeoM.Translate(-13, 15)
					screen.DrawImage(entityMid, entityOp)
				}
				if i == 7 {
					entityOp.GeoM.Translate(17, 15)
					screen.DrawImage(entityMid, entityOp)
				}
				if i == 6 {
					entityOp.GeoM.Translate(47, 15)
					screen.DrawImage(entityMid, entityOp)
				}
			}
			if cell.blocked {
				wallOp := &ebiten.DrawImageOptions{}
				screen.DrawImage(cell.view, wallOp)
			}
			btnOp := &ebiten.DrawImageOptions{}
			btnOp.GeoM.Translate(cell.btnX, cell.btnY)
			if cell.blocked {
				screen.DrawImage(cellBlockedImage, btnOp)
			} else {
				screen.DrawImage(cellUnblockedImage, btnOp)
			}
			if cell.entity {
				btnOp.GeoM.Translate(cellSize/4, cellSize/4)
				screen.DrawImage(cellEntityImage, btnOp)
			}
		}
		return nil
	}
}

func main() {
	if err := ebiten.Run(initGame(), screenWidth, screenHeight, 8, "First Person View Mockup"); err != nil {
		log.Fatal(err)
	}
}
