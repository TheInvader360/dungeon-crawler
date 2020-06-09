package main

func newGridMap() [][]int {
	gm := [][]int{
		{1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1, 0, 1},
		{1, 0, 1, 0, 0, 0, 1},
		{1, 0, 1, 0, 1, 0, 1},
		{1, 0, 0, 0, 1, 0, 1},
		{1, 1, 1, 0, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1},
	}
	return gm
}

func getCell(x int, y int, cells [][]int) int {
	return cells[y][x]
}

func getCells(x int, y int, w int, h int, cells [][]int) [][]int {
	//return nil if nonsense request
	if x < 0 || y < 0 || w < 1 || h < 1 {
		return nil
	}

	//shrink "w" or "h" if result would be out of bounds
	if x+w > len(cells[0]) {
		w = len(cells[0]) - x
	}
	if y+h > len(cells) {
		h = len(cells) - y
	}

	//temp copy prevents modification of underlying "cells"
	cellsCopy := make([][]int, len(cells))
	for i := range cells {
		cellsCopy[i] = make([]int, len(cells[i]))
		copy(cellsCopy[i], cells[i])
	}

	subset := cellsCopy[y : y+h]
	for i := range subset {
		subset[i] = subset[i][x : x+w]
	}
	return subset
}
