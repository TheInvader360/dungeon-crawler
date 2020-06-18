package dungeon

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type virtualButton int

const (
	u virtualButton = iota
	d
	l
	r
)

func IsPressed(btn virtualButton) bool {
	b := false

	switch btn {
	case u:
		b = ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsGamepadButtonPressed(0, ebiten.GamepadButton11)
	case d:
		b = ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsGamepadButtonPressed(0, ebiten.GamepadButton13)
	case l:
		b = ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsGamepadButtonPressed(0, ebiten.GamepadButton14)
	case r:
		b = ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsGamepadButtonPressed(0, ebiten.GamepadButton12)
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if isInBounds(btn, x, y) {
			b = true
		}
	}

	return b
}

func IsJustPressed(btn virtualButton) bool {
	b := false

	switch btn {
	case u:
		b = inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsGamepadButtonJustPressed(0, ebiten.GamepadButton11)
	case d:
		b = inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) || inpututil.IsGamepadButtonJustPressed(0, ebiten.GamepadButton13)
	case l:
		b = inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.IsGamepadButtonJustPressed(0, ebiten.GamepadButton14)
	case r:
		b = inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.IsGamepadButtonJustPressed(0, ebiten.GamepadButton12)
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if isInBounds(btn, x, y) {
			b = true
		}
	}

	return b
}

func isInBounds(btn virtualButton, x, y int) bool {
	switch btn {
	case u:
		return x >= 20 && x < 40 && y < 20
	case d:
		return x >= 20 && x < 40 && y >= 40
	case l:
		return y >= 20 && y < 40 && x < 20
	case r:
		return y >= 20 && y < 40 && x >= 40
	default:
		return false
	}
}
