package lib

// DefaultCommands is a map of default commands
// Size is a number of commands for brainfuck code
var (
	DefaultCommands = map[string]Command{
		">": IncrementDataPointer,
		"<": DecrementDataPointer,
		"+": IncrementCellValue,
		"-": DecrementCellValue,
		".": PrintDataPointer,
		",": SetCellValueFromInput,
	}
	Size = 30000
)
