package sudoku

import (
	"sudoku-solver/common"
	"testing"
)

func TestGetRow(t *testing.T) {
	size := 3

	// given
	field := CreateNewField(size)

	// when
	row := field.getRow(0)

	// then
	common.AssertInt(t, size, len(row), "row has wrong length")
	common.AssertArray(t, []int{0, 0, 0}, row, "wrong row values")

	// when
	field.SetValue('A', 0, 1)
	field.SetValue('B', 0, 2)
	field.SetValue('C', 0, 3)
	row = field.getRow(0)

	// then
	common.AssertArray(t, []int{1, 2, 3}, row, "wrong row values")
}

func TestGetColumn(t *testing.T) {
	size := 3

	// given
	field := CreateNewField(size)

	// when
	col := field.getCol('A')

	// then
	common.AssertInt(t, size, len(col), "column has wrong length")
	common.AssertArray(t, []int{0, 0, 0}, col, "wrong column values")

	// when
	field.SetValue('A', 0, 1)
	field.SetValue('A', 1, 2)
	field.SetValue('A', 2, 3)
	col = field.getCol('A')

	// then
	common.AssertArray(t, []int{1, 2, 3}, col, "wrong column values")
}

func TestGetValue(t *testing.T) {
	size := 3
	value := 1

	// given
	field := CreateNewField(size)
	field.SetValue('A', 0, value)

	// when
	common.AssertInt(t, value, field.GetValue('A', 0), "wrong value")
}

func TestGetRowValues(t *testing.T) {
	size := 3

	// given
	field := CreateNewField(size)
	field.SetValue('A', 0, 1)
	field.SetValue('B', 0, 2)

	// when
	row := field.GetRow(0)

	// then
	common.AssertInt(t, size, row.size, "not matching row size")
	common.AssertArray(t, []int{1, 2, 0}, row.cells, "not matching row cells")

	// when
	row = field.GetRow(1)

	// then
	common.AssertArray(t, []int{0, 0, 0}, row.cells, "not matching row cells")
}

func TestColRowValues(t *testing.T) {
	size := 3

	// given
	field := CreateNewField(size)
	field.SetValue('A', 0, 1)
	field.SetValue('A', 1, 2)

	// when
	col := field.GetCol(0)

	// then
	common.AssertInt(t, size, col.size, "not matching column size")
	common.AssertArray(t, []int{1, 2, 0}, col.cells, "not matching column cells")

	// when
	col = field.GetCol(1)
	common.AssertArray(t, []int{0, 0, 0}, col.cells, "not matching column cells")
}
