package hole

import (
	"math/rand/v2"
	"slices"
	"strings"
)

const (
	blockSize = 3
	boardSize = blockSize * blockSize
	cellCount = boardSize * boardSize
)

func printSudoku(board [boardSize][boardSize]int) string {
	var b strings.Builder
	b.Grow(1641) // Length in bytes of a printed Sudoku board.

	for i, row := range board {
		if i == 0 {
			b.WriteString("┏━━━┯━━━┯━━━┳━━━┯━━━┯━━━┳━━━┯━━━┯━━━┓\n")
		} else if i%blockSize == 0 {
			b.WriteString("┣━━━┿━━━┿━━━╋━━━┿━━━┿━━━╋━━━┿━━━┿━━━┫\n")
		} else {
			b.WriteString("┠───┼───┼───╂───┼───┼───╂───┼───┼───┨\n")
		}

		for j, number := range row {
			if j%blockSize == 0 {
				b.WriteRune('┃')
			} else {
				b.WriteRune('│')
			}

			b.WriteByte(' ')
			if number == 0 {
				b.WriteByte(' ')
			} else {
				b.WriteByte(byte('0' + number))
			}
			b.WriteByte(' ')
		}

		b.WriteString("┃\n")
	}

	b.WriteString("┗━━━┷━━━┷━━━┻━━━┷━━━┷━━━┻━━━┷━━━┷━━━┛")

	return b.String()
}

func solveSudoku(board *[boardSize][boardSize]int, cell int, count *int) bool {
	var i, j int

	for {
		if cell == cellCount {
			*count++
			return *count == 2
		}

		i = cell / boardSize
		j = cell % boardSize

		if board[i][j] == 0 {
			break
		}

		cell++
	}

	// Origin of block.
	i0 := i - i%blockSize
	j0 := j - j%blockSize

	// 1 - 9 in random order.
	numbers := rand.Perm(9)
	for i := range numbers {
		numbers[i]++
	}

Numbers:
	for _, number := range numbers {
		// number is already in the row.
		for _, numberInRow := range board[i] {
			if number == numberInRow {
				continue Numbers
			}
		}

		// number is already in the column.
		for _, row := range board {
			if number == row[j] {
				continue Numbers
			}
		}

		// number is already in the block.
		for _, row := range board[i0 : i0+blockSize] {
			for _, numberInRow := range row[j0 : j0+blockSize] {
				if number == numberInRow {
					continue Numbers
				}
			}
		}

		board[i][j] = number

		if solveSudoku(board, cell+1, count) {
			return true
		}
	}

	// No valid number for this cell, let's backtrack.
	board[i][j] = 0

	return false
}

func sudokuBoard(fillIn bool) Answer {
	var board [boardSize][boardSize]int
	var solutions int
	solveSudoku(&board, 0, &solutions)

	var args []string
	out := printSudoku(board)

	// Clear random cells while keeping a unique solution. Only clear up-to 50
	// cells so that brute force solutions are unlikely to time out.
	for _, cell := range rand.Perm(cellCount)[:50] {
		i := cell / boardSize
		j := cell % boardSize

		orig := board[i][j]
		board[i][j] = 0

		// Take a copy as solve will mutate it.
		boardCopy := ([boardSize][boardSize]int)(slices.Clone(board[:]))

		var solutions int
		solveSudoku(&boardCopy, 0, &solutions)

		// Removing this cell creates too many solutions, put it back.
		if solutions == 2 {
			board[i][j] = orig
		}
	}

	if fillIn {
		args = []string{printSudoku(board)}
	} else {
		args = make([]string, boardSize)

		for i, row := range board {
			var b strings.Builder

			for _, number := range row {
				if number == 0 {
					b.WriteByte('_')
				} else {
					b.WriteByte(byte('0' + number))
				}
			}

			args[i] = b.String()
		}
	}

	return Answer{Args: args, Answer: out}
}

var _ = answerFunc("sudoku", func() []Answer { return sudoku(false) })
var _ = answerFunc("sudoku-fill-in", func() []Answer { return sudoku(true) })

func sudoku(fillIn bool) []Answer {
	answers := make([]Answer, 3)
	for i := range answers {
		answers[i] = sudokuBoard(fillIn)
	}
	return answers
}
