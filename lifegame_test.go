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
	// Still geometries
	block := []Coord{Coord{1, 1}, Coord{1, 2}, Coord{2, 1}, Coord{2, 2}}
	beehive := []Coord{Coord{1, 2}, Coord{1, 3}, Coord{2, 1}, Coord{2, 4}, Coord{3, 2}, Coord{3, 3}}

	geometries := [][]Coord{block, beehive}

	for _, geometry := range geometries {
		board := New(8)
		for _, coord := range geometry {
			board.SetAt(coord, true)
		}
		board.Tick()

		for _, coord := range geometry {
			if !board.GetAt(coord) {
				t.Error("Expected board[", coord.X, "][", coord.Y, "] to be true, got false")
			}
		}

	}

}
