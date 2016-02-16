package goscheme

import (
	"fmt"
	"strings"
)

// Vector is a slice of Values.
type Vector []Value

func (s Vector) String() string {
	var a []string
	for _, v := range s {
		a = append(a, v.String())
	}
	return fmt.Sprintf(`[%v]`, strings.Join(a, " "))
}

// Inspect displays the contents of a Vector for debugging.
func (s Vector) Inspect() string {
	var arr []string
	for _, v := range s {
		arr = append(arr, v.Inspect())
	}
	return fmt.Sprintf(`[%v]`, strings.Join(arr, " "))
}
