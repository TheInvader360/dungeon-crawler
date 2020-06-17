package main

import (
	"log"

	"github.com/TheInvader360/dungeon-crawler/dungeon"
	"github.com/hajimehoshi/ebiten"
)

func main() {
	ebiten.SetFullscreen(true)
	//ebiten.SetWindowSize(dungeon.ScreenWidth*10, dungeon.ScreenHeight*10)
	ebiten.SetWindowTitle("Dungeon Crawler")
	if err := ebiten.RunGame(dungeon.NewGame()); err != nil {
		log.Fatal(err)
	}
}
