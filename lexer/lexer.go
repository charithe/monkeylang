package lexer

import (
	"bufio"
	"io"

	"github.com/charithe/monkeylang/token"
)

type Lexer struct {
	source *bufio.Reader
	eof    bool
	line   int
	column int
}

func New(input io.Reader) *Lexer {
	return &Lexer{
		source: bufio.NewReader(input),
		eof:    false,
	}
}

func (l *Lexer) NextToken() (*token.Token, error) {
	if l.eof {
		return nil, io.EOF
	}

	// have we reached the end?
	r, _, err := l.source.ReadRune()
	if err == io.EOF {
		l.eof = true
		return l.makeToken(token.EOF, ""), nil
	} else if err != nil {
		return nil, err
	}

	var tok *token.Token
	// increment column
	l.column = l.column + 1

	// determine token
	switch r {
	case '\n':
		tok = l.makeToken(token.NEWLINE, string(r))
		l.line = l.line + 1
		l.column = 1
	case '=':
		tok = l.makeToken(token.ASSIGN, string(r))
	case '+':
		tok = l.makeToken(token.PLUS, string(r))
	case ',':
		tok = l.makeToken(token.COMMA, string(r))
	case ';':
		tok = l.makeToken(token.SEMICOLON, string(r))
	case '(':
		tok = l.makeToken(token.LPAREN, string(r))
	case ')':
		tok = l.makeToken(token.RPAREN, string(r))
	case '{':
		tok = l.makeToken(token.LBRACE, string(r))
	case '}':
		tok = l.makeToken(token.RBRACE, string(r))
	default:
		tok = l.makeToken(token.ILLEGAL, string(r))
	}

	return tok, nil
}

func (l *Lexer) makeToken(tokenType token.TokenType, literal string) *token.Token {
	return &token.Token{
		Type:    tokenType,
		Literal: literal,
		Line:    l.line,
		Column:  l.column,
	}
}
