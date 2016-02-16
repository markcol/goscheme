package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/markcol/goscheme"
)

func main() {
	var file = flag.String("file", "", "File to interpret")
	var repl = flag.Bool("i", false, "Interactive mode")
	flag.Parse()

	if *repl {
		Repl()
	} else if *file != "" {
		if output, err := ioutil.ReadFile(*file); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		} else {
			if _, err := goscheme.EvalString(string(output)); err != nil {
				fmt.Printf("ERROR: %v\n", err)
			}
		}
	} else {
		if output, err := ioutil.ReadAll(bufio.NewReader(os.Stdin)); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		} else {
			goscheme.EvalString(string(output))
		}
	}
}

// Repl is the read-eval-print loop for goschemee.
func Repl() {
	fmt.Printf("Welcome to the GoScheme REPL\n")
	reader := bufio.NewReader(os.Stdin)
	expr := ""
	for {
		if expr == "" {
			fmt.Printf("\n> ")
		}
		line, _ := reader.ReadString('\n')
		expr = fmt.Sprintf("%v%v", expr, line)
		openCount := strings.Count(expr, "(")
		closeCount := strings.Count(expr, ")")
		if openCount < closeCount {
			fmt.Printf("ERROR: Malformed expression: %v", line)
			expr = ""
		} else if openCount == closeCount {
			if strings.TrimSpace(expr) != "" {
				if response, err := goscheme.EvalString(expr); err != nil {
					fmt.Printf("ERROR: %v\n", err)
				} else {
					if response == goscheme.Nil {
						fmt.Println(";Unspecified return value")
					} else {
						fmt.Printf(";Value: %v\n", response.Inspect())
					}
				}
			}
			expr = ""
		}
	}
}
