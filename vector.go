package goscheme

import (
	"fmt"
	"strings"
)

type Vector []Value

func (s Vector) String() string {
	var a []string
	for _, v := range s {
		a = append(a, v.String())
	}
	return fmt.Sprintf(`[%v]`, strings.Join(a, " "))
}
