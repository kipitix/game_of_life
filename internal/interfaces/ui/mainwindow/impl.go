package mainwindow

import "context"

type mainWindowImpl struct {
}

func NewMainWindow() MainWindow {
	return nil
}

var _ MainWindow = (*mainWindowImpl)(nil)

func (w *mainWindowImpl) Run(ctx context.Context) error {
	return nil
}
