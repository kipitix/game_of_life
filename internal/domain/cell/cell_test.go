package cell_test

import (
	"testing"

	"github.com/kipitix/game_of_life/internal/domain/cell"
)

func TestNewCell(t *testing.T) {
	c := cell.NewCell()

	if c.State() != cell.Dead {
		t.Errorf("default state must be dead")
	}

	c.SetState(cell.Life)

	if c.State() != cell.Life {
		t.Errorf("new state must be life")
	}
}
