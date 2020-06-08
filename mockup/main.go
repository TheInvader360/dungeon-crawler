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
	avatar *ebiten.Image
	bg     *ebiten.Image
	cells  []cell
	err    error
)

type cell struct {
	blocked bool
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
	for i := 0; i < 12; i++ {
		cell := cell{}
		if i == 3 || i == 4 || i == 7 || i == 11 {
			cell.blocked = true
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
	return func(screen *ebiten.Image) error {
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			for i, cell := range cells {
				if float64(x) > cell.btnX && float64(x) < cell.btnX+cellSize && float64(y) > cell.btnY && float64(y) < cell.btnY+cellSize {
					cells[i].blocked = !cells[i].blocked
				}
			}
		}
		if ebiten.IsDrawingSkipped() {
			return nil
		}
		screen.Fill(colornames.Grey)
		bgOp := &ebiten.DrawImageOptions{}
		//bgOp.GeoM.Scale(0.5, 0.5)
		screen.DrawImage(bg, bgOp)
		for _, cell := range cells {
			viewOp := &ebiten.DrawImageOptions{}
			//viewOp.GeoM.Scale(0.5, 0.5)
			if cell.blocked {
				screen.DrawImage(cell.view, viewOp)
			}
			btnOp := &ebiten.DrawImageOptions{}
			btnOp.GeoM.Translate(cell.btnX, cell.btnY)
			if cell.blocked {
				screen.DrawImage(cellBlockedImage, btnOp)
			} else {
				screen.DrawImage(cellUnblockedImage, btnOp)
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
