package lexer

import (
	"strings"
	"testing"

	"github.com/charithe/monkeylang/token"
	"github.com/stretchr/testify/assert"
)

func TestNextToken(t *testing.T) {
	source := strings.NewReader(`=+(){},;`)
	testCases := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(source)
	for _, tc := range testCases {
		tok, err := l.NextToken()
		assert.NoError(t, err)
		assert.Equal(t, tc.expectedType, tok.Type)
		assert.Equal(t, tc.expectedLiteral, tok.Literal)
	}
}
