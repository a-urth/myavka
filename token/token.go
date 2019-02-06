package token

const (
	ILLEGAL    = `ILLEGAL`
	EOF        = `EOF`
	FUNCTION   = `FUNCTION`
	LET        = `LET`
	IDENTIFIER = `IDENT`
	INT        = `INT`
	TRUE       = `TRUE`
	FALSE      = `FALSE`
	IF         = `IF`
	ELSE       = `ELSE`
	RETURN     = `RETURN`

	// operators
	ASSIGN   = `=`
	EQ       = `==`
	NOT_EQ   = `!=`
	PLUS     = `+`
	MINUS    = `-`
	BANG     = `!`
	ASTERISK = `*`
	SLASH    = `/`
	LT       = `<`
	GT       = `>`

	// mics
	COMMA     = `,`
	SEMICOLON = `;`
	LPAREN    = `(`
	RPAREN    = `)`
	LBRACE    = `{`
	RBRACE    = `}`
)

var keywords = map[string]TokenType{
	`fn`:     FUNCTION,
	`let`:    LET,
	`true`:   TRUE,
	`false`:  FALSE,
	`if`:     IF,
	`else`:   ELSE,
	`return`: RETURN,
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tokenType TokenType, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
	}
}

func LookupIdent(k string) TokenType {
	if tok, ok := keywords[k]; ok {
		return tok
	}
	return IDENTIFIER
}
