package main

type Lexer struct {
	source  string
	tokens  []Token
	current int
	start   int
	line    int
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
