package main

import (
	"fmt"
	"sudoku-solver/sudoku"
)

func main() {
	field := sudoku.CreateNewField(3)
	field.Print()
	fmt.Println("Setting cell (1,2) to value 9")
	field.SetCell(1, 2, 9)
	field.Print()
}
