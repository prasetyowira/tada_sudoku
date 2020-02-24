package solver

import (
	"fmt"
	"sort"
)

func (b *Board) IsValid() bool {
	return b.Checker(b.Rows) && b.Checker(b.Cols) && b.Checker(b.Grid)
}

func (b *Board) Checker(data [][]int) bool {
	for row := 0; row < N; row++ {
		sort.Ints(data[row])
		if b.Sum(data[row]) != 45 {
			fmt.Printf("value on row %d is not 45\n", row)
			return false
		}
		for col := 0; col < N; col++ {
			value := data[row][col]
			nextVal := 0
			if col != 8 {
				nextVal = data[row][col+1]
			}

			if value < 0 && value > 10 {
				fmt.Printf("invalid value: %d\n", value)
				return false
			}

			if col != 8 && value == nextVal {
				fmt.Printf("duplicated value: %d with %d\n", value, nextVal)
				return false
			}
		}
	}
	return true
}
