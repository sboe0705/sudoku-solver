package main

import (
	"sudoku-solver/sudoku"
)

func main() {
	samplField := createSampleField()
	samplField.Print()
}

func createSampleField() sudoku.Field {
	field := sudoku.CreateNewField(9)
	field.SetValue('A', 0, 1)
	field.SetValue('A', 2, 6)
	field.SetValue('A', 3, 4)
	field.SetValue('A', 4, 7)
	field.SetValue('A', 8, 9)

	field.SetValue('B', 3, 5)
	field.SetValue('B', 5, 1)
	field.SetValue('B', 6, 4)
	field.SetValue('B', 8, 8)

	field.SetValue('C', 3, 8)
	field.SetValue('C', 4, 3)
	field.SetValue('C', 8, 1)

	field.SetValue('D', 1, 2)
	field.SetValue('D', 6, 3)
	field.SetValue('D', 7, 5)

	field.SetValue('E', 2, 4)
	field.SetValue('E', 4, 2)
	field.SetValue('E', 6, 8)

	field.SetValue('F', 1, 1)
	field.SetValue('F', 2, 7)
	field.SetValue('F', 7, 4)

	field.SetValue('G', 0, 2)
	field.SetValue('G', 4, 5)
	field.SetValue('G', 5, 4)

	field.SetValue('H', 0, 6)
	field.SetValue('H', 2, 3)
	field.SetValue('H', 3, 7)
	field.SetValue('H', 5, 8)

	field.SetValue('I', 0, 4)
	field.SetValue('I', 4, 6)
	field.SetValue('I', 5, 3)
	field.SetValue('I', 6, 1)
	field.SetValue('I', 8, 5)

	return field
}
