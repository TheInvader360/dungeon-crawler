package dungeon

var dungeonSrcA [][]int
var dungeonSrcB [][]int

func init() {
	dungeonSrcA = [][]int{
		{10, 10, 10, 10, 10, 10},
		{10, 50, 00, 00, 51, 10},
		{10, 00, 10, 10, 00, 10},
		{10, 00, 32, 10, 00, 10},
		{10, 21, 31, 10, 12, 10},
		{10, 00, 10, 10, 00, 10},
		{10, 11, 30, 10, 20, 10},
		{10, 10, 10, 10, 10, 10},
	}
	dungeonSrcB = [][]int{
		{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		{10, 32, 00, 00, 00, 00, 10, 00, 00, 12, 00, 00, 00, 20, 10},
		{10, 00, 10, 00, 30, 00, 10, 31, 10, 10, 10, 10, 10, 10, 10},
		{10, 55, 10, 31, 00, 56, 10, 00, 00, 00, 00, 54, 00, 32, 10},
		{10, 00, 10, 10, 10, 10, 10, 51, 10, 10, 00, 10, 00, 00, 10},
		{10, 00, 12, 00, 00, 53, 00, 00, 00, 00, 31, 10, 00, 31, 10},
		{10, 10, 10, 10, 10, 10, 10, 00, 10, 10, 10, 10, 10, 10, 10},
		{10, 10, 32, 00, 10, 10, 10, 00, 10, 00, 00, 12, 00, 30, 10},
		{10, 10, 30, 00, 00, 50, 00, 00, 00, 52, 10, 10, 00, 10, 10},
		{10, 10, 31, 00, 10, 10, 10, 00, 10, 10, 10, 10, 00, 10, 10},
		{10, 10, 10, 10, 10, 10, 10, 00, 10, 10, 31, 10, 00, 31, 10},
		{10, 31, 31, 32, 10, 10, 00, 00, 00, 10, 31, 10, 00, 00, 10},
		{10, 30, 00, 00, 11, 11, 00, 00, 00, 10, 31, 10, 00, 00, 10},
		{10, 31, 31, 32, 10, 10, 00, 21, 00, 10, 00, 11, 54, 31, 10},
		{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
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
			if src[y][x] == 10 {
				c.wall = solid
			}
			if src[y][x] == 11 {
				c.wall = breakable
			}
			if src[y][x] == 12 {
				c.wall = locked
			}
			if src[y][x] == 20 {
				//TODO portalExit
			}
			if src[y][x] == 21 {
				//TODO portalStart (NESW?)
			}
			if src[y][x] == 30 {
				c.collectible = &key
			}
			if src[y][x] == 31 {
				c.collectible = &gold
			}
			if src[y][x] == 32 {
				c.collectible = &potion
			}
			if src[y][x] == 50 {
				c.enemy = &enemies[0]
			}
			if src[y][x] == 51 {
				c.enemy = &enemies[1]
			}
			if src[y][x] == 52 {
				c.enemy = &enemies[2]
			}
			if src[y][x] == 53 {
				c.enemy = &enemies[3]
			}
			if src[y][x] == 54 {
				c.enemy = &enemies[4]
			}
			if src[y][x] == 55 {
				c.enemy = &enemies[5]
			}
			if src[y][x] == 56 {
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
