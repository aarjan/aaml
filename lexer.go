package main

import (
	"regexp"
)

type Lexer struct {
	text       string
	currentPos int
}

func (l *Lexer) currentChar() []byte {
	if l.currentPos >= len(l.text) {
		return nil
	}
	return []byte{l.text[l.currentPos]}
}

func (l *Lexer) aheadChar() string {
	l.currentPos++
	if l.currentPos >= len(l.text) {
		return ""
	}
	return string(l.text[l.currentPos])
}

func (l *Lexer) advance() {
	l.currentPos++
}

// Whitespace may be a 'space' or a 'tab'
func (l *Lexer) skipWhitespace() {
	for l.currentChar() != nil && (string(l.currentChar()) == " " || l.currentChar()[0] == 9) {
		l.advance()
	}
}

func (l *Lexer) newLine() Token {
	for l.currentChar() != nil && string(l.currentChar()) == "\n" {
		l.advance()
	}
	return NewlineToken()
}

// Return appended text if it is a alphanumerical character
func (l *Lexer) variable() string {

	result := []byte{}
	regex := regexp.MustCompile("[0-9a-zA-z]")

	for l.currentChar() != nil && regex.Match(l.currentChar()) {
		result = append(result, l.currentChar()[0])
		l.advance()
	}
	return string(result)
}

func (l *Lexer) getNextToken() Token {
	regex := regexp.MustCompile("[0-9a-zA-z]")
	for {
		char := string(l.currentChar())
		switch {
		case char == "":
			return EOFToken()
		case char == " " || char == "\t":
			l.skipWhitespace()
			continue
		case regex.MatchString(char):
			return StringToken(l.variable())
		case char == "-" && l.aheadChar() == ">":
			l.advance()
			l.advance()
			return ArrowToken()
		case char == ",":
			l.advance()
			return CommaToken()
		case char == "\n":
			return l.newLine()
		default:
			Error()
		}
	}
}

// Error ...
func Error() {
	panic("Syntax Error")
}
