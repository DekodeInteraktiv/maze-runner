package game

type MazeTileType uint16

const (
	Floor MazeTileType = iota
	Wall
	Portal
)

func newGrid(size int) [][]MazeTileType {
	grid := make([][]MazeTileType, size)
	for i := range grid {
		grid[i] = make([]MazeTileType, size)
	}

	return grid
}
