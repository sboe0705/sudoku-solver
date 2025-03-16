package sudoku

import (
	"sudoku-solver/common"
	"testing"
)

func TestGetValues(t *testing.T) {
	// given
	line := CreateNewLine([]int{0, 1, 2})

	// when
	values := line.GetValues()

	// then
	common.AssertArray(t, []int{1, 2}, values, "wrong line values")
}
