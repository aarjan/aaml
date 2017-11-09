package main

const ARROW = "ARROW"
const COMMA = "COMMA"
const STRING = "STRING"
const NEWLINE = "NEWLINE"
const EOF = "EOF"

type Token struct {
	_type string
	value string
}

func ArrowToken() Token {
	return Token{_type: ARROW}
}

func CommaToken() Token {
	return Token{_type: COMMA}
}
func StringToken(val string) Token {
	return Token{_type: STRING, value: val}
}

func NewlineToken() Token {
	return Token{_type: NEWLINE}
}

func EOFToken() Token {
	return Token{_type: EOF}
}
