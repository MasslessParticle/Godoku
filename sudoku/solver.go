package sudoku
import (
	"strconv"
)

type Puzzle struct {
	Data [9][9]byte
}

func (puzzle Puzzle) String() string {
	solved := ""

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			solved += strconv.Itoa(int(puzzle.Data[i][j]))
		}
	}

	return solved
}

var puzzle Puzzle

func Solve(puzzleString string) Puzzle {
	parsePuzzle(puzzleString)
	backtrackSolve(0, 0)
	return puzzle
}

func parsePuzzle(puzzleString string) {
	puzzle = Puzzle{}

	for index, digit := range puzzleString {
		row := index / 9
		col := index % 9

		if string(digit) != "_" {
			puzzle.Data[row][col] = byte(digit - '0')
		}
	}
}

func backtrackSolve(row int, col int) {
	if row == 9 || puzzleSolved() {
		return
	}

	value := puzzle.Data[row][col]
	nextCol := (col + 1) % 9
	nextRow := nextRow(row, nextCol)

	if value == 0 {
		for i := byte(1); i < 10; i++ {
			if isValid(row, col, i) {
				puzzle.Data[row][col] = i
				backtrackSolve(nextRow, nextCol)

				if puzzleSolved() {
					return
				}
			}
		}

		puzzle.Data[row][col] = 0;
	} else {
		backtrackSolve(nextRow, nextCol)
	}
}


func nextRow(currentRow int, nextCol int) int {
	if nextCol == 0 {
		return currentRow + 1
	}
	return currentRow
}

func isValid(row int, col int, candidate byte) bool {
	return !rowContains(row, candidate) &&
	!colContains(col, candidate) &&
	!blockContains(row, col, candidate)
}

func rowContains(row int, candidate byte) bool {
	for i := 0; i < 9; i++ {
		if puzzle.Data[row][i] == candidate {
			return true
		}
	}

	return false
}

func colContains(col int, candidate byte) bool {
	for i := 0; i < 9; i++ {
		if puzzle.Data[i][col] == candidate {
			return true
		}
	}

	return false
}

func blockContains(row int, col int, candidate byte) bool {
	blockRow := (row / 3) * 3
	blockCol := (col / 3) * 3

	for i := blockRow; i < blockRow + 3; i++ {
		for j := blockCol; j < blockCol + 3; j++ {
			if puzzle.Data[i][j] == candidate {
				return true
			}
		}
	}

	return false
}

func puzzleSolved() bool {
	return rowsSolved() && colsSolved() && blocksSolved()
}

func rowsSolved() bool {
	for i := 0; i < 9; i++ {
		rowSum := 0
		for j := 0; j < 9; j++ {
			rowSum += int(puzzle.Data[i][j]);
		}

		if rowSum != 45 {
			return false
		}
	}

	return true
}

func colsSolved() bool {
	for i := 0; i < 9; i++ {
		colSum := 0
		for j := 0; j < 9; j++ {
			colSum += int(puzzle.Data[j][i]);
		}

		if colSum != 45 {
			return false
		}
	}

	return true
}

func blocksSolved() bool {
	for i := 0; i < 3; i ++ {
		for j := 0; j < 3; j++ {
			blockSum := 0

			for row := i * 3; row < (i * 3) + 3; row++ {
				for col := j * 3; col < (j * 3) + 3; col++ {
					blockSum += int(puzzle.Data[row][col]);
				}
			}

			if blockSum != 45 {
				return false
			}
		}
	}
	return true
}