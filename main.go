package main

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

type TokenType int

const (
	Identifier TokenType = iota
	Number
	Operator
	Keyword
)

type Token struct {
	Type    TokenType
	Literal string
}

func main() {
	input := "let x = 42 + y;"

	lexer := createLexer(input)

	for {
		token := lexer.Scan()
		if token == scanner.EOF {
			break
		}

		tokenType := determineTokenType(lexer.TokenText())
		fmt.Printf("Token Type: %s, Literal: %s\n", tokenType, lexer.TokenText())
	}
}

func createLexer(input string) *scanner.Scanner {
	var lexer scanner.Scanner
	lexer.Init(strings.NewReader(input))

	return &lexer
}

func determineTokenType(tokenLiteral string) TokenType {
	// Basic logic to determine token type, you can extend this based on your language's syntax
	if isKeyword(tokenLiteral) {
		return Keyword
	} else if isOperator(tokenLiteral) {
		return Operator
	} else if isNumber(tokenLiteral) {
		return Number
	}

	return Identifier
}

func isKeyword(token string) bool {
	// Example: check if token is a keyword in your language
	keywords := []string{"let", "if", "else", "for", "while", "func", "return", "true", "false", "nil", "and", "or", "const", "var", "struct", "int", "float"}
	for _, kw := range keywords {
		if token == kw {
			return true
		}
	}
	return false
}

func isOperator(token string) bool {
	// Example: check if token is an operator in your language
	operators := []string{"+", "-", "*", "/"}
	for _, op := range operators {
		if token == op {
			return true
		}
	}
	return false
}

func isNumber(token string) bool {
	// Example: check if token is a number in your language
	// You might want to use more complex logic based on your language's numeric literals
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}
