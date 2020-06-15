package main

import (
	"errors"
	"fmt"
	"image"
	"os"
	"strings"
	"testing"

	"github.com/hajimehoshi/ebiten"
)

func TestMain(m *testing.M) {
	code := 0
	// Run an Ebiten process so that (*Image).At is available.
	regularTermination := errors.New("regular termination")
	f := func(screen *ebiten.Image) error {
		code = m.Run()
		return regularTermination
	}
	if err := ebiten.Run(f, screenWidth, screenHeight, 1, "Test"); err != nil && err != regularTermination {
		panic(err)
	}
	os.Exit(code)
}

func TestMiniMap(t *testing.T) {
	players := []player{
		player{x: 1, y: 1, dir: north},
		player{x: 7, y: 1, dir: east},
		player{x: 4, y: 4, dir: south},
		player{x: 5, y: 5, dir: west},
		player{x: 8, y: 6, dir: north},
		player{x: 4, y: 8, dir: south},
	}
	src := [][]int{
		{10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		{10, 00, 10, 50, 00, 50, 10, 00, 00, 10},
		{10, 00, 10, 10, 00, 00, 50, 00, 00, 10},
		{10, 00, 00, 10, 00, 00, 00, 10, 10, 10},
		{10, 00, 00, 00, 00, 00, 00, 00, 50, 10},
		{10, 50, 00, 50, 00, 00, 10, 00, 00, 10},
		{10, 00, 10, 10, 00, 00, 10, 00, 00, 10},
		{10, 00, 10, 10, 00, 00, 10, 00, 00, 10},
		{10, 00, 50, 00, 00, 00, 10, 00, 50, 10},
		{10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
	}
	gm := buildGridMap(src)
	area := image.Rect(15, 15, 60, 60)
	var infos []string

	for i, p := range players {
		expected := essentialNewImageFromFile(fmt.Sprintf("../assets/testing/mini_map/%d.png", i))
		v, _ := ebiten.NewImage(screenWidth, screenHeight, ebiten.FilterNearest)
		found := renderMiniMapView(p, gm, v)
		if !isEqualImageArea(expected, found, area) {
			infos = append(infos, fmt.Sprintf("\nFailed scenario %d where player=%v", i, p))
		}
	}

	if len(infos) > 0 {
		t.Errorf(strings.Join(infos, ""))
	}
}
