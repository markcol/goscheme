package goscheme

import (
	"fmt"
	"strconv"
)

type valueType int8

const (
	nilValue valueType = iota
	symbolValue
	numberValue
	stringValue
	vectorValue
	procValue
	consValue
)

// Value represents all of the various atom and pointer values
type Value struct {
	typ valueType
	val interface{}
}

var (
	// Nil is the single unique nil value
	Nil = Value{nilValue, nil}
	// False is the single unique false value
	False = Value{symbolValue, "false"}
	// True is the single unique true value
	True = Value{symbolValue, "true"}
)

func (v Value) String() string {
	switch v.typ {
	case numberValue:
		return strconv.FormatFloat(v.val.(float64), 'f', -1, 64)
	case nilValue:
		return "()"
	default:
		return fmt.Sprintf("%v", v.val)
	}
}

// Cons returns a cons cell representation f the current value.
func (v Value) Cons() Cons {
	if v.typ == consValue {
		return *v.val.(*Cons)
	}
	return Cons{&v, &Nil}
}

func (v Value) Eval() (Value, error) {
	switch v.typ {
	case consValue:
		return v.Cons().Execute()
	case symbolValue:
		sym := v.String()
		if v, ok := scope.Get(sym); ok {
			return v, nil
		} else if sym == "true" || sym == "false" {
			return Value{symbolValue, sym}, nil
		} else {
			return Nil, fmt.Errorf("Unbound variable: %v", sym)
		}
	default:
		return v, nil
	}
}

func (v Value) Number() float64 {
	return v.val.(float64)
}
