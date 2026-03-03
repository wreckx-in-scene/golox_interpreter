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
