package token

type TokenType uint64

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

const (
	ILLEGAL TokenType = 1 << iota
	EOF
	NEWLINE
	IDENT  // Identifiers
	INT    // integral values
	ASSIGN // = opeerator
	PLUS   // + operator
	COMMA
	SEMICOLON
	LPAREN // (
	RPAREN // )
	LBRACE // {
	RBRACE // }
	FUNCTION
	LET
)
