package lib

// Stack is an implementation of a stack data structure
type Stack []int

// IsEmpty checks if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push appends the new value to the end of the stack
func (s *Stack) Push(str int) {
	*s = append(*s, str)
}

// Pop removes and returns top element of stack. It returns false if stack is empty.
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}
