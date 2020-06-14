package main

var demoMapSrc [][]int

func init() {
	demoMapSrc = [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 2, 1, 4, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 1},
		{1, 0, 1, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 2, 1},
		{1, 0, 0, 0, 1, 6, 6, 6, 0, 1, 1, 1, 1, 0, 1},
		{1, 0, 1, 0, 0, 2, 2, 2, 0, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 6, 5, 4, 3, 2, 0, 0, 3, 0, 0, 3, 1},
		{1, 0, 1, 0, 0, 5, 5, 5, 1, 1, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 1, 0, 1},
		{1, 1, 1, 1, 0, 1, 0, 0, 2, 0, 0, 1, 1, 4, 1},
		{1, 8, 0, 1, 0, 1, 1, 0, 4, 1, 1, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 1},
		{1, 7, 0, 0, 0, 0, 1, 6, 0, 0, 6, 0, 0, 5, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}
}

func buildGridMap(src [][]int) [][]cell {
	h := len(src)
	w := len(src[0])
	gm := make([][]cell, h)
	for y := 0; y < h; y++ {
		gm[y] = make([]cell, w)
		for x := 0; x < w; x++ {
			c := newCell()
			if src[y][x] == 1 {
				c.wall = true
			}
			if src[y][x] == 2 {
				c.enemy = &enemies[0]
			}
			if src[y][x] == 3 {
				c.enemy = &enemies[1]
			}
			if src[y][x] == 4 {
				c.enemy = &enemies[2]
			}
			if src[y][x] == 5 {
				c.enemy = &enemies[3]
			}
			if src[y][x] == 6 {
				c.enemy = &enemies[4]
			}
			if src[y][x] == 7 {
				c.enemy = &enemies[5]
			}
			if src[y][x] == 8 {
				c.enemy = &enemies[6]
			}
			gm[y][x] = c
		}
	}
	return gm
}

func setCell(x, y int, cells [][]cell, cell cell) {
	cells[y][x] = cell
}

func getCell(x, y int, cells [][]cell) cell {
	//return default if invalid request
	if x < 0 || y < 0 || x >= len(cells[0]) || y >= len(cells) {
		return newCell()
	}
	return cells[y][x]
}

func getCells(x int, y int, w int, h int, cells [][]cell) [][]cell {
	//return nil if invalid request
	if w < 1 || h < 1 {
		return nil
	}

	//return subset, padded with defaults if out of bounds
	c := make([][]cell, h)
	for i := 0; i < h; i++ {
		c[i] = make([]cell, w)
		for j := 0; j < w; j++ {
			c[i][j] = newCell()
			if i+y >= 0 && j+x >= 0 && i+y < len(cells) && j+x < len(cells[0]) {
				c[i][j] = cells[i+y][j+x]
			}
		}
	}

	return c
}
