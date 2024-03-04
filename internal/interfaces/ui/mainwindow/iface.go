package mainwindow

import "context"

type MainWindow interface {
	Run(context.Context) error
}
