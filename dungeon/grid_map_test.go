package dungeon

import (
	"testing"
)

func TestGetCell(t *testing.T) {
	src := [][]int{
		{10, 50, 00},
		{00, 00, 51},
		{52, 00, 53},
	}
	gm := buildGridMap(src)
	expected := newCell()
	expected.wall = solid
	found := getCell(0, 0, gm)
	if found != expected {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	expected.wall = none
	expected.enemy = &enemies[0]
	found = getCell(1, 0, gm)
	if found != expected {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	expected.enemy = &enemies[1]
	found = getCell(2, 1, gm)
	if found != expected {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	expected.enemy = &enemies[2]
	found = getCell(0, 2, gm)
	if found != expected {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	expected.enemy = &enemies[3]
	found = getCell(2, 2, gm)
	if found != expected {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//invalid x < 0
	expected = newCell()
	found = getCell(-1, 1, gm)
	if found != expected {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//invalid y < 0
	found = getCell(1, -1, gm)
	if found != expected {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//invalid x out of bounds
	found = getCell(3, 1, gm)
	if found != expected {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//invalid y out of bounds
	found = getCell(1, 3, gm)
	if found != expected {
		t.Errorf("Expected %v (found %v).", expected, found)
	}
}

func TestGetCells(t *testing.T) {
	src := [][]int{
		{00, 10, 50},
		{51, 52, 53},
		{54, 55, 56},
	}
	gm := buildGridMap(src)

	//complete set
	expectedSrc := [][]int{
		{00, 10, 50},
		{51, 52, 53},
		{54, 55, 56},
	}
	expected := buildGridMap(expectedSrc)
	found := getCells(0, 0, 3, 3, gm)
	if !IsEqual2DSliceCell(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//subset multirow multicol
	expectedSrc = [][]int{
		{00, 10},
		{51, 52},
	}
	expected = buildGridMap(expectedSrc)
	found = getCells(0, 0, 2, 2, gm)
	if !IsEqual2DSliceCell(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//subset multirow multicol
	expectedSrc = [][]int{
		{52, 53},
		{55, 56},
	}
	expected = buildGridMap(expectedSrc)
	found = getCells(1, 1, 2, 2, gm)
	if !IsEqual2DSliceCell(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//subset multirow singlecol
	expectedSrc = [][]int{
		{51},
		{54},
	}
	expected = buildGridMap(expectedSrc)
	found = getCells(0, 1, 1, 2, gm)
	if !IsEqual2DSliceCell(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//subset singlerow multicol
	expectedSrc = [][]int{
		{55, 56},
	}
	expected = buildGridMap(expectedSrc)
	found = getCells(1, 2, 2, 1, gm)
	if !IsEqual2DSliceCell(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//subset singlerow singlecol
	expectedSrc = [][]int{
		{53},
	}
	expected = buildGridMap(expectedSrc)
	found = getCells(2, 1, 1, 1, gm)
	if !IsEqual2DSliceCell(expected, found) {
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

	//out of top bounds
	expectedSrc = [][]int{
		{-1, -1, -1},
		{00, 10, 50},
		{51, 52, 53},
	}
	expected = buildGridMap(expectedSrc)
	found = getCells(0, -1, 3, 3, gm)
	if !IsEqual2DSliceCell(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of right bounds
	expectedSrc = [][]int{
		{52, 53, -1},
		{55, 56, -1},
	}
	expected = buildGridMap(expectedSrc)
	found = getCells(1, 1, 3, 2, gm)
	if !IsEqual2DSliceCell(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of bottom bounds
	expectedSrc = [][]int{
		{53},
		{56},
		{-1},
	}
	expected = buildGridMap(expectedSrc)
	found = getCells(2, 1, 1, 3, gm)
	if !IsEqual2DSliceCell(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of left bounds
	expectedSrc = [][]int{
		{-1, -1, 51, 52},
	}
	expected = buildGridMap(expectedSrc)
	found = getCells(-2, 1, 4, 1, gm)
	if !IsEqual2DSliceCell(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of top-left bounds
	expectedSrc = [][]int{
		{-1, -1, -1, -1},
		{-1, -1, 00, 10},
	}
	expected = buildGridMap(expectedSrc)
	found = getCells(-2, -1, 4, 2, gm)
	if !IsEqual2DSliceCell(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of top-right bounds
	expectedSrc = [][]int{
		{-1, -1, -1, -1},
		{10, 50, -1, -1},
	}
	expected = buildGridMap(expectedSrc)
	found = getCells(1, -1, 4, 2, gm)
	if !IsEqual2DSliceCell(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of bottom-left bounds
	expectedSrc = [][]int{
		{-1, 54, 55},
		{-1, -1, -1},
	}
	expected = buildGridMap(expectedSrc)
	found = getCells(-1, 2, 3, 2, gm)
	if !IsEqual2DSliceCell(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of bottom-right bounds
	expectedSrc = [][]int{
		{55, 56, -1},
		{-1, -1, -1},
	}
	expected = buildGridMap(expectedSrc)
	found = getCells(1, 2, 3, 2, gm)
	if !IsEqual2DSliceCell(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of top-bottom bounds
	expectedSrc = [][]int{
		{-1, -1},
		{00, 10},
		{51, 52},
		{54, 55},
		{-1, -1},
		{-1, -1},
	}
	expected = buildGridMap(expectedSrc)
	found = getCells(0, -1, 2, 6, gm)
	if !IsEqual2DSliceCell(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of left-right bounds
	expectedSrc = [][]int{
		{-1, -1, 51, 52, 53, -1, -1},
	}
	expected = buildGridMap(expectedSrc)
	found = getCells(-2, 1, 7, 1, gm)
	if !IsEqual2DSliceCell(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of left-top-right bounds
	expectedSrc = [][]int{
		{-1, -1, -1, -1, -1},
		{-1, 00, 10, 50, -1},
	}
	expected = buildGridMap(expectedSrc)
	found = getCells(-1, -1, 5, 2, gm)
	if !IsEqual2DSliceCell(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of top-right-bottom bounds
	expectedSrc = [][]int{
		{-1, -1, -1},
		{10, 50, -1},
		{52, 53, -1},
		{55, 56, -1},
		{-1, -1, -1},
	}
	expected = buildGridMap(expectedSrc)
	found = getCells(1, -1, 3, 5, gm)
	if !IsEqual2DSliceCell(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of right-bottom-left bounds
	expectedSrc = [][]int{
		{-1, 54, 55, 56, -1},
		{-1, -1, -1, -1, -1},
	}
	expected = buildGridMap(expectedSrc)
	found = getCells(-1, 2, 5, 2, gm)
	if !IsEqual2DSliceCell(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of bottom-left-top bounds
	expectedSrc = [][]int{
		{-1, -1, -1, -1},
		{-1, -1, -1, 00},
		{-1, -1, -1, 51},
		{-1, -1, -1, 54},
		{-1, -1, -1, -1},
		{-1, -1, -1, -1},
	}
	expected = buildGridMap(expectedSrc)
	found = getCells(-3, -1, 4, 6, gm)
	if !IsEqual2DSliceCell(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}

	//out of all bounds
	expectedSrc = [][]int{
		{-1, -1, -1, -1, -1},
		{-1, 00, 10, 50, -1},
		{-1, 51, 52, 53, -1},
		{-1, 54, 55, 56, -1},
		{-1, -1, -1, -1, -1},
	}
	expected = buildGridMap(expectedSrc)
	found = getCells(-1, -1, 5, 5, gm)
	if !IsEqual2DSliceCell(expected, found) {
		t.Errorf("Expected %v (found %v).", expected, found)
	}
}
