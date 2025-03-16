package solver

import (
	"slices"
	"sudoku-solver/sudoku"
)

func CreateSolver() Solver {
	return Solver{}
}

type Solver struct {
}

func (solver Solver) Solve(field sudoku.Field) {
	for x := 0; x < field.Size(); x++ {
		for y := 0; y < field.Size(); y++ {
			if field.GetValue(x, y) == 0 {
				singlePossibleValue := getSinglePossibleValue(field, x, y)
				if singlePossibleValue != 0 {
					field.SetValue(x, y, singlePossibleValue)
				}
			}
		}
	}
}

func getSinglePossibleValue(field sudoku.Field, x, y int) int {
	possibleValues := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	colValues := field.GetCol(x).GetValues()
	rowValues := field.GetRow(y).GetValues()
	possibleValues = removeElements(possibleValues, colValues)
	possibleValues = removeElements(possibleValues, rowValues)

	if len(possibleValues) == 1 {
		return possibleValues[0]
	} else {
		return 0
	}
}

func removeElements(array, elementsToRemove []int) []int {
	for i := 0; i < len(elementsToRemove); i++ {
		index := indexOf(array, elementsToRemove[i])
		if index != -1 {
			array = slices.Delete(array, index, index+1)
		}
	}
	return array
}

func indexOf(array []int, element int) int {
	for i := 0; i < len(array); i++ {
		if array[i] == element {
			return i
		}
	}
	return -1
}
