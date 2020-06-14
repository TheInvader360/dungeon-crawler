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
	testSkin     = "rgb" // 1bit / db16 / rgb
)

var (
	firstPersonImg *ebiten.Image
	miniMapImg     *ebiten.Image
	combatImg      *ebiten.Image
	err            error
)

func dungeonCrawler() *Game {
	return &Game{
		gridMap: buildGridMap(demoMapSrc),
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
	gridMap   [][]cell
	player    player
}

//Update ...
func (g *Game) Update(screen *ebiten.Image) error {
	switch g.gameState {
	case exploration:
		if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			x, y := g.player.getCoordInFront()
			target := getCell(x, y, g.gridMap)
			if !target.wall {
				g.player.moveTo(x, y)
			}
			if target.enemy != nil {
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
			c := getCell(g.player.x, g.player.y, g.gridMap).removeEnemy()
			setCell(g.player.x, g.player.y, g.gridMap, c)
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
		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			miniMapImg = renderMiniMapView(g.player, g.gridMap, miniMapImg)
			mmOp := &ebiten.DrawImageOptions{}
			screen.DrawImage(miniMapImg, mmOp)
			ebitenutil.DebugPrint(screen, fmt.Sprintf("%s", g.player))
		}
	case combat:
		enOp := &ebiten.DrawImageOptions{}
		screen.DrawImage(getCell(g.player.x, g.player.y, g.gridMap).enemy.nearImg, enOp)
		ebitenutil.DebugPrint(screen, "COMBAT\nSTUB")
	}
}

//Layout ...
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*4, screenHeight*4)
	ebiten.SetWindowTitle("Dungeon Crawler")
	if err := ebiten.RunGame(dungeonCrawler()); err != nil {
		log.Fatal(err)
	}
}
