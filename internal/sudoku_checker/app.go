package sudoku_checker

import (
	"fmt"
	"github.com/prasetyowira/tada_sudoku/internal/sudoku_checker/solver"
	"os"
)

func InitializeApp() {
	board, err := solver.GetBoardFrom(os.Stdin)

	if err != nil {
		fmt.Printf("Unable to process the input: %s\n", err)
		os.Exit(1)
	}

	if board.Backtrack() {
		fmt.Println("The Sudoku was solved successfully:")
		fmt.Println(board.String())
	} else {
		fmt.Printf("The Sudoku can't be solved.")
	}
}

func InitializeAppChecker() {
	board, err := solver.GetBoardFrom(os.Stdin)
	if err != nil {
		fmt.Printf("Unable to process the input: %s\n", err)
		os.Exit(1)
	}
	organizedBoard := board.ReorginzeData()
	fmt.Println(organizedBoard)
	if organizedBoard.IsValid() {
		fmt.Println("The Sudoku was solved correctly:")
	} else {
		fmt.Printf("The Sudoku was solved incorrectly.")
	}
}