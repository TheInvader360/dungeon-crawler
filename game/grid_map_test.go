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
	if !isEqual2DSliceInt(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//subset multirow multicol
	expected = [][]int{
		{0, 1},
		{3, 4},
	}
	found = getCells(0, 0, 2, 2, gm)
	if !isEqual2DSliceInt(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//subset multirow multicol
	expected = [][]int{
		{4, 5},
		{7, 8},
	}
	found = getCells(1, 1, 2, 2, gm)
	if !isEqual2DSliceInt(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//subset multirow singlecol
	expected = [][]int{
		{3},
		{6},
	}
	found = getCells(0, 1, 1, 2, gm)
	if !isEqual2DSliceInt(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//subset singlerow multicol
	expected = [][]int{
		{7, 8},
	}
	found = getCells(1, 2, 2, 1, gm)
	if !isEqual2DSliceInt(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//subset singlerow singlecol
	expected = [][]int{
		{5},
	}
	found = getCells(2, 1, 1, 1, gm)
	if !isEqual2DSliceInt(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//invalid w < 1
	if getCells(1, 1, 0, 1, gm) != nil {
		t.Errorf("Expected nil (w < 1)")
	}

	//invalid h < 1
	if getCells(1, 1, 1, 0, gm) != nil {
		t.Errorf("Expected nil (h < 1)")
	}

	//out of top bounds, padded with -1's to return valid result
	expected = [][]int{
		{-1, -1, -1},
		{0, 1, 2},
		{3, 4, 5},
	}
	found = getCells(0, -1, 3, 3, gm)
	if !isEqual2DSliceInt(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of right bounds, padded with -1's to return valid result
	expected = [][]int{
		{4, 5, -1},
		{7, 8, -1},
	}
	found = getCells(1, 1, 3, 2, gm)
	if !isEqual2DSliceInt(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of bottom bounds, padded with -1's to return valid result
	expected = [][]int{
		{5},
		{8},
		{-1},
	}
	found = getCells(2, 1, 1, 3, gm)
	if !isEqual2DSliceInt(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of left bounds, padded with -1's to return valid result
	expected = [][]int{
		{-1, -1, 3, 4},
	}
	found = getCells(-2, 1, 4, 1, gm)
	if !isEqual2DSliceInt(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of top-left bounds, padded with -1's to return valid result
	expected = [][]int{
		{-1, -1, -1, -1},
		{-1, -1, 0, 1},
	}
	found = getCells(-2, -1, 4, 2, gm)
	if !isEqual2DSliceInt(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of top-right bounds, padded with -1's to return valid result
	expected = [][]int{
		{-1, -1, -1, -1},
		{1, 2, -1, -1},
	}
	found = getCells(1, -1, 4, 2, gm)
	if !isEqual2DSliceInt(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of bottom-left bounds, padded with -1's to return valid result
	expected = [][]int{
		{-1, 6, 7},
		{-1, -1, -1},
	}
	found = getCells(-1, 2, 3, 2, gm)
	if !isEqual2DSliceInt(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of bottom-right bounds, padded with -1's to return valid result
	expected = [][]int{
		{7, 8, -1},
		{-1, -1, -1},
	}
	found = getCells(1, 2, 3, 2, gm)
	if !isEqual2DSliceInt(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of top-bottom bounds, padded with -1's to return valid result
	expected = [][]int{
		{-1, -1},
		{0, 1},
		{3, 4},
		{6, 7},
		{-1, -1},
		{-1, -1},
	}
	found = getCells(0, -1, 2, 6, gm)
	if !isEqual2DSliceInt(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of left-right bounds, padded with -1's to return valid result
	expected = [][]int{
		{-1, -1, 3, 4, 5, -1, -1},
	}
	found = getCells(-2, 1, 7, 1, gm)
	if !isEqual2DSliceInt(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of left-top-right bounds, padded with -1's to return valid result
	expected = [][]int{
		{-1, -1, -1, -1, -1},
		{-1, 0, 1, 2, -1},
	}
	found = getCells(-1, -1, 5, 2, gm)
	if !isEqual2DSliceInt(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of top-right-bottom bounds, padded with -1's to return valid result
	expected = [][]int{
		{-1, -1, -1},
		{1, 2, -1},
		{4, 5, -1},
		{7, 8, -1},
		{-1, -1, -1},
	}
	found = getCells(1, -1, 3, 5, gm)
	if !isEqual2DSliceInt(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of right-bottom-left bounds, padded with -1's to return valid result
	expected = [][]int{
		{-1, 6, 7, 8, -1},
		{-1, -1, -1, -1, -1},
	}
	found = getCells(-1, 2, 5, 2, gm)
	if !isEqual2DSliceInt(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of bottom-left-top bounds, padded with -1's to return valid result
	expected = [][]int{
		{-1, -1, -1, -1},
		{-1, -1, -1, 0},
		{-1, -1, -1, 3},
		{-1, -1, -1, 6},
		{-1, -1, -1, -1},
		{-1, -1, -1, -1},
	}
	found = getCells(-3, -1, 4, 6, gm)
	if !isEqual2DSliceInt(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of all bounds, padded with -1's to return valid result
	expected = [][]int{
		{-1, -1, -1, -1, -1},
		{-1, 0, 1, 2, -1},
		{-1, 3, 4, 5, -1},
		{-1, 6, 7, 8, -1},
		{-1, -1, -1, -1, -1},
	}
	found = getCells(-1, -1, 5, 5, gm)
	if !isEqual2DSliceInt(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}
}
