package lifegame

import "testing"

func TestBlinker(t *testing.T) {
	// Oscillators patterns
	blinker := Pattern{Coord{1, 0}, Coord{1, 1}, Coord{1, 2}}
	toad := Pattern{Coord{1, 1}, Coord{1, 2}, Coord{1, 3}, Coord{2, 0}, Coord{2, 1}, Coord{2, 2}}
	beacon := Pattern{Coord{0, 0}, Coord{1, 0}, Coord{0, 1}, Coord{2, 3}, Coord{3, 2}, Coord{3, 3}}
	patterns := map[string]Pattern{
		"blinker": blinker,
		"toad":    toad,
		"beacon":  beacon}
	periodByPattern := map[string]int{
		"blinker": 2,
		"toad":    2,
		"beacon":  2}

	for name, geometry := range patterns {
		board := New(4)
		if !board.AddGeometry(geometry) {
			t.Error("Board is to small for the", name, "pattern.")
			continue
		}

		for i := 0; i < periodByPattern[name]; i++ {
			board.Tick()
		}

		for _, coord := range geometry {
			if !board.GetAt(coord) {
				t.Error("Expected board[", coord.X, "][", coord.Y, "] to be true, got false.")
			}
		}
	}
}

func TestStillLifes(t *testing.T) {
	// Still patterns
	block := Pattern{Coord{1, 1}, Coord{0, 0}, Coord{0, 1}, Coord{1, 0}}
	beehive := Pattern{Coord{0, 1}, Coord{0, 2}, Coord{1, 0}, Coord{1, 3}, Coord{2, 1}, Coord{2, 2}}
	loaf := Pattern{Coord{0, 1}, Coord{0, 2}, Coord{1, 0}, Coord{1, 3}, Coord{2, 1}, Coord{2, 3}, Coord{3, 2}}
	boat := Pattern{Coord{0, 0}, Coord{0, 1}, Coord{1, 0}, Coord{1, 2}, Coord{2, 1}}

	patterns := map[string]Pattern{
		"block":   block,
		"beehive": beehive,
		"loaf":    loaf,
		"boat":    boat}

	for name, geometry := range patterns {
		board := New(4)
		if !board.AddGeometry(geometry) {
			t.Error("Board is to small for the", name, "pattern.")
			continue
		}
		board.Tick()

		for _, coord := range geometry {
			if !board.GetAt(coord) {
				t.Error("Expected board[", coord.X, "][", coord.Y, "] to be true, got false.")
			}
		}

	}

}
