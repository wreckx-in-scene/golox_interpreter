package main

import (
	"fmt"
	"strconv"
)

type Lexer struct {
	source  string
	tokens  []Token
	current int
	start   int
	line    int
}

// keyword map
var keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"fun":    FUN,
	"for":    FOR,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

// lexer functions
func (l *Lexer) isAtEnd() bool {
	return l.current >= len(l.source)
}

func (l *Lexer) advance() byte {
	temp := l.current
	l.current++
	return l.source[temp]
}

func (l *Lexer) addToken(tokType TokenType) {
	//creating a new token
	token := Token{
		Type:    tokType,
		Lexeme:  l.source[l.start:l.current],
		Literal: nil,
		Line:    l.line,
	}

	l.tokens = append(l.tokens, token)
}

func (l *Lexer) peek() byte {
	if l.isAtEnd() {
		return 0
	}
	return l.source[l.current]
}

func (l *Lexer) match(ch byte) bool {
	if l.isAtEnd() {
		return false
	}
	if ch == l.peek() {
		l.current++
		return true
	}
	return false

}

func (l *Lexer) scanToken() {
	curr := l.advance()
	switch curr {
	case '(':
		l.addToken(LEFT_PAREN)
	case ')':
		l.addToken(RIGHT_PAREN)
	case '{':
		l.addToken(LEFT_BRACE)
	case '}':
		l.addToken(RIGHT_BRACE)
	case ',':
		l.addToken(COMMA)
	case '.':
		l.addToken(DOT)
	case '-':
		l.addToken(MINUS)
	case '+':
		l.addToken(PLUS)
	case ';':
		l.addToken(SEMICOLON)
	case '/':
		if l.match('/') {
			for !l.isAtEnd() && l.peek() != '\n' {
				l.current++
			}
		} else {
			l.addToken(SLASH)
		}
	case '*':
		l.addToken(STAR)
	case '!':
		if l.match('=') {
			l.addToken(BANG_EQUAL)
		} else {
			l.addToken(BANG)
		}
	case '=':
		if l.match('=') {
			l.addToken(EQUAL_EQUAL)
		} else {
			l.addToken(EQUAL)
		}
	case '<':
		if l.match('=') {
			l.addToken(LESS_EQUAL)
		} else {
			l.addToken(LESS)
		}
	case '>':
		if l.match('=') {
			l.addToken(GREATER_EQUAL)
		} else {
			l.addToken(GREATER)
		}
	case ' ', '\r', '\t':

	case '\n':
		l.line++
	case '"':
		l.string()
	default:
		if l.isDigit(curr) {
			l.number()
		} else if l.isAlpha(curr) {
			l.identifier()
		} else {
			fmt.Println("Unexpected error on line ", l.line)
		}
	}
}

// function to use scan token
func (l *Lexer) ScanToken() []Token {
	for !l.isAtEnd() {
		l.start = l.current
		l.scanToken()
	}

	l.tokens = append(l.tokens, Token{Type: EOF, Lexeme: "", Literal: nil, Line: l.line})
	return l.tokens
}

// function to handle strings
func (l *Lexer) string() {
	//looping until closing " or end
	for l.peek() != '"' && !l.isAtEnd() {
		l.advance()
	}

	//if we reach end without closing "
	if l.isAtEnd() {
		fmt.Println("Error : unterminated string on line.", l.line)
		return
	}

	//consume closing "
	l.advance()

	//getting the required string value
	str := l.source[l.start+1 : l.current-1]

	//adding token
	l.tokens = append(l.tokens, Token{Type: STRING, Lexeme: l.source[l.start:l.current], Literal: str, Line: l.line})
}

// function to see if the char is a digit
func (l *Lexer) isDigit(ch byte) bool {
	if ch >= '0' && ch <= '9' {
		return true
	} else {
		return false
	}
}

// function to see if character is alphabet or underscore
func (l *Lexer) isAlpha(ch byte) bool {
	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_' {
		return true
	} else {
		return false
	}
}

// function to check if ch is alpha or number
func (l *Lexer) isAlphaNumeric(ch byte) bool {
	if l.isAlpha(ch) || l.isDigit(ch) {
		return true
	} else {
		return false
	}
}

// peeknext function
func (l *Lexer) peekNext() byte {
	if l.current+1 >= len(l.source) {
		return 0
	} else {
		return l.source[l.current+1]
	}
}

// function to handle the numbers
func (l *Lexer) number() {
	//looping till the next char is digit
	for !l.isAtEnd() && l.isDigit(l.peek()) {
		l.advance()
	}

	// if you find .
	if l.peek() == '.' && l.isDigit(l.peekNext()) {
		l.advance() //consume
		for !l.isAtEnd() && l.isDigit(l.peek()) {
			l.advance()
		}
	}

	//converting into the number
	str, _ := strconv.ParseFloat(l.source[l.start:l.current], 64)
	//adding the new token
	l.tokens = append(l.tokens, Token{Type: NUMBER, Lexeme: l.source[l.start:l.current], Literal: str, Line: l.line})
}

// function to handle identifiers
func (l *Lexer) identifier() {
	for !l.isAtEnd() && l.isAlphaNumeric(l.peek()) {
		l.advance()
	}

	//getting the word
	word := l.source[l.start:l.current]

	tokType, exists := keywords[word]
	if !exists {
		tokType = IDENTIFIER
	}

	l.addToken(tokType)
}
