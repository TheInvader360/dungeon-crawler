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

	expected := 0
	found := getCell(0, 0, gm)
	if found != expected {
		t.Errorf("Expected %d (found %d).", expected, found)
	}

	expected = 1
	found = getCell(1, 0, gm)
	if found != expected {
		t.Errorf("Expected %d (found %d).", expected, found)
	}

	expected = 5
	found = getCell(2, 1, gm)
	if found != expected {
		t.Errorf("Expected %d (found %d).", expected, found)
	}

	expected = 6
	found = getCell(0, 2, gm)
	if found != expected {
		t.Errorf("Expected %d (found %d).", expected, found)
	}

	expected = 8
	found = getCell(2, 2, gm)
	if found != expected {
		t.Errorf("Expected %d (found %d).", expected, found)
	}

	//invalid x < 0
	expected = -1
	found = getCell(-1, 1, gm)
	if found != expected {
		t.Errorf("Expected %d (found %d).", expected, found)
	}

	//invalid y < 0
	expected = -1
	found = getCell(1, -1, gm)
	if found != expected {
		t.Errorf("Expected %d (found %d).", expected, found)
	}

	//invalid x out of bounds
	expected = -1
	found = getCell(3, 1, gm)
	if found != expected {
		t.Errorf("Expected %d (found %d).", expected, found)
	}

	//invalid y out of bounds
	expected = -1
	found = getCell(1, 3, gm)
	if found != expected {
		t.Errorf("Expected %d (found %d).", expected, found)
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
