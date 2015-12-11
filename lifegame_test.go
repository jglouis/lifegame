package lifegame

import "testing"

func TestOscillators(t *testing.T) {
	// Oscillators patterns
	blinker := NewPattern(Coord{0, 0}, Coord{0, 1}, Coord{0, 2})
	toad := NewPattern(Coord{1, 1}, Coord{1, 2}, Coord{1, 3}, Coord{2, 0}, Coord{2, 1}, Coord{2, 2})
	beacon := NewPattern(Coord{0, 0}, Coord{1, 0}, Coord{0, 1}, Coord{2, 3}, Coord{3, 2}, Coord{3, 3})
	patterns := map[string]Pattern{
		"blinker": blinker,
		"toad":    toad,
		"beacon":  beacon}
	periodByPattern := map[string]int{
		"blinker": 2,
		"toad":    2,
		"beacon":  2}

	for name, pattern := range patterns {
		board := New()
		if !board.AddGeometry(pattern) {
			t.Error("Board is to small for the", name, "pattern.")
			continue
		}

		for i := 0; i < periodByPattern[name]; i++ {
			board.Tick()
		}

		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board); j++ {
				coord := Coord{i, j}
				_, ok := pattern[coord]
				if ok != board.GetAt(coord) {
					t.Error("Expected board[", coord.X, "][", coord.Y, "] to be", ok, "for pattern", name, ", got", !ok, ".")
				}
			}
		}
	}
}

func TestStillLifes(t *testing.T) {
	// Still patterns
	block := NewPattern(Coord{1, 1}, Coord{0, 0}, Coord{0, 1}, Coord{1, 0})
	beehive := NewPattern(Coord{0, 1}, Coord{0, 2}, Coord{1, 0}, Coord{1, 3}, Coord{2, 1}, Coord{2, 2})
	loaf := NewPattern(Coord{0, 1}, Coord{0, 2}, Coord{1, 0}, Coord{1, 3}, Coord{2, 1}, Coord{2, 3}, Coord{3, 2})
	boat := NewPattern(Coord{0, 0}, Coord{0, 1}, Coord{1, 0}, Coord{1, 2}, Coord{2, 1})

	patterns := map[string]Pattern{
		"block":   block,
		"beehive": beehive,
		"loaf":    loaf,
		"boat":    boat}

	for name, pattern := range patterns {
		board := New()
		if !board.AddGeometry(pattern) {
			t.Error("Board is to small for the", name, "pattern.")
			continue
		}
		board.Tick()

		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board); j++ {
				coord := Coord{i, j}
				_, ok := pattern[coord]
				if ok != board.GetAt(coord) {
					t.Error("Expected board[", coord.X, "][", coord.Y, "] to be", ok, "for pattern", name, ", got", !ok, ".")
				}
			}
		}

	}

}
