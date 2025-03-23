package solver_test

import (
	"sudoku-solver/common"
	"sudoku-solver/solver"
	"sudoku-solver/sudoku"
	"testing"
)

func TestSolveWithOneMissingNumberInRow(t *testing.T) {
	// given
	field := sudoku.CreateNewField(9)
	field.SetValue('A', 0, 1)
	field.SetValue('B', 0, 2)
	field.SetValue('C', 0, 3)
	field.SetValue('D', 0, 4)
	field.SetValue('E', 0, 5)
	field.SetValue('F', 0, 6)
	field.SetValue('G', 0, 7)
	field.SetValue('H', 0, 8)

	// when
	solver.CreateSolver().Solve(field)

	// then
	common.AssertInt(t, 9, field.GetValue('I', 0), "last possible number in row not correctly solved")
}

func TestSolveWithOneMissingNumberInColumnRow(t *testing.T) {
	// given
	field := sudoku.CreateNewField(9)
	field.SetValue('A', 0, 1)
	field.SetValue('A', 1, 2)
	field.SetValue('A', 2, 3)
	field.SetValue('A', 3, 4)
	field.SetValue('A', 4, 5)
	field.SetValue('A', 5, 6)
	field.SetValue('A', 6, 7)
	field.SetValue('A', 7, 8)

	// when
	solver.CreateSolver().Solve(field)

	// then
	common.AssertInt(t, 9, field.GetValue('A', 8), "last possible number in column not correctly solved")
}

func TestSolveWithOneMissingNumberInSubField(t *testing.T) {
	// given
	field := sudoku.CreateNewField(9)
	field.SetValue('A', 0, 1)
	field.SetValue('A', 1, 2)
	field.SetValue('A', 2, 3)
	field.SetValue('B', 0, 4)
	field.SetValue('B', 1, 5)
	field.SetValue('B', 2, 6)
	field.SetValue('C', 0, 7)
	field.SetValue('C', 1, 8)

	// when
	solver.CreateSolver().Solve(field)

	// then
	common.AssertInt(t, 9, field.GetValue('C', 2), "last possible number in sub field not correctly solved")
}

func TestSolveWithOneMissingNumberInRowColumnAndSubField(t *testing.T) {
	// given
	field := sudoku.CreateNewField(9)
	field.SetValue('A', 2, 1)
	field.SetValue('E', 2, 2)
	field.SetValue('I', 2, 3)

	field.SetValue('C', 1, 4)
	field.SetValue('C', 4, 5)
	field.SetValue('C', 8, 6)

	field.SetValue('A', 0, 7)
	field.SetValue('B', 1, 8)

	// when
	solver.CreateSolver().Solve(field)

	// then
	common.AssertInt(t, 9, field.GetValue('C', 2), "last possible number in row, column and sub field not correctly solved")
}
