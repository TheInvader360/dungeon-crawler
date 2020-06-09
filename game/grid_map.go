package main

func newGridMap() [][]int {
	gm := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1},
		{1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0, 1},
		{1, 1, 0, 1, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0, 1},
		{1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 1},
		{1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}
	return gm
}

func getCell(x int, y int, cells [][]int) int {
	//return -1 if invalid request
	if x < 0 || y < 0 || x >= len(cells[0]) || y >= len(cells) {
		return -1
	}
	return cells[y][x]
}

func getCells(x int, y int, w int, h int, cells [][]int) [][]int {
	//return nil if invalid request
	if w < 1 || h < 1 {
		return nil
	}

	//return subset, padded with -1 if out of bounds
	c := make([][]int, h)
	for i := 0; i < h; i++ {
		c[i] = make([]int, w)
		for j := 0; j < w; j++ {
			c[i][j] = -1
			if i+y >= 0 && j+x >= 0 && i+y < len(cells) && j+x < len(cells[0]) {
				c[i][j] = cells[i+y][j+x]
			}
		}
	}

	return c
}
