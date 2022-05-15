package lib

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type PPrinter struct {
	values []any
}

func (pp *PPrinter) printFunc(a ...any) (n int, err error) {
	pp.values = append(pp.values, a...)
	return
}

func expectError(t *testing.T, expectedError string) {
	if r := recover(); r != nil {
		assert.Equal(t, expectedError, r)
	}
}

func TestInterpreter_ExecuteDotCommandPrintsOutCurrentCellValue(t *testing.T) {
	pp := PPrinter{}
	interpreter := NewInterpreter(pp.printFunc, strings.NewReader("20"))
	reader := strings.NewReader(".[...+[..-]..]")
	interpreter.Execute(reader)
	assert.Equal(t, []any{0}, pp.values)
}

func TestInterpreter_ExecuteShouldAllowToAddCustomCommand(t *testing.T) {
	doubleCommand := func(i *Interpreter) {
		i.SetCurrentValue(i.GetCurrentValue() * 2)
	}
	pp := PPrinter{}
	interpreter := NewInterpreter(pp.printFunc, strings.NewReader("20"))
	reader := strings.NewReader(".+.*.+.*.")
	interpreter.AddCommand("*", doubleCommand)
	interpreter.Execute(reader)
	assert.Equal(t, []any{0, 1, 2, 3, 6}, pp.values)
}

func TestInterpreter_ExecutePanicIfTryToAddExistingDefaultCommand(t *testing.T) {
	expectedError := "command \".\" already exist in default commands"
	defer expectError(t, expectedError)
	doubleCommand := func(i *Interpreter) {
		i.SetCurrentValue(i.GetCurrentValue() * 2)
	}
	pp := PPrinter{}
	interpreter := NewInterpreter(pp.printFunc, strings.NewReader("20"))
	reader := strings.NewReader(".+.")
	interpreter.AddCommand(".", doubleCommand)
	interpreter.Execute(reader)
}

func TestInterpreter_ExecutePanicIfTryToAddExistingCustomCommand(t *testing.T) {
	expectedError := "command \"*\" already exist in custom commands"
	defer expectError(t, expectedError)
	doubleCommand := func(i *Interpreter) {
		i.SetCurrentValue(i.GetCurrentValue() * 2)
	}
	pp := PPrinter{}
	interpreter := NewInterpreter(pp.printFunc, strings.NewReader("20"))
	reader := strings.NewReader(".+.")
	interpreter.AddCommand("*", doubleCommand)
	interpreter.AddCommand("*", doubleCommand)
	interpreter.Execute(reader)
}

func TestInterpreter_ExecutePanicIfTryToRemoveNonExistingCustomCommand(t *testing.T) {
	expectedError := "command \"*\" does not exist in custom commands"
	defer expectError(t, expectedError)
	pp := PPrinter{}
	interpreter := NewInterpreter(pp.printFunc, strings.NewReader("20"))
	reader := strings.NewReader(".+.")
	interpreter.RemoveCommand("*")
	interpreter.Execute(reader)
}

func TestInterpreter_ExecuteAllowToRemoveExistingCustomCommand(t *testing.T) {
	expectedError := "SyntaxError: unknown command \"*\" on position 2"
	defer expectError(t, expectedError)
	doubleCommand := func(i *Interpreter) {
		i.SetCurrentValue(i.GetCurrentValue() * 2)
	}
	pp := PPrinter{}
	interpreter := NewInterpreter(pp.printFunc, strings.NewReader("20"))
	reader := strings.NewReader(".+*.")
	interpreter.AddCommand("*", doubleCommand)
	interpreter.RemoveCommand("*")
	interpreter.Execute(reader)
}

func TestInterpreter_ExecuteUnknownCommand(t *testing.T) {
	expectedError := "SyntaxError: unknown command \":\" on position 2"
	defer expectError(t, expectedError)
	pp := PPrinter{}
	interpreter := NewInterpreter(pp.printFunc, strings.NewReader("20"))
	reader := strings.NewReader(".+:")
	interpreter.Execute(reader)
}

func TestInterpreter_ExecuteNoMatchingClosingBracketZeroValueBeforeLoop(t *testing.T) {
	expectedError := "SyntaxError: no matching pair closing bracket for a loop"
	defer expectError(t, expectedError)
	pp := PPrinter{}
	interpreter := NewInterpreter(pp.printFunc, strings.NewReader("20"))
	reader := strings.NewReader(".[")
	interpreter.Execute(reader)
}

func TestInterpreter_ExecuteNoMatchingClosingBracketNonZeroValueBeforeLoop(t *testing.T) {
	expectedError := "SyntaxError: no matching pair closing bracket for a loop"
	defer expectError(t, expectedError)
	pp := PPrinter{}
	interpreter := NewInterpreter(pp.printFunc, strings.NewReader("20"))
	reader := strings.NewReader("+[..[.-]")
	interpreter.Execute(reader)
}

func TestInterpreter_ExecuteClosingBracketBeforeOpeningOne(t *testing.T) {
	expectedError := "SyntaxError: closing bracket came before opening bracket on position 5"
	defer expectError(t, expectedError)
	pp := PPrinter{}
	interpreter := NewInterpreter(pp.printFunc, strings.NewReader("20"))
	reader := strings.NewReader("+[.-]]")
	interpreter.Execute(reader)
}
