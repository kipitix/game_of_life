package iterator

import (
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

type position struct {
	x, y int
}

// Makes new iteration of life cycle.
// Rules of game:
// 1. Any live cell with fewer than two live neighbors dies, as if by underpopulation.
// 2. Any live cell with two or three live neighbors lives on to the next generation.
// 3. Any live cell with more than three live neighbors dies, as if by overpopulation.
// 4. sAny dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
func (i *iteratorImpl) Process(grid grid.Grid) {
	for x, column := range grid.Cells() {
		for y, curCell := range column {
			nearPos := i.findNeighborsPositions(position{x: x, y: y}, grid.Width()-1, grid.Height()-1)
			lifeNeighbors := countLifeNeighbors(grid.Cells(), nearPos)
			if curCell.State() == cell.Life {
				if lifeNeighbors < 2 || lifeNeighbors > 3 {
					curCell.SetState(cell.Dead)
				}
			} else if curCell.State() == cell.Dead {
				if lifeNeighbors == 3 {
					curCell.SetState(cell.Life)
				}
			}
		}
	}
}

// Find positions for neighbor cells.
func (i *iteratorImpl) findNeighborsPositions(pos position, maxX, maxY int) []position {
	res := make([]position, 8)

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			nextPos := position{
				x: pos.x + dx,
				y: pos.y + dy,
			}

			// TODO make case with cropField == false
			if (nextPos == pos) || (i.cropField && (nextPos.x < 0 || nextPos.y < 0 || nextPos.x > maxX || nextPos.y > maxY)) {
				continue
			}

			res = append(res, nextPos)
		}
	}

	return res
}

func countLifeNeighbors(cellsGrid grid.CellsGrid, nearPositions []position) int {
	res := 0
	for _, pos := range nearPositions {
		if cellsGrid[pos.x][pos.y].State() == cell.Life {
			res++
		}
	}
	return res
}
