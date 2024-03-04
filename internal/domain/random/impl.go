package random

import (
	"math/rand"

	"github.com/kipitix/game_of_life/internal/domain/cell"
	"github.com/kipitix/game_of_life/internal/domain/grid"
)

type randomImpl struct {
}

func NewRandom() Random {
	return &randomImpl{}
}

var _ Random = (*randomImpl)(nil)

func (r *randomImpl) Fill(gr grid.Grid) error {

	states := []cell.State{
		cell.Dead,
		cell.Life,
	}

	for _, col := range gr.Cells() {
		for _, cl := range col {
			cl.SetState(states[rand.Intn(len(states))])
		}
	}

	return nil
}
