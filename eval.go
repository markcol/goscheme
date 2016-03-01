package goscheme

import "io"

// Eval evaluates the contents of the reader in the current context.
func Eval(r io.Reader) (Value, error) {
	tokens, err := NewTokens(r)
	if err != nil {
		return Nil, err
	}
	expanded, err := tokens.Expand()
	if err != nil {
		return Nil, err
	}
	parsed, err := expanded.Parse()
	if err != nil {
		return Nil, err
	}
	evaled, err := parsed.Eval()
	if err != nil {
		return Nil, err
	}
	scope.Create("_", evaled)
	return evaled, nil
}
