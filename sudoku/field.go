package sudoku

import "fmt"

func CreateNewField(size int) Field {
	cells := make([][]int, size*3)
	for y := 0; y < size*3; y++ {
		cells[y] = make([]int, size*3)
	}
	return Field{cells}
}

type Field struct {
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

func (field Field) SetCell(x, y, value int) {
	field.cells[y][x] = value
}
