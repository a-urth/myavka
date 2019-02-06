package lexer

import (
	"github.com/a-urth/myavka/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := Lexer{input: input}
	l.readChar()
	return &l
}

func (l *Lexer) NextToken() token.Token {
	var (
		tokenType  token.TokenType
		tokenValue string
	)

	l.skipWhitespaces()

	tokenValue = string(l.ch)

	switch l.ch {
	case '=':
		tokenType = token.ASSIGN
		if l.peekChar() == '=' {
			tokenType = token.EQ
			tokenValue = "=="
			l.readChar()
		}
	case ';':
		tokenType = token.SEMICOLON
	case '(':
		tokenType = token.LPAREN
	case ')':
		tokenType = token.RPAREN
	case '{':
		tokenType = token.LBRACE
	case '}':
		tokenType = token.RBRACE
	case ',':
		tokenType = token.COMMA
	case '+':
		tokenType = token.PLUS
	case '!':
		tokenType = token.BANG
		if l.peekChar() == '=' {
			tokenType = token.NOT_EQ
			tokenValue = "!="
			l.readChar()
		}
	case '-':
		tokenType = token.MINUS
	case '*':
		tokenType = token.ASTERISK
	case '<':
		tokenType = token.LT
	case '>':
		tokenType = token.GT
	case '/':
		tokenType = token.SLASH
	case 0:
		tokenType = token.EOF
		tokenValue = ""
	default:
		tokenType = token.ILLEGAL

		if isLetter(l.ch) {
			tokenValue = l.readTerm(isLetter)
			tokenType = token.LookupIdent(tokenValue)
			return token.NewToken(tokenType, tokenValue)
		} else if isDigit(l.ch) {
			tokenValue = l.readTerm(isDigit)
			tokenType = token.INT
			return token.NewToken(tokenType, tokenValue)
		}
	}

	l.readChar()

	return token.NewToken(tokenType, tokenValue)
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) skipWhitespaces() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readTerm(pred func(byte) bool) string {
	position := l.position
	for pred(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isLetter(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_'
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}
