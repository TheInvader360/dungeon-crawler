package dungeon

import (
	"bytes"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func IsEqual2DSliceInt(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func IsEqual2DSliceCell(a, b [][]cell) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func IsEqualImageArea(a, b *ebiten.Image, area image.Rectangle) bool {
	for y := area.Min.Y; y < area.Max.Y; y++ {
		for x := area.Min.X; x < area.Max.X; x++ {
			//log.Println(fmt.Sprintf("%d,%d", x, y))
			if !IsEqualColor(a.At(x, y), b.At(x, y)) {
				return false
			}
		}
	}
	return true
}

func IsEqualColor(a, b color.Color) bool {
	ar, ag, ab, aa := a.RGBA()
	br, bg, bb, ba := b.RGBA()
	//log.Println(fmt.Sprintf("a(%d,%d,%d,%d) b(%d,%d,%d,%d)", ar, ag, ab, aa, br, bg, bb, ba))
	return ar == br && ag == bg && ab == bb && aa == ba
}

func EssentialNewImageFromEncoded(encoded []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(encoded))
	if err != nil {
		panic(err)
	}
	sprite, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	return sprite
}

func essentialNewImageFromFile(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	return img
}
