package main

import "fmt"

func main() {
	text := `
	name -> aarjan
	age -> 21
	friends -> robus,rajan,raman
	`
	lexer := Lexer{text, 0}
	for {
		curToken := lexer.getNextToken()
		if curToken._type == EOF {
			break
		}
		parser := Parser{lexer, curToken}
		list := parser.statementList()

		fmt.Printf("%+v\n", list.nodes)
	}

}
