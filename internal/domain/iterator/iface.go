// Iterator - part of domain.
// This package contains logic of cells handling.
package iterator

import "github.com/kipitix/game_of_life/internal/domain/grid"

// Iterator goes thought all cells of grid and calculates new state for cells.
type Iterator interface {
	// Make process for grid of cells.
	Process(current grid.Grid, result grid.Grid) error
}
