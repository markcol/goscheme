# Scheme R5RS Compliance

- [Predicate Uniqueness](http://www.schemers.org/Documents/Standards/R5RS/HTML/r5rs-Z-H-2.html#%_toc_%_sec_3.2)

  - boolean?
  - symbol?
  - char?
  - vector?
  - pair?
  - number?
  - string?
  - port?

- [Tail Recursion](http://www.schemers.org/Documents/Standards/R5RS/HTML/r5rs-Z-H-2.html#%_toc_%_sec_3.5)
- [Conditionals](http://www.schemers.org/Documents/Standards/R5RS/HTML/r5rs-Z-H-2.html#%_toc_%_sec_4.2.1)

  - cond _(library)_
  - case _(library)_
  - and  _(library)_
  - or  _(library)_

- [Binding Constructs](http://www.schemers.org/Documents/Standards/R5RS/HTML/r5rs-Z-H-2.html#%_toc_%_sec_4.2.2)

  - let _(library)_
  - let* _(library)_
  - letrec _(library)_

- [Sequencing](http://www.schemers.org/Documents/Standards/R5RS/HTML/r5rs-Z-H-2.html#%_toc_%_sec_4.2.3)

  - begin _(library)_

- [Iteration](http://www.schemers.org/Documents/Standards/R5RS/HTML/r5rs-Z-H-2.html#%_toc_%_sec_4.2.4)

  - do _(library)_
  - let _(library)_

- [Delayed Evalutation](http://www.schemers.org/Documents/Standards/R5RS/HTML/r5rs-Z-H-2.html#%_toc_%_sec_4.2.5)

  - delay _(library)_

- [Quasiquotation](http://www.schemers.org/Documents/Standards/R5RS/HTML/r5rs-Z-H-2.html#%_toc_%_sec_4.2.6)

  - quasiquote
  - qq

- [Macros](http://www.schemers.org/Documents/Standards/R5RS/HTML/r5rs-Z-H-2.html#%_toc_%_sec_4.3.1)

  - let-syntax
  - letrec-syntax
  - syntax-rules

- [Equivalence Predicates](http://www.schemers.org/Documents/Standards/R5RS/HTML/r5rs-Z-H-2.html#%_toc_%_sec_6.1)

  - eqv?
  - eq?
  - equal?

- [Numerical Operations](http://www.schemers.org/Documents/Standards/R5RS/HTML/r5rs-Z-H-2.html#%_toc_%_sec_6.2.5)

  - number?
  - complex?
  - real?
  - rational?
  - integer?
  - exact?
  - inexact?
  - zero?
  - positive?
  - negative?
  - odd?
  - even?
  - max
  - min
  - =, <, >, <=, >=, +, -, x (asterisk), /,
  - quotient
  - remainder
  - modulo
  - gcd _(library)_
  - lcm _(library)_
  - numerator
  - denominator
  - floor
  - ceiling
  - truncate
  - round
  - exp
  - log
  - sin
  - cos
  - tan
  - asin
  - acos
  - atan
  - atan x y
  - sqrt
  - expt
  - make-rectangular
  - make-polar
  - real-part
  - imag-part
  - magnitude
  - angle
  - exact->inexact
  - inexact->exact
  - number->string
  - string->number

- [Booleans](http://www.schemers.org/Documents/Standards/R5RS/HTML/r5rs-Z-H-2.html#%_toc_%_sec_6.3.1)

  - not _(library)_
  - boolean? _(library)_

- [Pairs and Lists](http://www.schemers.org/Documents/Standards/R5RS/HTML/r5rs-Z-H-2.html#%_toc_%_sec_6.3.2)

  - pair?
  - cons
  - car
  - cdr
  - set-car!
  - set-cdr!
  - null? _(library)_
  - list? _(library)_
  - list _(library)_
  - length _(library)_
  - append _(library)_
  - reverse _(library)_
  - list-tail _(library)_
  - list-ref _(library)_
  - memq _(library)_
  - memv _(library)_
  - member _(library)_
  - assq _(library)_
  - assv _(library)_
  - assoc _(library)_

- [Symbols](http://www.schemers.org/Documents/Standards/R5RS/HTML/r5rs-Z-H-2.html#%_toc_%_sec_6.3.3)

  - symbol?
  - symbol->string
  - string->symbol

- [Characters](http://www.schemers.org/Documents/Standards/R5RS/HTML/r5rs-Z-H-2.html#%_toc_%_sec_6.3.4)

  - char?
  - char=?
  - char<?
  - char>?
  - char<=?
  - char>=?
  - char-ci=? _(library)_
  - char-ci<? _(library)_
  - char-ci>? _(library)_
  - char-ci<=? _(library)_
  - char-ci>=? _(library)_
  - char-alphabetic? _(library)_
  - char-numeric? _(library)_
  - char-whitespace? _(library)_
  - char-upper-case? _(library)_
  - char-lower-case? _(library)_
  - char->integer
  - integer->char
  - char-upcase _(library)_
  - char-downcase _(library)_

- [Strings](http://www.schemers.org/Documents/Standards/R5RS/HTML/r5rs-Z-H-2.html#%_toc_%_sec_6.3.5)

  - string?
  - make-string
  - string _(library)_
  - string-length
  - string-ref
  - string-set!
  - string=? _(library)_
  - string-ci=? _(library)_
  - string<? _(library)_
  - string>? _(library)_
  - string<=? _(library)_
  - string>=? _(library)_
  - string-ci<? _(library)_
  - string-ci>? _(library)_
  - string-ci<=? _(library)_
  - string-ci>=? _(library)_
  - substring _(library)_
  - string-append _(library)_
  - string->list _(library)_
  - list->string _(library)_
  - string-copy _(library)_
  - string-fill! _(library)_

- [Vectors](http://www.schemers.org/Documents/Standards/R5RS/HTML/r5rs-Z-H-2.html#%_toc_%_sec_6.3.6)

  - vector?
  - make-vector
  - vector _(library)_
  - vector-length
  - vector-ref
  - vector-set!
  - vector->list _(library)_
  - list-vector _(library)_
  - vector-fill _(library)_

- [Control Features](http://www.schemers.org/Documents/Standards/R5RS/HTML/r5rs-Z-H-2.html#%_toc_%_sec_6.4)

  - procedure?
  - apply
  - map _(library)_
  - for-each _(library)_
  - force _(library)_
  - call-with-current-continuation
  - values
  - call-with-values
  - dynamic-wind

- [Eval](http://www.schemers.org/Documents/Standards/R5RS/HTML/r5rs-Z-H-2.html#%_toc_%_sec_6.5)

  - eval
  - scheme-report-environment
  - null-environment
  - interaction-environment _(optional)_

- [Ports](http://www.schemers.org/Documents/Standards/R5RS/HTML/r5rs-Z-H-2.html#%_toc_%_sec_6.6.1)

  - call-with-input-file
  - call-with-output-file
  - input-port?
  - output-port?
  - current-input-port
  - current-output-port
  - with-input-from-file _(optional)_
  - with-output-to-file _(optional)_
  - open-input-file
  - open-output-file
  - close-input-port
  - close-output-port

- [Input](http://www.schemers.org/Documents/Standards/R5RS/HTML/r5rs-Z-H-2.html#%_toc_%_sec_6.6.2)

  - read
  - read-char
  - peek-char
  - eof-object?
  - char-ready?

- [Output](http://www.schemers.org/Documents/Standards/R5RS/HTML/r5rs-Z-H-2.html#%_toc_%_sec_6.6.3)

  - write
  - display
  - newline
  - write-char
  - current-output-port

- [System Interface](http://www.schemers.org/Documents/Standards/R5RS/HTML/r5rs-Z-H-2.html#%_toc_%_sec_6.6.4)

  - load _(optional)_
  - transcript-on _(optional)_
  - transcript-off _(optional)_

# Implementation Changes

- Replace usage of Cons.Sexp() w/ Cons.Map and Cons.Len
- Replace evalValue(Value) w/ Eval(Cons)
- Use interface for Value.val instead of interface{}
- Move tokens.Expand() to be called before Parse(), like tokens.Expand().Parse() sort of
- Rename Sexp => Vector
- Add _cons_ command
- Remove typecasting commands on Value
- Remove naked returns
- Add methods to Cons
- Add support for Line/pos in error messages

  - Store pos in token
  - Store token in value
  - Calculate line/pos lazily when outputting

- Add short forms

  - (define (inc x) (+ x 1)) => (define inc (lambda (x) (+ x 1)))
  - (let ((a 1) (b 2)) (+ a b)) => ((lambda (a b) (+ a b)) 1 2)
  - (list a b) => (cons a (cons b ()))
  - (a . b) => (cons a b)

# Language Extensions

- Implement tapping into Go packages?

  - Use [pkgreflect](https://github.com/ungerik/pkgreflect) to create maps like [Clojure Java interop](http://clojure.org/java_interop)

    ```lisp
    (fmt/Printf "Hello, world!")
    (.Body (first (http/Get "[http://example.com/](http://example.com/)")))
    (. (. (. http (Get "[http://example.com/](http://example.com/)")) (first)) (Body))
    (.. (http (Get "[http://example.com/](http://example.com/)")) (first) (Body))
    (-> (http (Get "[http://example.com/](http://example.com/)")) (first) (Body))
    (doto (http/Get "[http://example.com/](http://example.com/)") (first) (Body))
    ```

- Support for [modules](http://www.htus.org/Book/Staging/how-to-use-modules/)?

# Sceme Requests for Implementation (SFRIs)

- [SRFI 0](http://srfi.schemers.org/srfi-0/srfi-0.html): Feature-based conditional expansion construct
- [SRFI 1](http://srfi.schemers.org/srfi-1/srfi-1.html): List Library
- [SRFI 2](http://srfi.schemers.org/srfi-2/srfi-2.html): AND-LET_: an AND with local bindings, a guarded LET_ special form
- [SRFI 4](http://srfi.schemers.org/srfi-4/srfi-4.html): Homogeneous numeric vector datatypes
- [SRFI 5](http://srfi.schemers.org/srfi-5/srfi-5.html): A compatible let form with signatures and rest arguments
- [SRFI 6](http://srfi.schemers.org/srfi-6/srfi-6.html): Basic String Ports
- [SRFI 7](http://srfi.schemers.org/srfi-7/srfi-7.html): Feature-based program configuration language
- [SRFI 8](http://srfi.schemers.org/srfi-8/srfi-8.html): receive: Binding to multiple values
- [SRFI 9](http://srfi.schemers.org/srfi-9/srfi-9.html): Defining Record Types
- [SRFI 10](http://srfi.schemers.org/srfi-10/srfi-10.html): Sharp-Comma External Form
- [SRFI 11](http://srfi.schemers.org/srfi-11/srfi-11.html): Syntax for receiving multiple values
- [SRFI 13](http://srfi.schemers.org/srfi-13/srfi-13.html): String Library
- [SRFI 14](http://srfi.schemers.org/srfi-14/srfi-14.html): Character-Set Library
- [SRFI 16](http://srfi.schemers.org/srfi-16/srfi-16.html): Syntax for procedures of variable arity
- [SRFI 17](http://srfi.schemers.org/srfi-17/srfi-17.html): Generalized set!
- [SRFI 18](http://srfi.schemers.org/srfi-18/srfi-18.html): Multithreading support
- [SRFI 19](http://srfi.schemers.org/srfi-19/srfi-19.html): Time Data Types and Procedures
- [SRFI 21](http://srfi.schemers.org/srfi-21/srfi-21.html): Real-time multithreading support
- [SRFI 22](http://srfi.schemers.org/srfi-22/srfi-22.html): Running Scheme Scripts on Unix
- [SRFI 23](http://srfi.schemers.org/srfi-23/srfi-23.html): Error reporting mechanism
- [SRFI 25](http://srfi.schemers.org/srfi-25/srfi-25.html): Multi-dimensional Array Primitives
- [SRFI 26](http://srfi.schemers.org/srfi-26/srfi-26.html): Notation for Specializing Parameters without Currying
- [SRFI 27](http://srfi.schemers.org/srfi-27/srfi-27.html): Sources of Random Bits
- [SRFI 28](http://srfi.schemers.org/srfi-28/srfi-28.html): Basic Format Strings
- [SRFI 29](http://srfi.schemers.org/srfi-29/srfi-29.html): Localization
- [SRFI 30](http://srfi.schemers.org/srfi-30/srfi-30.html): Nested Multi-line Comments
- [SRFI 31](http://srfi.schemers.org/srfi-31/srfi-31.html): A special form for recursive evaluation
- [SRFI 34](http://srfi.schemers.org/srfi-34/srfi-34.html): Exception Handling for Programs
- [SRFI 35](http://srfi.schemers.org/srfi-35/srfi-35.html): Conditions
- [SRFI 36](http://srfi.schemers.org/srfi-36/srfi-36.html): I/O Conditions
- [SRFI 37](http://srfi.schemers.org/srfi-37/srfi-37.html): args-fold: a program argument processor
- [SRFI 38](http://srfi.schemers.org/srfi-38/srfi-38.html): External Representation for Data With Shared Structure
- [SRFI 39](http://srfi.schemers.org/srfi-39/srfi-39.html): Parameter objects
- [SRFI 40](http://srfi.schemers.org/srfi-40/srfi-40.html): A Library of Streams (deprecated)
- [SRFI 41](http://srfi.schemers.org/srfi-41/srfi-41.html): Streams
- [SRFI 42](http://srfi.schemers.org/srfi-42/srfi-42.html): Eager Comprehensions
- [SRFI 43](http://srfi.schemers.org/srfi-43/srfi-43.html): Vector Library
- [SRFI 44](http://srfi.schemers.org/srfi-44/srfi-44.html): Collections
- [SRFI 45](http://srfi.schemers.org/srfi-45/srfi-45.html): Primitives for expressing iterative lazy algorithms
- [SRFI 46](http://srfi.schemers.org/srfi-46/srfi-46.html): Basic Syntax-rules Extensions
- [SRFI 47](http://srfi.schemers.org/srfi-47/srfi-47.html): Array
- [SRFI 48](http://srfi.schemers.org/srfi-48/srfi-48.html): Intermediate Format Strings
- [SRFI 49](http://srfi.schemers.org/srfi-49/srfi-49.html): Indentation-sensitive syntax
- [SRFI 51](http://srfi.schemers.org/srfi-51/srfi-51.html): Handling rest list
- [SRFI 54](http://srfi.schemers.org/srfi-54/srfi-54.html): Formatting
- [SRFI 55](http://srfi.schemers.org/srfi-55/srfi-55.html): require-extension
- [SRFI 57](http://srfi.schemers.org/srfi-57/srfi-57.html): Records
- [SRFI 58](http://srfi.schemers.org/srfi-58/srfi-58.html): Array Notation
- [SRFI 59](http://srfi.schemers.org/srfi-59/srfi-59.html): Vicinity
- [SRFI 60](http://srfi.schemers.org/srfi-60/srfi-60.html): Integers as Bits
- [SRFI 61](http://srfi.schemers.org/srfi-61/srfi-61.html): A more general cond clause
- [SRFI 62](http://srfi.schemers.org/srfi-62/srfi-62.html): S-expression comments
- [SRFI 63](http://srfi.schemers.org/srfi-63/srfi-63.html): Homogeneous and Heterogeneous Arrays
- [SRFI 64](http://srfi.schemers.org/srfi-64/srfi-64.html): A Scheme API for test suites
- [SRFI 66](http://srfi.schemers.org/srfi-66/srfi-66.html): Octet Vectors
- [SRFI 67](http://srfi.schemers.org/srfi-67/srfi-67.html): Compare Procedures
- [SRFI 69](http://srfi.schemers.org/srfi-69/srfi-69.html): Basic hash tables
- [SRFI 70](http://srfi.schemers.org/srfi-70/srfi-70.html): Numbers
- [SRFI 71](http://srfi.schemers.org/srfi-71/srfi-71.html): LET-syntax for multiple values
- [SRFI 72](http://srfi.schemers.org/srfi-72/srfi-72.html): Simple hygienic macros
- [SRFI 74](http://srfi.schemers.org/srfi-74/srfi-74.html): Octet-Addressed Binary Blocks
- [SRFI 78](http://srfi.schemers.org/srfi-78/srfi-78.html): Lightweight testing
- [SRFI 86](http://srfi.schemers.org/srfi-86/srfi-86.html): MU and NU simulating VALUES & CALL-WITH-VALUES, and their related LET-syntax
- [SRFI 87](http://srfi.schemers.org/srfi-87/srfi-87.html): => in case clauses
- [SRFI 88](http://srfi.schemers.org/srfi-88/srfi-88.html): Keyword Objects
- [SRFI 89](http://srfi.schemers.org/srfi-89/srfi-89.html): Optional and named parameters
- [SRFI 90](http://srfi.schemers.org/srfi-90/srfi-90.html): Extensible hash table constructor
- [SRFI 94](http://srfi.schemers.org/srfi-94/srfi-94.html): Type-Restricted Numerical Functions
- [SRFI 95](http://srfi.schemers.org/srfi-95/srfi-95.html): Sorting and Merging
- [SRFI 96](http://srfi.schemers.org/srfi-96/srfi-96.html): SLIB Prerequisites
- [SRFI 97](http://srfi.schemers.org/srfi-97/srfi-97.html): SRFI Libraries
- [SRFI 98](http://srfi.schemers.org/srfi-98/srfi-98.html): An interface to access environment variables
- [SRFI 99](http://srfi.schemers.org/srfi-99/srfi-99.html): ERR5RS Records
- [SRFI 100](http://srfi.schemers.org/srfi-100/srfi-100.html): define-lambda-object
- [SRFI 101](http://srfi.schemers.org/srfi-101/srfi-101.html): Purely Functional Random-Access Pairs and Lists
- [SRFI 105](http://srfi.schemers.org/srfi-105/srfi-105.html): Curly-infix-expressions
- [SRFI 106](http://srfi.schemers.org/srfi-106/srfi-106.html): Basic socket interface
- [SRFI 107](http://srfi.schemers.org/srfi-107/srfi-107.html): XML reader syntax
- [SRFI 108](http://srfi.schemers.org/srfi-108/srfi-108.html): Named quasi-literal constructors
- [SRFI 109](http://srfi.schemers.org/srfi-109/srfi-109.html): Extended string quasi-literals
- [SRFI 110](http://srfi.schemers.org/srfi-110/srfi-110.html): Sweet-expressions (t-expressions)
- [SRFI 111](http://srfi.schemers.org/srfi-111/srfi-111.html): Boxes
- [SRFI 112](http://srfi.schemers.org/srfi-112/srfi-112.html): Environment Inquiry
- [SRFI 113](http://srfi.schemers.org/srfi-113/srfi-113.html): Sets and Bags
- [SRFI 114](http://srfi.schemers.org/srfi-114/srfi-114.html): Comparators
- [SRFI 115](http://srfi.schemers.org/srfi-115/srfi-115.html): Scheme Regular Expressions
- [SRFI 116](http://srfi.schemers.org/srfi-116/srfi-116.html): Immutable List Library
- [SRFI 117](http://srfi.schemers.org/srfi-117/srfi-117.html): Queues based on lists
- [SRFI 118](http://srfi.schemers.org/srfi-118/srfi-118.html): Simple adjustable-size strings
- [SRFI 119](http://srfi.schemers.org/srfi-119/srfi-119.html): wisp: simpler indentation-sensitive scheme
- [SRFI 120](http://srfi.schemers.org/srfi-120/srfi-120.html): Timer APIs
- [SRFI 121](http://srfi.schemers.org/srfi-121/srfi-121.html): Generators
- [SRFI 123](http://srfi.schemers.org/srfi-123/srfi-123.html): Generic accessor and modifier operators
- [SRFI 124](http://srfi.schemers.org/srfi-124/srfi-124.html): Ephemerons
- [SRFI 126](http://srfi.schemers.org/srfi-126/srfi-126.html): R6RS-based hashtables
- [SRFI 127](http://srfi.schemers.org/srfi-127/srfi-127.html): Lazy Sequences
- [SRFI 128](http://srfi.schemers.org/srfi-128/srfi-128.html): Comparators (reduced)
- [SRFI 131](http://srfi.schemers.org/srfi-131/srfi-131.html): ERR5RS Record Syntax (reduced)
