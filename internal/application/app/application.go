package app

import (
	"context"
	"fmt"
	"time"

	"github.com/kipitix/game_of_life/internal/domain/grid"
	"github.com/kipitix/game_of_life/internal/domain/iterator"
	"github.com/kipitix/game_of_life/internal/interfaces/ui/mainwindow"
)

type appImpl struct {
	evenGrid          grid.Grid
	oddGrid           grid.Grid
	iterator          iterator.Iterator
	mainWindow        mainwindow.MainWindow
	evenCycle         bool
	lifeCycleDuration time.Duration
}

type AppCfg struct {
	LifeCycleDuration string `arg:"--life-cycle-duration,env:LIFE_CYCLE_DURATION" default:"200ms" help:"Duration of life cycle"`
}

func NewApp(
	evenGrid grid.Grid,
	oddGrid grid.Grid,
	iterator iterator.Iterator,
	mainWindow mainwindow.MainWindow,
	cfg AppCfg) (App, error) {

	res := &appImpl{
		evenGrid:   evenGrid,
		oddGrid:    oddGrid,
		iterator:   iterator,
		mainWindow: mainWindow,
		evenCycle:  true,
	}

	if lifeCycleDuration, err := time.ParseDuration(cfg.LifeCycleDuration); err != nil {
		return nil, fmt.Errorf("failed to parse life cycle duration: %w", err)
	} else {
		res.lifeCycleDuration = lifeCycleDuration
	}

	return res, nil
}

var _ App = (*appImpl)(nil)

func (a *appImpl) Run(ctx context.Context) error {

	a.mainWindow.Render(a.oddGrid, a.evenGrid)

	workFunc := func() error {
		ticker := time.NewTicker(a.lifeCycleDuration)

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
