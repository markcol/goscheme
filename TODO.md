- Replace usage of Cons.Sexp() w/ Cons.Map and Cons.Len
- Replace evalValue(Value) w/ Eval(Cons)
- Use interface for Value.val instead of interface{}
- Move tokens.Expand() to be called before Parse(), like tokens.Expand().Parse() sort of
- Rename Sexp => Vector
- Add cons command
- Remove typecasting commands on Value
- Remove naked returns
- Add methods to Cons
- Pairs
  - [http://groups.csail.mit.edu/mac/ftpdir/scheme-7.4/doc-html/scheme_8.html#SEC73](http://groups.csail.mit.edu/mac/ftpdir/scheme-7.4/doc-html/scheme_8.html#SEC73)
  - [http://docs.racket-lang.org/guide/Pairs__Lists__and_Racket_Syntax.html](http://docs.racket-lang.org/guide/Pairs__Lists__and_Racket_Syntax.html)

- Add support for Line/pos in error messages
  - Store pos in token
  - Store token in value
  - Calculate line/pos lazily when outputting

- Add string builtins:
-   [http://groups.csail.mit.edu/mac/ftpdir/scheme-7.4/doc-html/scheme_7.html#SEC61](http://groups.csail.mit.edu/mac/ftpdir/scheme-7.4/doc-html/scheme_7.html#SEC61)
-   string-concatenate...
- Add list builtins:
-   [http://groups.csail.mit.edu/mac/ftpdir/scheme-7.4/doc-html/scheme_8.html#SEC72](http://groups.csail.mit.edu/mac/ftpdir/scheme-7.4/doc-html/scheme_8.html#SEC72)
-   cons, car, cdr...
- Macros
  - [http://www.willdonnelly.net/blog/scheme-syntax-rules/](http://www.willdonnelly.net/blog/scheme-syntax-rules/)
  - [http://docs.racket-lang.org/guide/pattern-macros.html](http://docs.racket-lang.org/guide/pattern-macros.html)

- Implement tapping into packages
  - Use pkgreflect to create maps
    - [https://github.com/ungerik/pkgreflect](https://github.com/ungerik/pkgreflect)

  - (fmt/Printf "Hello, world!")
  - (.Body (first (http/Get "[http://example.com/](http://example.com/)")))
  - (. (. (. http (Get "[http://example.com/](http://example.com/)")) (first)) (Body))
  - (.. (http (Get "[http://example.com/](http://example.com/)")) (first) (Body))
  - (-> (http (Get "[http://example.com/](http://example.com/)")) (first) (Body))
  - (doto (http/Get "[http://example.com/](http://example.com/)") (first) (Body))
  - Compare [http://clojure.org/java_interop](http://clojure.org/java_interop)

- Add short forms
  - (define (inc x) (+ x 1)) => (define inc (lambda (x) (+ x 1)))
  - (let ((a 1) (b 2)) (+ a b)) => ((lambda (a b) (+ a b)) 1 2)
  - (list a b) => (cons a (cons b ()))
  - (a . b) => (cons a b)

- Support for modules
  - [http://www.htus.org/Book/Staging/how-to-use-modules/](http://www.htus.org/Book/Staging/how-to-use-modules/)

- More support for Scheme
  - [http://people.csail.mit.edu/jaffer/r5rs_toc.html](http://people.csail.mit.edu/jaffer/r5rs_toc.html)
