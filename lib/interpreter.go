package lib

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Printer is used for text output
type Printer func(a ...any) (n int, err error)

// Interpreter reads brainfuck code and implements it
type Interpreter struct {
	state          State
	commands       map[string]Command
	customCommands map[string]Command
	printer        Printer
	reader         io.Reader
}

// NewInterpreter creates the Interpreter struct instance and returns a pointer
func NewInterpreter(printer Printer, reader io.Reader) (i *Interpreter) {
	i = new(Interpreter)
	i.state = State{position: 0, cells: make([]int, Size, Size), size: Size}
	i.commands = DefaultCommands
	i.customCommands = map[string]Command{}

	if printer == nil {
		i.printer = fmt.Println
	} else {
		i.printer = printer
	}

	if reader == nil {
		i.reader = os.Stdin
	} else {
		i.reader = reader
	}

	return
}

// SetCurrentValue sets the given value to the current cell
func (i *Interpreter) SetCurrentValue(value int) {
	i.state.setCurrentValue(value)
}

// GetCurrentValue returns the value of the current cell
func (i *Interpreter) GetCurrentValue() int {
	return i.state.getCurrentValue()
}

// AddCommand adds the given custom command to the customCommands map
func (i *Interpreter) AddCommand(commandName string, command Command) {
	if _, ok := i.commands[commandName]; ok {
		panic(fmt.Sprintf("command %q already exist in default commands", commandName))
	}

	if _, ok := i.customCommands[commandName]; ok {
		panic(fmt.Sprintf("command %q already exist in custom commands", commandName))
	}

	i.customCommands[commandName] = command
}

// RemoveCommand removes the given custom command from the customCommands map
func (i *Interpreter) RemoveCommand(commandName string) {
	if _, ok := i.customCommands[commandName]; !ok {
		panic(fmt.Sprintf("command %q does not exist in custom commands", commandName))
	}

	delete(i.customCommands, commandName)
}

// Execute executes brainfuck code
func (i *Interpreter) Execute(r io.Reader) {
	code, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	i.ExecuteWith(code)
}

// joinCommands joins default commands and custom commands and returns all commands
func (i *Interpreter) joinCommands() (allCommands map[string]Command) {
	allCommands = map[string]Command{}

	for k, v := range i.commands {
		allCommands[k] = v
	}

	for k, v := range i.customCommands {
		allCommands[k] = v
	}

	return
}

// ExecuteWith executes brainfuck code
func (i *Interpreter) ExecuteWith(code []byte) {
	var stack Stack
	idx := 0
	lastIdx := len(code) - 1
	allCommands := i.joinCommands()

	for {
		if idx > lastIdx {
			break
		}

		commandName := string(code[idx])

		if commandFunc, ok := allCommands[commandName]; ok {
			commandFunc(i)
			idx++
		} else {
			if commandName == "[" {
				if i.GetCurrentValue() > 0 {
					i.getClosingBracketPositionOrPanic(idx, lastIdx, code)
					stack.Push(idx)
					idx++
				} else {
					idx = i.getClosingBracketPositionOrPanic(idx, lastIdx, code)
					idx++
				}
			} else if commandName == "]" {
				position, hasValue := stack.Pop()

				if !hasValue {
					panic(
						fmt.Sprintf("SyntaxError: closing bracket came before opening bracket on position %d", idx),
					)
				} else {
					idx = position
				}
			} else {
				panic(fmt.Sprintf("SyntaxError: unknown command %q on position %d", commandName, idx))
			}
		}
	}
}

// getClosingBracketPositionOrPanic tries to find a closing bracket and returns a position of a closing bracket
func (i *Interpreter) getClosingBracketPositionOrPanic(startIndex int, lastIndex int, code []byte) int {
	idx := startIndex
	openingBrackets := 1

	for openingBrackets > 0 {
		idx++

		if idx > lastIndex {
			panic("SyntaxError: no matching pair closing bracket for a loop")
		}

		command := string(code[idx])

		if command == "[" {
			openingBrackets++
		} else if command == "]" {
			openingBrackets--
		}
	}

	return idx
}
