package main

import (
	"fmt"
	"sudoku-solver/sudoku"
)

func main() {
	field := sudoku.CreateNewField(3)
	field.Print()
	fmt.Println("Setting cell (A,2) to value 9")
	field.SetValue('C', 2, 9)
	field.Print()
}
