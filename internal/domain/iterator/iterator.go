package iterator

import (
	"fmt"

	"github.com/kipitix/game_of_life/internal/domain/cell"
	"github.com/kipitix/game_of_life/internal/domain/grid"
)

type iteratorImpl struct {
	cropField bool
}

func NewIterator() Iterator {
	res := &iteratorImpl{
		cropField: true,
	}
	return res
}

var _ Iterator = (*iteratorImpl)(nil)

// Makes new iteration of life cycle.
// Rules of game:
// 1. Any live cell with fewer than two live neighbors dies, as if by underpopulation.
// 2. Any live cell with two or three live neighbors lives on to the next generation.
// 3. Any live cell with more than three live neighbors dies, as if by overpopulation.
// 4. sAny dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
func (i *iteratorImpl) Process(current grid.Grid, result grid.Grid) error {
	if current.Width() != result.Width() {
		return fmt.Errorf("current and result grids must have equal width")
	} else if current.Height() != result.Height() {
		return fmt.Errorf("current and result grids must have equal height")
	}

	for x, column := range current.Cells() {
		for y, curCell := range column {

			// nearPos := i.findNeighborsPositions(position{x: x, y: y}, current.Width()-1, current.Height()-1)
			lifeNeighbors := i.countLifeNeighbors(current.Cells(), x, y, current.Width()-1, current.Height()-1)

			if curCell.State() == cell.Life {
				if lifeNeighbors < 2 || lifeNeighbors > 3 {
					result.Cells()[x][y].SetState(cell.Dead)
					continue
				}
			} else if curCell.State() == cell.Dead {
				if lifeNeighbors == 3 {
					result.Cells()[x][y].SetState(cell.Life)
					continue
				}
			}

			result.Cells()[x][y].SetState(curCell.State())
		}
	}

	return nil
}

func (i *iteratorImpl) countLifeNeighbors(cellsGrid grid.CellsGrid, x, y, maxX, maxY int) int {
	res := 0

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {

			nx := x + dx
			ny := y + dy

			// TODO make case with cropField == false

			if (dx == 0 && dy == 0) || (i.cropField && (nx < 0 || ny < 0 || nx > maxX || ny > maxY)) {
				continue
			}

			if cellsGrid[nx][ny].State() == cell.Life {
				res++
			}
		}
	}

	return res
}
