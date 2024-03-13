package loader

import (
	"github.com/kipitix/game_of_life/internal/domain/grid"
)

type loaderImpl struct {
}

func NewLoader() Loader {
	return nil
}

var _ Loader = (*loaderImpl)(nil)

func (l *loaderImpl) Load(path string) (grid.Grid, error) {

	// data, err := os.ReadFile(path)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil // TODO: implement this
}
