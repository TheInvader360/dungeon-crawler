package main

type cell struct {
	wall  bool
	enemy *enemy
}

func newCell() cell {
	return cell{wall: false, enemy: nil}
}

func (c cell) removeEnemy() cell {
	c.enemy = nil
	return c
}
