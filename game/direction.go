package main

type direction int

const (
	north direction = iota
	east
	south
	west
)

func (d direction) String() string {
	return [...]string{"N", "E", "S", "W"}[d]
}
