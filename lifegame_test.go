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
