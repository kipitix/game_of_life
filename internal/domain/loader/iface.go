package loader

import "github.com/kipitix/game_of_life/internal/domain/grid"

type Loader interface {
	Load(string) (grid.Grid, error)
}
