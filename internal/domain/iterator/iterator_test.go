package iterator_test

import (
	"testing"

	"github.com/kipitix/game_of_life/internal/domain/cell"
	"github.com/kipitix/game_of_life/internal/domain/grid"
	"github.com/kipitix/game_of_life/internal/domain/iterator"
)

func TestNew(t *testing.T) {
	it := iterator.NewIterator()
	if it == nil {
		t.Errorf("iterator can`t be nil")
	}
}

func TestIterateEmpty(t *testing.T) {
	it := iterator.NewIterator()

	src := grid.NewGrid(3, 3)
	dst := grid.NewGrid(3, 3)

	err := it.Process(src, dst)
	if err != nil {
		t.Errorf("fail with process: %s", err)
	}
}

func TestIterateWrongWidth(t *testing.T) {
	it := iterator.NewIterator()

	src := grid.NewGrid(3, 3)
	dst := grid.NewGrid(2, 3)

	err := it.Process(src, dst)
	if err == nil {
		t.Errorf("error can`t be nil")
	}
}

func TestIterateWrongHeight(t *testing.T) {
	it := iterator.NewIterator()

	src := grid.NewGrid(3, 3)
	dst := grid.NewGrid(3, 2)

	err := it.Process(src, dst)
	if err == nil {
		t.Errorf("error can`t be nil")
	}
}

func TestIterateOne(t *testing.T) {
	it := iterator.NewIterator()

	src := grid.NewGrid(3, 3)
	dst := grid.NewGrid(3, 3)

	src.Cells()[1][1].SetState(cell.Life)

	it.Process(src, dst)

	if src.Cells()[1][1].State() != cell.Life {
		t.Errorf("cell must be life")
	}
	if dst.Cells()[1][1].State() != cell.Dead {
		t.Errorf("cell must be dead")
	}
}

func TestIterateTwo(t *testing.T) {
	it := iterator.NewIterator()

	src := grid.NewGrid(3, 3)
	dst := grid.NewGrid(3, 3)

	src.Cells()[0][1].SetState(cell.Life)
	src.Cells()[1][1].SetState(cell.Life)

	it.Process(src, dst)

	if src.Cells()[0][1].State() != cell.Life {
		t.Errorf("cell must be life")
	}
	if src.Cells()[1][1].State() != cell.Life {
		t.Errorf("cell must be life")
	}

	if dst.Cells()[0][1].State() != cell.Dead {
		t.Errorf("cell must be dead")
	}
	if dst.Cells()[1][1].State() != cell.Dead {
		t.Errorf("cell must be dead")
	}
}

// 000    010
// 111 -> 010
// 000    010
func TestIterateThree(t *testing.T) {
	it := iterator.NewIterator()

	src := grid.NewGrid(3, 3)
	dst := grid.NewGrid(3, 3)

	src.Cells()[0][1].SetState(cell.Life)
	src.Cells()[1][1].SetState(cell.Life)
	src.Cells()[2][1].SetState(cell.Life)

	it.Process(src, dst)

	if src.Cells()[0][0].State() != cell.Dead {
		t.Errorf("cell must be dead")
	}
	if src.Cells()[1][0].State() != cell.Dead {
		t.Errorf("cell must be dead")
	}
	if src.Cells()[2][0].State() != cell.Dead {
		t.Errorf("cell must be dead")
	}
	if src.Cells()[0][1].State() != cell.Life {
		t.Errorf("cell must be life")
	}
	if src.Cells()[1][1].State() != cell.Life {
		t.Errorf("cell must be life")
	}
	if src.Cells()[2][1].State() != cell.Life {
		t.Errorf("cell must be life")
	}
	if src.Cells()[0][2].State() != cell.Dead {
		t.Errorf("cell must be dead")
	}
	if src.Cells()[1][2].State() != cell.Dead {
		t.Errorf("cell must be dead")
	}
	if src.Cells()[2][2].State() != cell.Dead {
		t.Errorf("cell must be dead")
	}

	if dst.Cells()[0][0].State() != cell.Dead {
		t.Errorf("cell must be dead")
	}
	if dst.Cells()[1][0].State() != cell.Life {
		t.Errorf("cell must be life")
	}
	if dst.Cells()[2][0].State() != cell.Dead {
		t.Errorf("cell must be dead")
	}
	if dst.Cells()[0][1].State() != cell.Dead {
		t.Errorf("cell must be dead")
	}
	if dst.Cells()[1][1].State() != cell.Life {
		t.Errorf("cell must be life")
	}
	if dst.Cells()[2][1].State() != cell.Dead {
		t.Errorf("cell must be dead")
	}
	if dst.Cells()[0][2].State() != cell.Dead {
		t.Errorf("cell must be dead")
	}
	if dst.Cells()[1][2].State() != cell.Life {
		t.Errorf("cell must be life")
	}
	if dst.Cells()[2][2].State() != cell.Dead {
		t.Errorf("cell must be dead")
	}
}

// 100    100
// 111 -> 110
// 000    010
func TestIterateFour(t *testing.T) {
	it := iterator.NewIterator()

	src := grid.NewGrid(3, 3)
	dst := grid.NewGrid(3, 3)

	src.Cells()[0][0].SetState(cell.Life)
	src.Cells()[0][1].SetState(cell.Life)
	src.Cells()[1][1].SetState(cell.Life)
	src.Cells()[2][1].SetState(cell.Life)

	it.Process(src, dst)

	if src.Cells()[0][0].State() != cell.Life {
		t.Errorf("cell must be life")
	}
	if src.Cells()[1][0].State() != cell.Dead {
		t.Errorf("cell must be dead")
	}
	if src.Cells()[2][0].State() != cell.Dead {
		t.Errorf("cell must be dead")
	}
	if src.Cells()[0][1].State() != cell.Life {
		t.Errorf("cell must be life")
	}
	if src.Cells()[1][1].State() != cell.Life {
		t.Errorf("cell must be life")
	}
	if src.Cells()[2][1].State() != cell.Life {
		t.Errorf("cell must be life")
	}
	if src.Cells()[0][2].State() != cell.Dead {
		t.Errorf("cell must be dead")
	}
	if src.Cells()[1][2].State() != cell.Dead {
		t.Errorf("cell must be dead")
	}
	if src.Cells()[2][2].State() != cell.Dead {
		t.Errorf("cell must be dead")
	}

	if dst.Cells()[0][0].State() != cell.Life {
		t.Errorf("cell must be life")
	}
	if dst.Cells()[1][0].State() != cell.Dead {
		t.Errorf("cell must be dead")
	}
	if dst.Cells()[2][0].State() != cell.Dead {
		t.Errorf("cell must be dead")
	}
	if dst.Cells()[0][1].State() != cell.Life {
		t.Errorf("cell must be life")
	}
	if dst.Cells()[1][1].State() != cell.Life {
		t.Errorf("cell must be life")
	}
	if dst.Cells()[2][1].State() != cell.Dead {
		t.Errorf("cell must be dead")
	}
	if dst.Cells()[0][2].State() != cell.Dead {
		t.Errorf("cell must be dead")
	}
	if dst.Cells()[1][2].State() != cell.Life {
		t.Errorf("cell must be life")
	}
	if dst.Cells()[2][2].State() != cell.Dead {
		t.Errorf("cell must be dead")
	}
}
