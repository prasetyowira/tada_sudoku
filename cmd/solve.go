package cmd

import (
	"github.com/prasetyowira/tada_sudoku/internal/sudoku_checker"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var solveCmd = &cobra.Command{
	Use:   "solve",
	Short: "Solving Sudoku",
	Long:  `Solving Sudoku from txt`,
	Run: func(cmd *cobra.Command, args []string) {
		sudoku_checker.InitializeApp()
	},
}

func init() {
	rootCmd.AddCommand(solveCmd)
}
