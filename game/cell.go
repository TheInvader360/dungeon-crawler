package main

type wallType int

const (
	none wallType = iota
	solid
	breakable
	locked
)

type cell struct {
	wall  wallType
	enemy *enemy
}

func newCell() cell {
	return cell{wall: none, enemy: nil}
}

func (c cell) removeWall() cell {
	c.wall = none
	return c
}

func (c cell) removeEnemy() cell {
	c.enemy = nil
	return c
}
