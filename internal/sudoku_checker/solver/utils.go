package solver

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// N defines the Sudoku's board size
const N = 9

// Board represents the state of a Sudoku.
type Board struct {
	Cells	[N][N]int
	Rows	[][]int
	Cols	[][]int
	Grid	[][]int
}

// GetBoardFrom takes an io.Reader and returns a Board with the input data.
// Invalid or malformed inputs will be rejected.
func GetBoardFrom(source io.Reader) (*Board, error) {
	scanner := bufio.NewScanner(source)
	board := new(Board)
	curRow := 0

	for ; scanner.Scan() && curRow < N; curRow++ {
		// remove all spaces of the current row
		inputRow := strings.Split(strings.TrimSpace(scanner.Text()), " ")

		// check if the input row fits into the board
		if len(inputRow) != N {
			return nil, fmt.Errorf("row %d needs to contain %d fields", curRow, N)
		}

		// convert values and insert them
		for curCol, val := range inputRow {
			digit, err := validateInputCell(val)
			if err != nil {
				return nil, err
			}
			board.Cells[curRow][curCol] = digit
		}
	}

	// check if rows are missing
	if curRow < N {
		return nil, fmt.Errorf("the board contains only %d rows", curRow)
	}

	return board, nil
}

// validateInputCell returns 0 if the cell contains an underscore as placeholder.
// An error will be returned if the conversion fails or the value isn't between 1 and N.
func validateInputCell(val string) (int, error) {
	if val == "_" {
		return 0, nil
	}

	digit, err := strconv.Atoi(val)
	if err != nil || digit < 1 || digit > N {
		return 0, fmt.Errorf("only digits from 1 to %d and _ as placeholder are allowed values", N)
	}

	return digit, nil
}

func (b *Board) ReorginzeData() *Board {
	newBoard := &Board{
		Cells: b.Cells,
		Rows:  [][]int{},
		Cols:  [][]int{},
		Grid:  [][]int{},
	}

	rows := make([][]int, N)
	cols := make([][]int, N)
	grids := make([][]int, N)
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			// Transform column to row
			rows[col] = append(rows[col], b.Cells[row][col])
			cols[row] = append(cols[row], b.Cells[row][col])

			//// Calculate grid identifiers
			gridRow := row / 3
			gridCol := col / 3
			gridIndex := gridRow * 3 + gridCol
			cellData := b.Cells[row][col]
			grids[gridIndex] = append(grids[gridIndex], cellData)
		}
	}
	newBoard.Rows = rows
	newBoard.Cols = cols
	newBoard.Grid = grids
	return newBoard
}

func (b *Board) Sum(listNumber []int) int {
	sum := 0
	for i:=0;i<len(listNumber);i++{
		sum += listNumber[i]
	}
	return sum
}

// PrintBoard returns a formatted version of the  Sudoku's
// current state. All digits are seperated by a whitespace.
func (b *Board) String() string {
	var output string

	for i, row := range b.Cells {
		str := fmt.Sprint(row)
		output += str[1 : len(str)-1] // trim the array's backets

		if i+1 < N {
			output += "\n"
		}
	}

	return output
}
