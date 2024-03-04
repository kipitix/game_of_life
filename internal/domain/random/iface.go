package random

import "github.com/kipitix/game_of_life/internal/domain/grid"

type Random interface {
	Fill(grid.Grid) error
}
