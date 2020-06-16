package main

import (
	"fmt"
	"log"
	"os"

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
		gridMap: buildGridMap(dungeonSrcA),
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
			if target.wall == none {
				g.player.moveTo(x, y)
			} else if target.wall == breakable {
				c := getCell(x, y, g.gridMap).removeWall()
				setCell(x, y, g.gridMap, c)
				g.player.moveTo(x, y)
			} else if target.wall == locked {
				if g.player.keys > 0 {
					c := getCell(x, y, g.gridMap).removeWall()
					setCell(x, y, g.gridMap, c)
					g.player.keys--
					g.player.moveTo(x, y)
				}
			}
			if target.collectible != nil {
				if target.collectible == &key {
					g.player.keys++
				} else if target.collectible == &gold {
					g.player.gold++
				} else if target.collectible == &potion {
					g.player.hp = g.player.hpMax
				}
				c := getCell(x, y, g.gridMap).removeCollectible()
				setCell(x, y, g.gridMap, c)
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
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
	return nil
}

//Draw ...
func (g *Game) Draw(screen *ebiten.Image) {
	switch g.gameState {
	case exploration:
		if !ebiten.IsKeyPressed(ebiten.KeyDown) {
			firstPersonImg = renderFirstPersonView(g.player, g.gridMap, firstPersonImg)
			fpOp := &ebiten.DrawImageOptions{}
			screen.DrawImage(firstPersonImg, fpOp)
			ebitenutil.DebugPrint(screen, fmt.Sprintf("%s", g.player.dir))
		} else {
			miniMapImg = renderMiniMapView(g.player, g.gridMap, miniMapImg)
			mmOp := &ebiten.DrawImageOptions{}
			screen.DrawImage(miniMapImg, mmOp)
			ebitenutil.DebugPrint(screen, fmt.Sprintf("%s", g.player))
		}
	case combat:
		enOp := &ebiten.DrawImageOptions{}
		screen.DrawImage(getCell(g.player.x, g.player.y, g.gridMap).enemy.nearImg, enOp)
		ebitenutil.DebugPrint(screen, "COMBAT")
	}
}

//Layout ...
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	//ebiten.SetFullscreen(true)
	ebiten.SetWindowSize(screenWidth*10, screenHeight*10)
	ebiten.SetWindowTitle("Dungeon Crawler")
	if err := ebiten.RunGame(dungeonCrawler()); err != nil {
		log.Fatal(err)
	}
}
