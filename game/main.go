package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const (
	screenWidth  = 80
	screenHeight = 60
)

var (
	firstPersonImg *ebiten.Image
	miniMapImg     *ebiten.Image
	err            error
)

func dungeonCrawler() *Game {
	return &Game{
		gridMap: newGridMap(),
		player:  newPlayer(),
	}
}

func init() {
	firstPersonImg, _ = ebiten.NewImage(screenWidth, screenHeight, ebiten.FilterNearest)
	miniMapImg, _ = ebiten.NewImage(screenWidth, screenHeight, ebiten.FilterNearest)
}

//Game ...
type Game struct {
	gridMap [][]int
	player  player
}

//Update ...
func (g *Game) Update(screen *ebiten.Image) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		x, y := g.player.getCoordInFront()
		if getCell(x, y, g.gridMap) != 1 {
			g.player.moveTo(x, y)
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		g.player.turnLeft()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		g.player.turnRight()
	}
	return nil
}

//Draw ...
func (g *Game) Draw(screen *ebiten.Image) {
	if !ebiten.IsKeyPressed(ebiten.KeyZ) {
		firstPersonImg = renderFirstPersonView(g, firstPersonImg)
		fpOp := &ebiten.DrawImageOptions{}
		screen.DrawImage(firstPersonImg, fpOp)
	}
	if !ebiten.IsKeyPressed(ebiten.KeyX) {
		miniMapImg = renderMiniMapView(g, miniMapImg)
		mmOp := &ebiten.DrawImageOptions{}
		screen.DrawImage(miniMapImg, mmOp)
	}
	if !ebiten.IsKeyPressed(ebiten.KeyC) {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("%s", g.player))
	}
}

//Layout ...
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*10, screenHeight*10)
	ebiten.SetWindowTitle("Dungeon Crawler")
	if err := ebiten.RunGame(dungeonCrawler()); err != nil {
		log.Fatal(err)
	}
}
