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
	IDENT       // Identifiers
	INT         // integral values
	ASSIGN      // =
	PLUS        // +
	MINUS       // -
	BANG        // !
	ASTERISK    // *
	SLASH       // /
	BITWISE_AND // &
	BITWISE_OR  // |
	EQ          // ==
	NEQ         // !=
	AND         // &&
	OR          // ||
	COMMA
	SEMICOLON
	LPAREN // (
	RPAREN // )
	LBRACE // {
	RBRACE // }
	LT     // <
	GT     // >
	FUNCTION
	LET
	TRUE
	FALSE
	IF
	ELSE
	RETURN
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) (TokenType, string) {
	if tt, ok := keywords[ident]; ok {
		return tt, ident
	}

	return IDENT, ident
}
