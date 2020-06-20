package dungeon

type gameState int

const (
	initialize gameState = iota
	exploration
	combat
	gameOver
)
