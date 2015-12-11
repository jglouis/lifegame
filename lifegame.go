package lifegame

type Board [][]bool

func New(size uint) Board {
	b := make(Board, size)
	for i := range b {
		b[i] = make([]bool, size)
	}
	return b
}

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

// One tick of game of life
func (b Board) Tick() {
	// for each cell:
	newBoard := New(3)
	for i, row := range b {
		for j, cell := range row {
			if cell {
				adj_alife := b.CountAdjacentCells(i, j, true)
				if adj_alife == 2 || adj_alife == 3 {
					newBoard[i][j] = true
				} else {
					{
						newBoard[i][j] = false
					}
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

func (b Board) CountAdjacentCells(x, y int, isActive bool) int {
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
			if b[i][j] == isActive {
				count++
			}

		}
	}
	return count
}
