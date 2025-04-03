package main

import (
	"bufio"
	"carametal/CaraScript/evaluator"
	"carametal/CaraScript/lexer"
	"carametal/CaraScript/parser"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to CaraScript.")
	fmt.Println("Type \"exit\" to exit interpreter.")

	for {
		fmt.Print("> ")
		scanned := scanner.Scan()
		if !scanned {
			break
		}

		line := scanner.Text()
		if line == "exit" {
			break
		}

		l := lexer.New(line)

		p := parser.New(l)
		program := p.ParseProgram()

		if program != nil {
			evaluated := evaluator.Eval(program)
			if evaluated != nil {
				fmt.Println(evaluated.String())
			}
		}
	}
	fmt.Println("Exit CaraScript. Good bye!!")
}
