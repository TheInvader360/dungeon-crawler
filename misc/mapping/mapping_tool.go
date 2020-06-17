package main

import (
	"errors"
	"flag"
	"fmt"
	"image/color"
	"log"
	"os"
	"strings"

	"github.com/TheInvader360/dungeon-crawler/dungeon"
	"github.com/hajimehoshi/ebiten"
)

var mappings = map[string]string{
	"#deeed6": "00", // white:  none
	"#140c1c": "10", // black:  solid
	"#442434": "11", // purple: breakable
	"#757161": "12", // grey:   locked
	"#d04648": "20", // red:    finish
	"#6daa2c": "PS", // green:  start
	"#dad45e": "30", // yellow: key
	"#d27d2c": "31", // orange: gold
	"#d2aa99": "32", // pink:   potion
	"#597dce": "CE", // blue:   enemy
}

func main() {
	code := 0
	// Run an Ebiten process so that (*Image).At is available.
	regularTermination := errors.New("regular termination")
	f := func(screen *ebiten.Image) error {
		code = process()
		return regularTermination
	}
	if err := ebiten.Run(f, 100, 100, 1, "Test"); err != nil && err != regularTermination {
		panic(err)
	}
	os.Exit(code)
}

func process() int {
	path := flag.String("path", "", "path to grid map image")
	flag.Parse()
	if *path == "" {
		log.Fatal("Path arg is required... Try: go run mapping_tool.go -path=./pix0.png")
	}

	img := dungeon.EssentialNewImageFromFile(*path)
	stringGrid := stringGridFromPixels(img)
	prettyPrint(stringGrid)
	fmt.Println("Change PS to 21/22/23/24 for player start dir N/E/S/W")
	fmt.Println("Change CE to valid enemy ref")

	return 0
}

func stringGridFromPixels(img *ebiten.Image) [][]string {
	w, h := img.Size()
	stringGrid := make([][]string, h)
	for i := 0; i < h; i++ {
		stringGrid[i] = make([]string, w)
		for j := 0; j < w; j++ {
			stringGrid[i][j] = getValForPixel(j, i, img)
		}
	}
	return stringGrid
}

func getValForPixel(x, y int, img *ebiten.Image) string {
	hex := hexColor(img.At(x, y))
	return mappings[hex]
}

func hexColor(c color.Color) string {
	rgba := color.RGBAModel.Convert(c).(color.RGBA)
	return fmt.Sprintf("#%.2x%.2x%.2x", rgba.R, rgba.G, rgba.B)
}

func prettyPrint(stringGrid [][]string) {
	var output strings.Builder
	for y := 0; y < len(stringGrid); y++ {
		output.WriteString("\n{")
		for x := 0; x < len(stringGrid[y]); x++ {
			output.WriteString(stringGrid[y][x])
			if x == len(stringGrid[y])-1 {
				output.WriteString("},")
			} else {
				output.WriteString(",")
			}
		}
	}
	fmt.Println(output.String())
}
