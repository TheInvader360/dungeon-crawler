package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const (
	screenWidth  = 60
	screenHeight = 60
	testSkin     = "rgb"    // 1bit / db16 / rgb
	testEnemy    = "enemy4" // entity / enemy0 / enemy1 / enemy2 / enemy3 / enemy4
)

var (
	firstPersonImg *ebiten.Image
	miniMapImg     *ebiten.Image
	combatImg      *ebiten.Image
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
	combatImg, _ = ebiten.NewImage(screenWidth, screenHeight, ebiten.FilterNearest)
}

//Game ...
type Game struct {
	gameState gameState
	gridMap   [][]int
	player    player
}

//Update ...
func (g *Game) Update(screen *ebiten.Image) error {
	switch g.gameState {
	case exploration:
		if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			x, y := g.player.getCoordInFront()
			target := getCell(x, y, g.gridMap)
			if target != 1 {
				g.player.moveTo(x, y)
			}
			if target > 1 {
				g.gameState = combat
			}
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
			g.player.turnLeft()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
			g.player.turnRight()
		}
	case combat:
		if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			setCell(g.player.x, g.player.y, 0, g.gridMap)
			g.gameState = exploration
		}
	}
	return nil
}

//Draw ...
func (g *Game) Draw(screen *ebiten.Image) {
	firstPersonImg = renderFirstPersonView(g.player, g.gridMap, firstPersonImg)
	fpOp := &ebiten.DrawImageOptions{}
	screen.DrawImage(firstPersonImg, fpOp)
	switch g.gameState {
	case exploration:
		if ebiten.IsKeyPressed(ebiten.KeyZ) {
			ebitenutil.DebugPrint(screen, fmt.Sprintf("%s", g.player))
		}
		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			miniMapImg = renderMiniMapView(g.player, g.gridMap, miniMapImg)
			mmOp := &ebiten.DrawImageOptions{}
			screen.DrawImage(miniMapImg, mmOp)
		}
	case combat:
		enOp := &ebiten.DrawImageOptions{}
		screen.DrawImage(entityNearImg, enOp)
		ebitenutil.DebugPrint(screen, "COMBAT\nSTUB")
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
