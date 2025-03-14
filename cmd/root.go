package cmd

import (
	"fmt"
	"os"

	"github.com/prasetyowira/tada_sudoku/internal/sudoku_checker"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "./bin/tada_sudoku < ./{input image}",
	Short: "Run Sudoku Checker based on image",
	Long: `Run Sudo Checker based on image or text file`,
	Run: func(cmd *cobra.Command, args []string) {
		sudoku_checker.InitializeAppChecker()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
