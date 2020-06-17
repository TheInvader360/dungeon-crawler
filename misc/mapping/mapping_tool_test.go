package main

import (
	"errors"
	"os"
	"testing"

	"github.com/TheInvader360/dungeon-crawler/dungeon"
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
	if err := ebiten.Run(f, 100, 100, 1, "Test"); err != nil && err != regularTermination {
		panic(err)
	}
	os.Exit(code)
}

func TestStringGridFromPixels(t *testing.T) {
	expected := [][]string{
		{"10", "10", "10", "10", "10", "10"},
		{"10", "CE", "00", "00", "CE", "10"},
		{"10", "00", "10", "10", "00", "10"},
		{"10", "00", "32", "10", "00", "10"},
		{"10", "PS", "31", "10", "12", "10"},
		{"10", "00", "10", "10", "00", "10"},
		{"10", "11", "30", "10", "20", "10"},
		{"10", "10", "10", "10", "10", "10"},
	}
	img := dungeon.EssentialNewImageFromFile("./test.png")
	found := stringGridFromPixels(img)
	if !dungeon.IsEqual2DSliceString(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}
}
