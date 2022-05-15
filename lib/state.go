package lib

// State is a struct for working with brainfuck code
type State struct {
	position int
	cells    []int
	size     int
}

// getCurrentValue extracts the value of the current cell
func (s *State) getCurrentValue() int {
	return s.cells[s.position]
}

// setCurrentValue set the given value for the current cell
func (s *State) setCurrentValue(value int) {
	s.cells[s.position] = value
}

// movePointerForward increments the data pointer (to point to the next cell to the right)
func (s *State) movePointerForward() {
	if s.position == len(s.cells)-1 {
		s.position = 0
	} else {
		s.position += 1
	}
}

// movePointerBackwards decrements the data pointer (to point to the next cell to the left)
func (s *State) movePointerBackwards() {
	if s.position == 0 {
		s.position = len(s.cells) - 1
	} else {
		s.position -= 1
	}
}
