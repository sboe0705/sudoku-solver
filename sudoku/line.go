package sudoku

func CreateNewLine(cells []int) Line {
	return Line{len(cells), cells}
}

type Line struct {
	size  int
	cells []int
}

func (line Line) GetValues() []int {
	return getValues(line.cells)
}

func getValues(line []int) []int {
	values := []int{}
	for i := 0; i < len(line); i++ {
		value := line[i]
		if value != 0 {
			values = append(values, value)
		}
	}
	return values
}
