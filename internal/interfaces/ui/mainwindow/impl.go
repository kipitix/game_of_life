package mainwindow

import (
	"context"
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/kipitix/game_of_life/internal/domain/cell"
	"github.com/kipitix/game_of_life/internal/domain/grid"
)

const (
	_size     = 4
	_margin   = 1
	_fullSize = _size + _margin
)

var (
	_black = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	_white = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
)

type mainWindowImpl struct {
	app           fyne.App
	window        fyne.Window
	container     *fyne.Container
	width, height int
	cells         [][]*canvas.Rectangle
	animate       bool
}

func NewMainWindow(width, height int) MainWindow {
	res := &mainWindowImpl{
		width:   width,
		height:  height,
		animate: false,
	}

	res.app = app.New()
	res.window = res.app.NewWindow("GAME OF LIFE")

	tmpCells := make([]fyne.CanvasObject, width*height)

	res.cells = make([][]*canvas.Rectangle, width)
	for x := range res.cells {
		res.cells[x] = make([]*canvas.Rectangle, height)
		for y := range res.cells[x] {
			rect := canvas.NewRectangle(_black)
			rect.Resize(fyne.NewSize(_size, _size))
			rect.Move(fyne.NewPos(float32(x*_fullSize), float32(y*_fullSize)))
			res.cells[x][y] = rect

			tmpCells[x+y*width] = rect
		}
	}

	res.container = container.NewWithoutLayout(tmpCells...)

	res.window.SetContent(res.container)

	return res
}

var _ MainWindow = (*mainWindowImpl)(nil)

func (w *mainWindowImpl) Run(ctx context.Context) error {

	w.window.Resize(fyne.NewSize(float32(w.width*_fullSize), float32(w.height*_fullSize)))
	w.window.SetPadded(false)
	w.window.ShowAndRun()

	return nil
}

func (w *mainWindowImpl) Render(old grid.Grid, new grid.Grid) error {

	if w.width != old.Width() {
		return fmt.Errorf("view and model grids must have equal width")
	} else if w.width != new.Width() {
		return fmt.Errorf("view and model grids must have equal width")
	} else if w.height != old.Height() {
		return fmt.Errorf("view and model grids must have equal height")
	} else if w.height != new.Height() {
		return fmt.Errorf("view and model grids must have equal height")
	}

	for x, col := range new.Cells() {
		for y, cl := range col {

			if old.Cells()[x][y].State() != new.Cells()[x][y].State() {
				rect := w.cells[x][y]

				oldColor := _black
				newColor := _black

				if cl.State() == cell.Life {
					newColor = _white
				} else {
					oldColor = _white
				}

				if w.animate {
					canvas.NewColorRGBAAnimation(oldColor, newColor, time.Millisecond*200, func(c color.Color) {
						rect.FillColor = c
						canvas.Refresh(rect)
					}).Start()
				} else {
					rect.FillColor = newColor
				}
			}
		}
	}

	w.container.Refresh()

	return nil
}
