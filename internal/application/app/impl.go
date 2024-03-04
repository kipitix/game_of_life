package app

import (
	"context"

	"github.com/kipitix/game_of_life/internal/domain/grid"
	"github.com/kipitix/game_of_life/internal/domain/iterator"
	"github.com/kipitix/game_of_life/internal/interfaces/ui/mainwindow"
)

type appImpl struct {
	grid       grid.Grid
	iterator   iterator.Iterator
	mainWindow mainwindow.MainWindow
}

func NewApp(
	grid grid.Grid,
	iterator iterator.Iterator,
	mainWindow mainwindow.MainWindow) App {

	return &appImpl{
		grid:       grid,
		iterator:   iterator,
		mainWindow: mainWindow,
	}
}

var _ App = (*appImpl)(nil)

func (a *appImpl) Run(ctx context.Context) error {
	return nil
}
