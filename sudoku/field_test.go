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
	row := field.GetRow(0)

	// then
	common.AssertInt(t, size, len(row.cells), "row has wrong length")
	common.AssertArray(t, []int{0, 0, 0}, row.cells, "wrong row values")

	// when
	field.SetValue('A', 0, 1)
	field.SetValue('B', 0, 2)
	field.SetValue('C', 0, 3)
	row = field.GetRow(0)

	// then
	common.AssertArray(t, []int{1, 2, 3}, row.cells, "wrong row values")
}

func TestGetColumn(t *testing.T) {
	size := 3

	// given
	field := CreateNewField(size)

	// when
	col := field.GetCol('A')

	// then
	common.AssertInt(t, size, len(col.cells), "column has wrong length")
	common.AssertArray(t, []int{0, 0, 0}, col.cells, "wrong column values")

	// when
	field.SetValue('A', 0, 1)
	field.SetValue('A', 1, 2)
	field.SetValue('A', 2, 3)
	col = field.GetCol('A')

	// then
	common.AssertArray(t, []int{1, 2, 3}, col.cells, "wrong column values")
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
	col := field.GetCol('A')

	// then
	common.AssertInt(t, size, col.size, "not matching column size")
	common.AssertArray(t, []int{1, 2, 0}, col.cells, "not matching column cells")

	// when
	col = field.GetCol('B')

	// then
	common.AssertArray(t, []int{0, 0, 0}, col.cells, "not matching column cells")
}
