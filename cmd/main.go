package main

import (
	"context"

	"github.com/alexflint/go-arg"
	"github.com/kipitix/game_of_life/internal/application/app"
	"github.com/kipitix/game_of_life/internal/domain/grid"
	"github.com/kipitix/game_of_life/internal/domain/iterator"
	"github.com/kipitix/game_of_life/internal/domain/random"
	"github.com/kipitix/game_of_life/internal/interfaces/ui/mainwindow"

	log "github.com/sirupsen/logrus"
)

func main() {

	var cfg struct {
		grid.GridCfg
		app.AppCfg
	}

	arg.MustParse(&cfg)

	log.WithField("cfg", cfg).Info("start configuration")

	srcGrid := grid.NewGrid(cfg.GridCfg)

	rnd := random.NewRandom()
	rnd.Fill(srcGrid)

	app, err := app.NewApp(
		srcGrid,
		grid.NewGrid(cfg.GridCfg),
		iterator.NewIterator(),
		mainwindow.NewMainWindow(cfg.GridCfg.Width, cfg.GridCfg.Height),
		cfg.AppCfg,
	)

	if err != nil {
		log.WithError(err).Fatal("failed to create app")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app.Run(ctx)
}
