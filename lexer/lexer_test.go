package lexer

import (
	"fmt"
	"strings"
	"testing"

	"github.com/charithe/monkeylang/token"
	"github.com/stretchr/testify/assert"
)

func TestIndividualTokens(t *testing.T) {
	source := strings.NewReader(`=+(){},;-!*/<>`)
	testCases := []struct {
		name            string
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{"assign", token.ASSIGN, "="},
		{"plus", token.PLUS, "+"},
		{"lparen", token.LPAREN, "("},
		{"rparen", token.RPAREN, ")"},
		{"lbrace", token.LBRACE, "{"},
		{"rbrace", token.RBRACE, "}"},
		{"comma", token.COMMA, ","},
		{"semicolon", token.SEMICOLON, ";"},
		{"minus", token.MINUS, "-"},
		{"bang", token.BANG, "!"},
		{"asterisk", token.ASTERISK, "*"},
		{"slash", token.SLASH, "/"},
		{"lt", token.LT, "<"},
		{"gt", token.GT, ">"},
		{"eof", token.EOF, ""},
	}

	l := New(source)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tok, err := l.NextToken()
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedType, tok.Type)
			assert.Equal(t, tc.expectedLiteral, tok.Literal)
		})
	}
}

func TestTokensFromSource(t *testing.T) {
	source := strings.NewReader(`let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
};

let result = add(five, ten);
if a < b || b > c && a == c || b != c {
	return true
} else {
	return false
}`)
	testCases := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.IDENT, "a"},
		{token.LT, "<"},
		{token.IDENT, "b"},
		{token.OR, "||"},
		{token.IDENT, "b"},
		{token.GT, ">"},
		{token.IDENT, "c"},
		{token.AND, "&&"},
		{token.IDENT, "a"},
		{token.EQ, "=="},
		{token.IDENT, "c"},
		{token.OR, "||"},
		{token.IDENT, "b"},
		{token.NEQ, "!="},
		{token.IDENT, "c"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(source)
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("token_%d", i), func(t *testing.T) {
			tok, err := l.NextToken()
			fmt.Printf("%d -> %s -> %s\n", i, tc.expectedLiteral, tok.Literal)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedType, tok.Type)
			assert.Equal(t, tc.expectedLiteral, tok.Literal)
		})
	}
}
