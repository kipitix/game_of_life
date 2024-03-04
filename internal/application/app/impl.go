package app

import (
	"context"
	"time"

	"github.com/kipitix/game_of_life/internal/domain/grid"
	"github.com/kipitix/game_of_life/internal/domain/iterator"
	"github.com/kipitix/game_of_life/internal/interfaces/ui/mainwindow"
)

type appImpl struct {
	evenGrid   grid.Grid
	oddGrid    grid.Grid
	iterator   iterator.Iterator
	mainWindow mainwindow.MainWindow
	evenCycle  bool
}

func NewApp(
	evenGrid grid.Grid,
	oddGrid grid.Grid,
	iterator iterator.Iterator,
	mainWindow mainwindow.MainWindow) App {

	return &appImpl{
		evenGrid:   evenGrid,
		oddGrid:    oddGrid,
		iterator:   iterator,
		mainWindow: mainWindow,
		evenCycle:  true,
	}
}

var _ App = (*appImpl)(nil)

func (a *appImpl) Run(ctx context.Context) error {

	a.mainWindow.Render(a.oddGrid, a.evenGrid)

	workFunc := func() error {
		ticker := time.NewTicker(time.Millisecond * 100)

		for {
			select {

			case <-ticker.C:
				if a.evenCycle {
					if err := a.iterator.Process(a.evenGrid, a.oddGrid); err != nil {
						return err
					}
					a.mainWindow.Render(a.evenGrid, a.oddGrid)
				} else {
					if err := a.iterator.Process(a.oddGrid, a.evenGrid); err != nil {
						return err
					}
					a.mainWindow.Render(a.oddGrid, a.evenGrid)
				}
				a.evenCycle = !a.evenCycle

			case <-ctx.Done():
				return nil
			}
		}
	}

	// var err error
	go func() {
		workFunc()
	}()

	if err := a.mainWindow.Run(ctx); err != nil {
		return err
	}

	return nil
}
