package dungeon

import "fmt"

type player struct {
	x     int
	y     int
	dir   direction
	hpMax int
	hp    int
	keys  int
	gold  int
}

func newPlayer() player {
	return player{x: 0, y: 0, dir: north, hpMax: 5, hp: 3, keys: 0, gold: 0}
}

func (p player) String() string {
	return fmt.Sprintf("%s(%d,%d)\n%d/%d\nK%d\nG%d", p.dir, p.x, p.y, p.hp, p.hpMax, p.keys, p.gold)
}

func (p *player) getCoordInFront() (int, int) {
	coordX := -1
	coordY := -1
	switch p.dir {
	case north:
		coordX = p.x
		coordY = p.y - 1
	case east:
		coordX = p.x + 1
		coordY = p.y
	case south:
		coordX = p.x
		coordY = p.y + 1
	case west:
		coordX = p.x - 1
		coordY = p.y
	}
	return coordX, coordY
}

func (p *player) moveTo(x int, y int) {
	p.x = x
	p.y = y
}

func (p *player) turnLeft() {
	switch p.dir {
	case north:
		p.dir = west
	case east:
		p.dir = north
	case south:
		p.dir = east
	case west:
		p.dir = south
	}
}

func (p *player) turnRight() {
	switch p.dir {
	case north:
		p.dir = east
	case east:
		p.dir = south
	case south:
		p.dir = west
	case west:
		p.dir = north
	}
}
