package grid_test

import (
	"testing"

	"github.com/kipitix/game_of_life/internal/domain/grid"
)

func TestNewGrid(t *testing.T) {
	g := grid.NewGrid(100, 50)

	if g.Width() != 100 {
		t.Errorf("wrong width")
	}

	if g.Height() != 50 {
		t.Errorf("wrong height")
	}

	cells := g.Cells()

	if len(cells) != 100 {
		t.Errorf("wrong width")
	}

	for _, column := range cells {
		if len(column) != 50 {
			t.Errorf("wrong height")
		}
	}
}
