package main

import "fmt"

func main() {
	text := `
	name -> rojesh
	age -> 
	friends -> robus,rajan,raman
	`
	lexer := Lexer{text, 0}
	curToken := lexer.getNextToken()
	parser := Parser{lexer, curToken}
	interpreter := NewInterpreter(parser)
	interpreter.interpret()
	fmt.Println(interpreter.data)
}
