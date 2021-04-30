package main

import (
	"strconv"
)

type TOKENKind int

const (
	TK_RESERVED TOKENKind = iota + 1 // 記号
	TK_NUM
	TK_EOF
)

type Token struct {
	Kind TOKENKind
	Next *Token
	Val  int
	Str  string
}

var CurToken *Token

func isSpace(s byte) bool {
	switch s {
	case ' ', '\t', '\r', '\n':
		return true
	}
	return false
}

func isDigit(s byte) bool {
	if '0' <= s && s <= '9' {
		return true
	}
	return false
}

func strtol(a string, i *int) string {
	var result []byte
	for *i < len(a) {
		s := a[*i]
		if '0' <= s && s <= '9' {
			result = append(result, s)
			*i++
		} else {
			break
		}
	}
	return string(result)
}

func NewToken(kind TOKENKind, curToken *Token, str string) *Token {
	token := &Token{
		Kind: kind,
		Str:  str,
	}

	curToken.Next = token
	return token
}

func Tokenize(input string) *Token {
	head := &Token{}
	cur := head

	var i int

	for i < len(input) {
		token := input[i]

		if isSpace(token) {
			i++
			continue
		}

		if token == '+' || token == '-' {
			cur = NewToken(TK_RESERVED, cur, string(token))
			i++
			continue
		}

		if isDigit(token) {
			num := strtol(input, &i)
			cur = NewToken(TK_NUM, cur, num)
			cur.Val, _ = strconv.Atoi(num)
			continue
		}

		printfErr("予期しない文字です： '%s'", string(input[i]))
	}

	NewToken(TK_EOF, cur, "")

	return head.Next
}
