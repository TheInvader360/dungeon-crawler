package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/inpututil"

	"golang.org/x/image/colornames"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 80
	screenHeight = 108
	cellSize     = 16
)

var (
	avatar     *ebiten.Image
	bg         *ebiten.Image
	entityFar  *ebiten.Image
	entityMid  *ebiten.Image
	entityNear *ebiten.Image
	cells      []cell
	err        error
)

type cell struct {
	blocked bool
	entity  bool
	btnX    float64
	btnY    float64
	view    *ebiten.Image
}

func initGame() func(*ebiten.Image) error {
	avatar, _, err = ebitenutil.NewImageFromFile("../assets/avatar.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	bg, _, err = ebitenutil.NewImageFromFile("../assets/rgb/bg.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	entityFar, _, err = ebitenutil.NewImageFromFile("../assets/entityFar.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	entityMid, _, err = ebitenutil.NewImageFromFile("../assets/entityMid.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	entityNear, _, err = ebitenutil.NewImageFromFile("../assets/entityNear.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 12; i++ {
		cell := cell{}
		if i == 3 || i == 4 || i == 7 || i == 11 {
			cell.blocked = true
		}
		if i < 5 || i > 6 {
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
		} else if i < 10 {
			cell.btnY = screenHeight - 2*cellSize
			if i == 5 {
				cell.btnX = 0
			} else if i == 6 {
				cell.btnX = 4 * cellSize
			} else if i == 7 {
				cell.btnX = cellSize
			} else if i == 8 {
				cell.btnX = 3 * cellSize
			} else {
				cell.btnX = 2 * cellSize
			}
		} else {
			cell.btnY = screenHeight - cellSize
			if i == 10 {
				cell.btnX = cellSize
			} else {
				cell.btnX = 3 * cellSize
			}
		}
		var path = fmt.Sprintf("../assets/rgb/%d.png", i)
		var view, _, err = ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
		cell.view = view
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
				if i == 0 {
					entityOp.GeoM.Translate(-10, 20)
					screen.DrawImage(entityFar, entityOp)
				}
				if i == 2 {
					entityOp.GeoM.Translate(10, 20)
					screen.DrawImage(entityFar, entityOp)
				}
				if i == 4 {
					entityOp.GeoM.Translate(30, 20)
					screen.DrawImage(entityFar, entityOp)
				}
				if i == 3 {
					entityOp.GeoM.Translate(50, 20)
					screen.DrawImage(entityFar, entityOp)
				}
				if i == 1 {
					entityOp.GeoM.Translate(70, 20)
					screen.DrawImage(entityFar, entityOp)
				}
				//MID, L-R
				if i == 7 {
					entityOp.GeoM.Translate(-5, 15)
					screen.DrawImage(entityMid, entityOp)
				}
				if i == 9 {
					entityOp.GeoM.Translate(25, 15)
					screen.DrawImage(entityMid, entityOp)
				}
				if i == 8 {
					entityOp.GeoM.Translate(55, 15)
					screen.DrawImage(entityMid, entityOp)
				}
				//NEAR, L-R
				if i == 10 {
					entityOp.GeoM.Translate(-41, 3)
					screen.DrawImage(entityNear, entityOp)
				}
				if i == 11 {
					entityOp.GeoM.Translate(67, 3)
					screen.DrawImage(entityNear, entityOp)
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
		avatarOp := &ebiten.DrawImageOptions{}
		avatarOp.GeoM.Translate(2*cellSize, screenHeight-cellSize)
		screen.DrawImage(avatar, avatarOp)
		return nil
	}
}

func main() {
	if err := ebiten.Run(initGame(), screenWidth, screenHeight, 8, "Dungeon Crawler Mockup"); err != nil {
		log.Fatal(err)
	}
}
