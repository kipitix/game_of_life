package mainwindow

import (
	"context"

	"github.com/kipitix/game_of_life/internal/domain/grid"
)

type MainWindow interface {
	Run(context.Context) error
	Render(old grid.Grid, new grid.Grid) error
}
