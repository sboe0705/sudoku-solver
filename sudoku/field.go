package sudoku

import (
	"fmt"
	"slices"
)

func CreateNewField(size int) Field {
	cells := make([][]int, size)
	for y := 0; y < size; y++ {
		cells[y] = make([]int, size)
	}
	return Field{size, cells}
}

type Field struct {
	size  int
	cells [][]int
}

func (field Field) Print() {
	for y := 0; y < len(field.cells); y++ {
		for x := 0; x < len(field.cells[y]); x++ {
			fmt.Print(field.cells[y][x])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func (field Field) getRow(y int) []int {
	return slices.Clone(field.cells[y])
}

func (field Field) getCol(x int) []int {
	x = field.toIndex(x)
	col := make([]int, field.size)
	for y := 0; y < len(field.cells); y++ {
		col[y] = field.cells[y][x]
	}
	return col
}

func (field Field) SetValue(x, y, value int) {
	x = field.toIndex(x)
	field.cells[y][x] = value
}

func (field Field) GetValue(x, y int) int {
	x = field.toIndex(x)
	return field.cells[y][x]
}

func (field Field) GetRowValues(y int) []int {
	row := field.getRow(y)
	return getLineValues(row)
}

func (field Field) GetColValues(x int) []int {
	x = field.toIndex(x)
	col := field.getCol(x)
	return getLineValues(col)
}

func getLineValues(line []int) []int {
	values := []int{}
	for i := 0; i < len(line); i++ {
		value := line[i]
		if value != 0 {
			values = append(values, value)
		}
	}
	return values
}

func (field Field) toIndex(letterOrIndex int) int {
	if letterOrIndex >= field.size && letterOrIndex >= 'A' {
		return letterOrIndex - 'A'
	}
	return letterOrIndex
}
