package main

import (
	"testing"
)

func TestGetCell(t *testing.T) {
	gm := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}

	if getCell(0, 0, gm) != 0 {
		t.Errorf("Expected 0 (found %v).", len(gm))
	}

	if getCell(1, 0, gm) != 1 {
		t.Errorf("Expected 1 (found %v).", len(gm))
	}

	if getCell(2, 1, gm) != 5 {
		t.Errorf("Expected 5 (found %v).", len(gm))
	}

	if getCell(0, 2, gm) != 6 {
		t.Errorf("Expected 6 (found %v).", len(gm))
	}

	if getCell(2, 2, gm) != 8 {
		t.Errorf("Expected 8 (found %v).", len(gm))
	}
}

func TestGetCells(t *testing.T) {
	gm := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}

	//complete set
	expected := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	found := getCells(0, 0, 3, 3, gm)
	if !isEqual(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//subset multirow multicol
	expected = [][]int{
		{0, 1},
		{3, 4},
	}
	found = getCells(0, 0, 2, 2, gm)
	if !isEqual(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//subset multirow multicol
	expected = [][]int{
		{4, 5},
		{7, 8},
	}
	found = getCells(1, 1, 2, 2, gm)
	if !isEqual(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//subset multirow singlecol
	expected = [][]int{
		{3},
		{6},
	}
	found = getCells(0, 1, 1, 2, gm)
	if !isEqual(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//subset singlerow multicol
	expected = [][]int{
		{7, 8},
	}
	found = getCells(1, 2, 2, 1, gm)
	if !isEqual(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//subset singlerow singlecol
	expected = [][]int{
		{5},
	}
	found = getCells(2, 1, 1, 1, gm)
	if !isEqual(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//invalid x < 0
	if getCells(-1, 1, 1, 1, gm) != nil {
		t.Errorf("Expected nil (x < 0)")
	}

	//invalid y < 0
	if getCells(1, -1, 1, 1, gm) != nil {
		t.Errorf("Expected nil (y < 0)")
	}

	//invalid w < 1
	if getCells(1, 1, 0, 1, gm) != nil {
		t.Errorf("Expected nil (w < 1)")
	}

	//invalid h < 1
	if getCells(1, 1, 1, 0, gm) != nil {
		t.Errorf("Expected nil (h < 1)")
	}

	//x+w out of bounds, scope reduced to return valid result
	expected = [][]int{
		{4, 5},
		{7, 8},
	}
	found = getCells(1, 1, 10, 2, gm)
	if !isEqual(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//y+h out of bounds, scope reduced to return valid result
	expected = [][]int{
		{4, 5},
		{7, 8},
	}
	found = getCells(1, 1, 2, 10, gm)
	if !isEqual(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//x+y and y+h out of bounds, scope reduced to return valid result
	expected = [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	found = getCells(0, 0, 10, 10, gm)
	if !isEqual(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}
}
