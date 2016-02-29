# goscheme

A R5RS-compliant Scheme interpreter in Go.

Go is a great language for or building a Scheme interpreter, since
it's portable, garbage collected and fast. The result is a one file
distribution, which makes it easy to deploy. The garbage collection
unloads that work from the interpreter, using Go's GC.

## Setup

### Run the example

```shell
$ goscheme -file $GOPATH/src/github.com/markcol/goscheme/example.scm
```

### Run it interactively

```shell
$ goscheme -i
```

## Syntax

--------------------------------------------------------------------------------

### (quote exp)

Return *exp* literally; do not evaluate it.

```lisp
(quote (a b c))
```

or

```
'(a b c)
```

Returns:

```lisp
(a b c)
```

### (if test conseq alt)

Evaluate *test*; if true, evaluate and return *conseq*; otherwise evaluate
and return *alt*.

```lisp
(if (< 10 20) 
    (+ 1 1)
  (+ 3 3))
```

Returns:

```
2
```

### (set! var exp)

Evaluate *exp* and assign that value to *var*, which must have been
previously defined (with *define* or as a parameter to an enclosing
procedure).

```lisp
(set! x2 (* x x))
```

### (define var exp)

Define a new variable and give it the value of evaluating the
expression *exp*.

```lisp
(define r 3)
(define square (lambda (x) (* x x)))
```

### (lambda (var...) body)

Create a procedure with parameter(s) named *var* and the expression
*body* as the body.

```lisp
(lambda (r) (* 3.141592653 (* r r)))
```

### (begin exp...)

Evaluate each of the expressions in left-to-right order, and return
the final value.

```lisp
(begin (define x 1) (set! x (+ x 1)) (* x 2))
```

Returns:

```lisp
4
```

### (*proc* exp...)

If *proc* is anything other than one of the symbols *if*, *set!*,
*define*, *lambda*, *begin*, or *quote* then it is treated as a
procedure. It is evaluated using the same rules defined here. All the
expressions are evaluated, and then the procedure is called with the
list of expressions as arguments.

```lisp
(square 12)
```

Returns:

```
144
```
