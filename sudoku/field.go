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
			value := field.cells[y][x]
			if value == 0 {
				fmt.Print("_")
			} else {
				fmt.Print(field.cells[y][x])
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func (field Field) Size() int {
	return field.size
}

func (field Field) SetValue(x, y, value int) {
	x = field.toIndex(x)
	field.cells[y][x] = value
}

func (field Field) GetValue(x, y int) int {
	x = field.toIndex(x)
	return field.cells[y][x]
}

func (field Field) GetRow(y int) Line {
	row := slices.Clone(field.cells[y])
	return CreateNewLine(row)
}

func (field Field) GetCol(x int) Line {
	x = field.toIndex(x)

	// collect values
	col := make([]int, field.size)
	for y := 0; y < len(field.cells); y++ {
		col[y] = field.cells[y][x]
	}

	return CreateNewLine(col)
}

func (field Field) GetSubField(x, y int) Field {
	subField := CreateNewField(3)
	x = field.toIndex(x)
	// start indexes
	xs := x * 3
	ys := y * 3
	for iy := ys; iy < ys+3; iy++ {
		for ix := xs; ix < xs+3; ix++ {
			subField.SetValue(ix%3, iy%3, field.GetValue(ix, iy))
		}
	}
	return subField
}

func (field Field) toIndex(letterOrIndex int) int {
	if letterOrIndex >= 'A' && letterOrIndex < 'A'+field.size {
		return letterOrIndex - 'A'
	} else if letterOrIndex >= 'a' && letterOrIndex < 'a'+field.size {
		return letterOrIndex - 'a'
	}
	return letterOrIndex
}
