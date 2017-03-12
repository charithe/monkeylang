package lexer

import (
	"bufio"
	"bytes"
	"io"
	"unicode"

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

	var tok *token.Token
	var err error

	r, err := l.skipWhitespaceAndReadNext()
	// have we reached the end?
	if err == io.EOF {
		l.eof = true
		return l.makeToken(token.EOF, ""), nil
	} else if err != nil {
		return nil, err
	}

	// determine token
	switch r {
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
		if l.isIdentifierStartChar(r) {
			tok, err = l.readIdentifer(r)
		} else if l.isDigit(r) {
			tok, err = l.readNumber(r)
		} else {
			tok = l.makeToken(token.ILLEGAL, string(r))
		}
	}

	return tok, err
}

func (l *Lexer) makeToken(tokenType token.TokenType, literal string) *token.Token {
	return &token.Token{
		Type:    tokenType,
		Literal: literal,
		Line:    l.line,
		Column:  l.column,
	}
}

func (l *Lexer) skipWhitespaceAndReadNext() (rune, error) {
	for {
		r, _, err := l.source.ReadRune()
		if err != nil {
			return 0, err
		}

		if r == '\n' {
			l.line = l.line + 1
			l.column = 1
		} else {
			l.column = l.column + 1
		}

		if unicode.IsSpace(r) {
			continue
		}

		return r, nil
	}
}

func (l *Lexer) isIdentifierStartChar(r rune) bool {
	return unicode.IsLetter(r)
}

func (l *Lexer) isIdentifierChar(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'
}

func (l *Lexer) readIdentifer(r rune) (*token.Token, error) {
	var identifier bytes.Buffer
	identifier.WriteRune(r)
	for {
		nr, _, err := l.source.ReadRune()
		if err != nil {
			return nil, err
		}

		if l.isIdentifierChar(nr) {
			identifier.WriteRune(nr)
			l.column = l.column + 1
		} else {
			if err := l.source.UnreadRune(); err != nil {
				return nil, err
			}
			return l.makeToken(token.LookupIdent(identifier.String())), nil
		}
	}
}

func (l *Lexer) isDigit(r rune) bool {
	return unicode.IsDigit(r)
}

func (l *Lexer) readNumber(r rune) (*token.Token, error) {
	//TODO floats
	var num bytes.Buffer
	num.WriteRune(r)
	for {
		nr, _, err := l.source.ReadRune()
		if err != nil {
			return nil, err
		}

		if l.isDigit(nr) {
			num.WriteRune(nr)
			l.column = l.column + 1
		} else {
			if err := l.source.UnreadRune(); err != nil {
				return nil, err
			}
			return l.makeToken(token.INT, num.String()), nil
		}
	}
}
