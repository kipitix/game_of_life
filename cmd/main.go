package main

import (
	"context"

	"github.com/kipitix/game_of_life/internal/application/app"
	"github.com/kipitix/game_of_life/internal/domain/grid"
	"github.com/kipitix/game_of_life/internal/domain/iterator"
	"github.com/kipitix/game_of_life/internal/domain/random"
	"github.com/kipitix/game_of_life/internal/interfaces/ui/mainwindow"
)

const (
	_width  = 200
	_height = 100
)

func main() {

	srcGrid := grid.NewGrid(_width, _height)

	rnd := random.NewRandom()
	rnd.Fill(srcGrid)

	app := app.NewApp(
		srcGrid,
		grid.NewGrid(_width, _height),
		iterator.NewIterator(),
		mainwindow.NewMainWindow(_width, _height),
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app.Run(ctx)
}
