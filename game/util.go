package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

func isEqual2DSliceInt(a, b [][]int) bool {
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

func isEqualImageArea(a, b *ebiten.Image, area image.Rectangle) bool {
	for y := area.Min.Y; y < area.Max.Y; y++ {
		for x := area.Min.X; x < area.Max.X; x++ {
			//log.Println(fmt.Sprintf("%d,%d", x, y))
			if !isEqualColor(a.At(x, y), b.At(x, y)) {
				return false
			}
		}
	}
	return true
}

func isEqualColor(a, b color.Color) bool {
	ar, ag, ab, aa := a.RGBA()
	br, bg, bb, ba := b.RGBA()
	//log.Println(fmt.Sprintf("a(%d,%d,%d,%d) b(%d,%d,%d,%d)", ar, ag, ab, aa, br, bg, bb, ba))
	return ar == br && ag == bg && ab == bb && aa == ba
}
