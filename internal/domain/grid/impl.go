package grid

import "github.com/kipitix/game_of_life/internal/domain/cell"

type gridImpl struct {
	width  int
	height int
	cells  CellsGrid
}

func NewGrid(width, height int) Grid {
	res := &gridImpl{
		width:  width,
		height: height,
		cells:  make(CellsGrid, width),
	}

	for x := 0; x < width; x++ {
		res.cells[x] = make([]cell.Cell, height)

		for y := 0; y < height; y++ {
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
