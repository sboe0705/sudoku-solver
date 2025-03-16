package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sudoku-solver/solver"
	"sudoku-solver/sudoku"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	field := createSampleField()
	solver := solver.CreateSolver()

	for {
		clearConsole()
		field.Print()

		fmt.Println("\nOptions:")
		fmt.Println("1. Set value")
		fmt.Println("2. Solve")
		fmt.Println("0. Exit")
		fmt.Print("Select an option: ")

		option := readString()

		switch option {
		case "1":
			fmt.Print("Input (Format: \"X Y ?\"): ")
			params := readParams()
			field.SetValue(params[0], params[1], params[2])
		case "2":
			solver.Solve(field)
		case "0":
			os.Exit(0)
		}
	}
}

func createSampleField() sudoku.Field {
	field := sudoku.CreateNewField(9)
	field.SetValue('A', 0, 1)
	field.SetValue('A', 2, 6)
	field.SetValue('A', 3, 4)
	field.SetValue('A', 4, 7)
	field.SetValue('A', 8, 9)

	field.SetValue('B', 3, 5)
	field.SetValue('B', 5, 1)
	field.SetValue('B', 6, 4)
	field.SetValue('B', 8, 8)

	field.SetValue('C', 3, 8)
	field.SetValue('C', 4, 3)
	field.SetValue('C', 8, 1)

	field.SetValue('D', 1, 2)
	field.SetValue('D', 6, 3)
	field.SetValue('D', 7, 5)

	field.SetValue('E', 2, 4)
	field.SetValue('E', 4, 2)
	field.SetValue('E', 6, 8)

	field.SetValue('F', 1, 1)
	field.SetValue('F', 2, 7)
	field.SetValue('F', 7, 4)

	field.SetValue('G', 0, 2)
	field.SetValue('G', 4, 5)
	field.SetValue('G', 5, 4)

	field.SetValue('H', 0, 6)
	field.SetValue('H', 2, 3)
	field.SetValue('H', 3, 7)
	field.SetValue('H', 5, 8)

	field.SetValue('I', 0, 4)
	field.SetValue('I', 4, 6)
	field.SetValue('I', 5, 3)
	field.SetValue('I', 6, 1)
	field.SetValue('I', 8, 5)

	return field
}

func readParams() []int {
	input := readString()
	var a, b, c int
	_, err := fmt.Sscanf(input, "%d %d %d", &a, &b, &c)
	if err != nil {
		fmt.Println("Fehler beim Parsen:", err)
		return nil
	}
	return []int{a, b, c}
}

func readString() string {
	value, _ := reader.ReadString('\n')
	value = value[:len(value)-1] // remove new-line
	return value
}

func clearConsole() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Print("\033[H\033[2J")
	}
}
