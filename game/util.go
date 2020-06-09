package main

func isEqual(a [][]int, b [][]int) bool {
	equal := true
	for i := range a {
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				equal = false
			}
		}
	}
	return equal
}
