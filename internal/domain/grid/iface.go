package grid

import "github.com/kipitix/game_of_life/internal/domain/cell"

type CellsGrid [][]cell.Cell

type Grid interface {
	Width() int
	Height() int

	Cells() CellsGrid
}
