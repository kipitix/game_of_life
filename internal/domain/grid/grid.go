package grid

import "github.com/kipitix/game_of_life/internal/domain/cell"

type gridImpl struct {
	width  int
	height int
	cells  CellsGrid
}

type GridCfg struct {
	Width  int `arg:"--grid-width,env:GRID_WIDTH" default:"100" help:"Width of the grid"`
	Height int `arg:"--grid-height,env:GRID_HEIGHT" default:"100" help:"Height of the grid"`
}

func NewGrid(cfg GridCfg) Grid {
	res := &gridImpl{
		width:  cfg.Width,
		height: cfg.Height,
		cells:  make(CellsGrid, cfg.Width),
	}

	for x := 0; x < cfg.Width; x++ {
		res.cells[x] = make([]cell.Cell, cfg.Height)

		for y := 0; y < cfg.Height; y++ {
			res.cells[x][y] = cell.NewCell()
		}
	}

	return res
}

var _ Grid = (*gridImpl)(nil)

func (g *gridImpl) Width() int {
	return g.width
}

func (g *gridImpl) Height() int {
	return g.height
}

func (g *gridImpl) Cells() CellsGrid {
	return g.cells
}
