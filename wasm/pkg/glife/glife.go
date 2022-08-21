package glife

import "math/rand"

type WorldState [][]int

// Creates a new world state.
func NewWorldState(colSize, rowSize int) WorldState {
	b := make([][]int, rowSize)

	for i := 0; i < rowSize; i++ {
		b[i] = make([]int, colSize)
	}

	return b
}

func Seed(world WorldState, colSize, rowSize int) {
	for i := 0; i < rowSize; i++ {
		for j := 0; j < colSize; j++ {
			world[i][j] = rand.Intn(2) // either 0 or 1
		}
	}
}

func WorldTick(currState WorldState, colSize, rowSize int) WorldState {
	nextState := NewWorldState(colSize, rowSize)

	for i := 0; i < rowSize; i++ {
		for j := 0; j < colSize; j++ {
			neighboursAlive := countLiveNeighbours(currState, colSize, rowSize, i, j)
			currCellState := currState[i][j]
			nextCellState := cellTick(currCellState, neighboursAlive)

			nextState[i][j] = nextCellState
		}
	}
	return nextState
}

func cellTick(currState, neighboursAlive int) int {
	var nextState int
	if isAlive(currState) {
		if neighboursAlive == 2 || neighboursAlive == 3 {
			// cell survives
			nextState = 1
		} else {
			// cell is now dead
			nextState = 0
		}
	} else {
		if neighboursAlive == 3 {
			// cell is now alive
			nextState = 1
		} else {
			// cell remains dead
			nextState = 0
		}
	}
	return nextState
}

func outOfBounds(x, y, offsetX, offsetY, lowBound, highBoundCol, highBoundRow int) bool {
	outOfBoundsRow := (x+offsetX) < lowBound || (x+offsetX) > highBoundRow
	outOfBoundsCol := (y+offsetY) < lowBound || (y+offsetY) > highBoundCol

	return outOfBoundsRow || outOfBoundsCol
}

func isAlive(c int) bool {
	return c == 1
}

func countLiveNeighbours(world WorldState, colSize, rowSize, x, y int) int {
	liveNeighbours := 0

	for offsetX := -1; offsetX < 2; offsetX++ {
		for offsetY := -1; offsetY < 2; offsetY++ {
			if offsetX == 0 && offsetY == 0 {
				continue
			}
			if outOfBounds(x, y, offsetX, offsetY, 0, colSize-1, rowSize-1) {
				continue
			}

			if world[x+offsetX][y+offsetY] == 1 {
				liveNeighbours += 1
			}
		}
	}
	return liveNeighbours
}
