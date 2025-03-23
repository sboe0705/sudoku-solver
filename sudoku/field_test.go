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

func TestGetSubField(t *testing.T) {
	// given
	field := CreateNewField(9)
	field.SetValue('A', 0, 1)
	field.SetValue('B', 1, 2)
	field.SetValue('C', 2, 3)
	field.SetValue('D', 3, 4)
	field.SetValue('E', 4, 5)
	field.SetValue('F', 5, 6)
	field.SetValue('G', 6, 7)
	field.SetValue('H', 7, 8)
	field.SetValue('I', 8, 9)

	field.SetValue('A', 8, 1)
	field.SetValue('B', 7, 2)
	field.SetValue('C', 6, 3)

	// when
	subField := field.GetSubField('A', 0)

	// then
	common.AssertInt(t, 1, subField.GetValue('A', 0), "wrong value in first subfield")
	common.AssertInt(t, 2, subField.GetValue('B', 1), "wrong value in first subfield")
	common.AssertInt(t, 3, subField.GetValue('C', 2), "wrong value in first subfield")

	// when
	subField = field.GetSubField('B', 1)

	// then
	common.AssertInt(t, 4, subField.GetValue('A', 0), "wrong value in second subfield")
	common.AssertInt(t, 5, subField.GetValue('B', 1), "wrong value in second subfield")
	common.AssertInt(t, 6, subField.GetValue('C', 2), "wrong value in second subfield")

	// when
	subField = field.GetSubField('C', 2)

	// then
	common.AssertInt(t, 7, subField.GetValue('A', 0), "wrong value in third subfield")
	common.AssertInt(t, 8, subField.GetValue('B', 1), "wrong value in third subfield")
	common.AssertInt(t, 9, subField.GetValue('C', 2), "wrong value in third subfield")

	// when
	subField = field.GetSubField('A', 2)

	// then
	common.AssertInt(t, 1, subField.GetValue('A', 2), "wrong value in fourth subfield")
	common.AssertInt(t, 2, subField.GetValue('B', 1), "wrong value in fourth subfield")
	common.AssertInt(t, 3, subField.GetValue('C', 0), "wrong value in fourth subfield")
}
