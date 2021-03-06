package lifegame

// Coord X Y
type Coord struct {
	X, Y int
}

// Pattern is a set of Coord.
type Pattern map[Coord]struct{}

// NewPattern creates a pattern.
func NewPattern(coords ...Coord) Pattern {
	p := Pattern{}
	for _, coord := range coords {
		p[coord] = struct{}{}
	}
	return p
}

// Board is a two dimensional boolean grid.
type Board map[Coord]bool

// New creates a new zeroed instance of board.
func New() Board {
	return make(Board)
}

// Set the Board cell at the given coordinates.
func (b Board) SetAt(c Coord, toSet bool) {
	b[c] = toSet
	if toSet {
		// If set to true, then add the adjacent coordinates to the map.
		x := c.X
		y := c.Y
		for i := x - 1; i <= x+1; i++ {
			for j := y - 1; j <= y+1; j++ {
				adjCoord := Coord{i, j}
				b[adjCoord] = b.GetAt(adjCoord)
			}
		}
	}
}

// GetAt Get the cell state at given coordinates.
func (b Board) GetAt(c Coord) bool {
	value, ok := b[c]
	return ok && value
}

// AddPattern adds a given geometry to the board.
func (b Board) AddPattern(pattern Pattern) {
	for coord := range pattern {
		b.SetAt(coord, true)
	}
}

// GetBoundaries gets the current boundaries of the map.
func (b Board) GetBoundaries() (minX, minY, maxX, maxY int) {
	for coord := range b {
		if coord.X < minX {
			minX = coord.X
		}
		if coord.Y < minY {
			minY = coord.Y
		}
		if coord.X > maxX {
			maxX = coord.X
		}
		if coord.Y > maxY {
			maxY = coord.Y
		}
	}
	return
}

// String returns a string representation of the Board state.
func (b Board) String() string {
	// Get the current boundaries of the map.
	minX, minY, maxX, maxY := b.GetBoundaries()

	str := ""
	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			if b.GetAt(Coord{i, j}) {
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
	// Copy the old state.
	oldBoard := New()
	for k, v := range b {
		oldBoard[k] = v
	}

	for coord, cell := range b {
		adj_alife := oldBoard.CountAdjacentCells(coord.X, coord.Y, true)
		if cell {
			b.SetAt(coord, adj_alife == 2 || adj_alife == 3)
		} else {
			b.SetAt(coord, adj_alife == 3)
		}
	}
}

// CountAdjacentCells Counts the number of adjacent cells with the given state.
func (b Board) CountAdjacentCells(x, y int, state bool) int {
	count := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if x == i && y == j {
				continue
			}
			if b.GetAt(Coord{i, j}) == state {
				count++
			}
		}
	}
	return count
}
