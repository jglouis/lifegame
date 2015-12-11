package lifegame

// Coord X Y
type Coord struct {
	X, Y int
}

// Pattern is a slice of Coord.
type Pattern []Coord

// Board is a two dimensional boolean grid.
type Board [][]bool

// New creates a new zeroed instance of board.
func New(size int) Board {
	b := make(Board, size)
	for i := range b {
		b[i] = make([]bool, size)
	}
	return b
}

// Set the Board cell at the given coordinates.
func (b Board) SetAt(c Coord, toSet bool) {
	b[c.X][c.Y] = toSet
}

// Get the cell state at given coordinates.
func (b Board) GetAt(c Coord) bool {
	return b[c.X][c.Y]
}

// AddGeometry adds a given geometry to the board.
// Returns false if the geometry doesn't fit on the board.
func (b Board) AddGeometry(pattern Pattern) bool {
	// first check if the geometry can be added to the board
	for _, coord := range pattern {
		if coord.X > len(b)-1 || coord.Y > len(b)-1 {
			return false
		}
	}

	for _, coord := range pattern {
		b.SetAt(coord, true)
	}
	return true
}

// Get a string representation of the Board state.
func (b Board) String() string {
	str := ""
	for _, row := range b {
		for _, cell := range row {
			if cell {
				str += "1"
			} else {
				str += "0"
			}
		}
		str += "\n"
	}
	return str
}

// Tick computes the next state of the game of life.
func (b Board) Tick() {
	// for each cell:
	newBoard := New(len(b))
	for i, row := range b {
		for j, cell := range row {
			if cell {
				adj_alife := b.CountAdjacentCells(i, j, true)
				if adj_alife == 2 || adj_alife == 3 {
					newBoard[i][j] = true
				}
			} else {
				adj_alife := b.CountAdjacentCells(i, j, true)
				if adj_alife == 3 {
					newBoard[i][j] = true
				}
			}
		}
	}
	for i, row := range newBoard {
		for j, cell := range row {
			b[i][j] = cell
		}
	}

}

// Counts the number of adjacent cells with the given state.
func (b Board) CountAdjacentCells(x, y int, state bool) int {
	size := len(b)
	count := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if x == i && y == j {
				continue
			}
			if i < 0 || j < 0 || i > size-1 || j > size-1 {
				continue
			}
			if b[i][j] == state {
				count++
			}

		}
	}
	return count
}
