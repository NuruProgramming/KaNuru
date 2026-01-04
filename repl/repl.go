// Package repl provides a Read-Eval-Print Loop implementation for interactive
// command execution and evaluation.
package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/NuruProgramming/KaNuru/evaluator"
	"github.com/NuruProgramming/KaNuru/lexer"
	"github.com/NuruProgramming/KaNuru/object"
	"github.com/NuruProgramming/KaNuru/parser"
)

const prompt = ">>> "

func Read(contents string) {
	env := object.NewEnvironment()

	l := lexer.New(contents)
	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		fmt.Println("Kuna Errors Zifuatazo:")

		for _, msg := range p.Errors() {
			fmt.Println("\t" + msg)
		}

	}
	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		if evaluated.Type() != object.NULL_OBJ {
			fmt.Println(evaluated.Inspect())
		}
	}
}

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if strings.TrimSpace(line) == "exit()" || strings.TrimSpace(line) == "toka()" {
			fmt.Println("Karibu Tena")
			os.Exit(0)
		}
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			if evaluated.Type() != object.NULL_OBJ {
				io.WriteString(out, evaluated.Inspect())
				io.WriteString(out, "\n")
			}
		}
	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Kuna Errors Zifuatazo:\n")

	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
