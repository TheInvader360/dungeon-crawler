package dungeon

type gameState int

const (
	initialize gameState = iota
	slideShow
	exploration
	combat
)
