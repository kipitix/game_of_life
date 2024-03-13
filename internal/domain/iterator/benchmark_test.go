package iterator_test

import (
	"testing"

	"github.com/kipitix/game_of_life/internal/domain/grid"
	"github.com/kipitix/game_of_life/internal/domain/iterator"
	"github.com/kipitix/game_of_life/internal/domain/random"
)

func BenchmarkIterator(b *testing.B) {
	const size = 500

	gr1 := grid.NewGrid(grid.GridCfg{Width: size, Height: size})
	gr2 := grid.NewGrid(grid.GridCfg{Width: size, Height: size})

	rnd := random.NewRandom()
	if err := rnd.Fill(gr1); err != nil {
		b.Error(err)
	}

	it := iterator.NewIterator()

	for i := 0; i < 500; i++ {
		if err := it.Process(gr1, gr2); err != nil {
			b.Error(err)
		}
		if err := it.Process(gr2, gr1); err != nil {
			b.Error(err)
		}
	}
}
