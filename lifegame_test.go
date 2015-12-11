package lifegame

import "testing"

func TestBlinker(t *testing.T) {
	board := New(3)
	board[1][0] = true
	board[1][1] = true
	board[1][2] = true
	board.Tick()
	board.Tick()

	for i := 0; i <= 2; i++ {
		if !board[1][i] {
			t.Error("Expected board[1][", i, "] to be true, got false")
		}
	}
}

func TestStillLifes(t *testing.T) {
	// Still patterns
	block := Pattern{Coord{1, 1}, Coord{1, 2}, Coord{2, 1}, Coord{2, 2}}
	beehive := Pattern{Coord{1, 2}, Coord{1, 3}, Coord{2, 1}, Coord{2, 4}, Coord{3, 2}, Coord{3, 3}}
	loaf := Pattern{Coord{1, 2}, Coord{1, 3}, Coord{2, 1}, Coord{2, 4}, Coord{3, 2}, Coord{3, 4}, Coord{4, 3}}
	boat := Pattern{Coord{1, 1}, Coord{1, 2}, Coord{2, 1}, Coord{2, 3}, Coord{3, 2}}

	patterns := map[string]Pattern{
		"block":   block,
		"beehive": beehive,
		"loaf":    loaf,
		"boat":    boat}

	for name, geometry := range patterns {
		board := New(5)
		if !board.AddGeometry(geometry) {
			t.Fatal("Board is to small for the", name, "pattern.")
		}
		board.Tick()

		for _, coord := range geometry {
			if !board.GetAt(coord) {
				t.Error("Expected board[", coord.X, "][", coord.Y, "] to be true, got false.")
			}
		}

	}

}
