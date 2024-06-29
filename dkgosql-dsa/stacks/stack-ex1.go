package main

import (
	"fmt"
	"strings"
)

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}
func IsVerified(brcstr string) bool {
	var stack Stack // create a stack variable of type Stack
	brackets := map[string]bool{"{}": true, "[]": true, "()": true}
	for _, v := range brcstr {
		stack.Push(string(v))
		if len(stack) >= 2 {
			last := stack[len(stack)-1:]
			seclast := stack[len(stack)-2 : len(stack)-1]
			laststr := strings.Join(last, "")
			seclaststr := strings.Join(seclast, "")
			vt := seclaststr + laststr
			if _, ok := brackets[string(vt)]; ok {
				_, _ = stack.Pop()
				_, _ = stack.Pop()
			}
		}
	}

	if stack.IsEmpty() == true && len(brcstr) > 1 {
		return true
	}

	return false
}
func main() {

	brcstr := "{{[]}}[()]"
	// brcstr = "{"
	isValid := IsVerified(brcstr)
	if isValid {
		fmt.Println("Verification passed")
	} else {
		fmt.Println("Verification failed")
	}

}
