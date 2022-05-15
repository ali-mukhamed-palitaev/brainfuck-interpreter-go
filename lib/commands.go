package lib

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// Command type is a function implementation of a command
type Command func(i *Interpreter)

// IncrementDataPointer (">") increments the data pointer (to point to the next cell to the right)
func IncrementDataPointer(i *Interpreter) {
	i.state.movePointerForward()
}

// DecrementDataPointer ("<") decrements the data pointer (to point to the next cell to the left)
func DecrementDataPointer(i *Interpreter) {
	i.state.movePointerBackwards()
}

// IncrementCellValue ("+") increments (increase by one) the byte at the data pointer
func IncrementCellValue(i *Interpreter) {
	currentValue := i.GetCurrentValue()
	i.SetCurrentValue(currentValue + 1)
}

// DecrementCellValue ("-") decrements (decrease by one) the byte at the data pointer
func DecrementCellValue(i *Interpreter) {
	currentValue := i.GetCurrentValue()
	i.SetCurrentValue(currentValue - 1)
}

// PrintDataPointer (".") outputs the byte at the data pointer
func PrintDataPointer(i *Interpreter) {
	i.printer(i.GetCurrentValue())
}

// SetCellValueFromInput (",") accepts one byte of input, storing its value in the byte at the data pointer
func SetCellValueFromInput(i *Interpreter) {
	reader := bufio.NewReader(i.reader)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	input = strings.TrimSuffix(input, "\n")
	value, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println(err)
		return
	}

	i.SetCurrentValue(value)
}
