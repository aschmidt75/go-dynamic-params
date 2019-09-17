package main

import (
	"fmt"
)

type mode int
type tokenType int

const (
	modeNorm mode = iota
	modeBeginParam
	modeInParam
)

const (
	typeStaticPart tokenType = iota
	typeParamPart
)

// TokenizeError is
type TokenizeError struct {
	what  string
	pos   int
	token *Token
}

func (e *TokenizeError) Error() string {
	return fmt.Sprintf("%s at pos %d (%s)", e.what, e.pos, e.token.part)
}

// Token is a single Token split by the tokenizer
type Token struct {
	part            string
	tkType          tokenType
	withNestedParam bool
}

// Tokenizer is a
type Tokenizer struct {
	// TODO: bytes..
	in string
}

// NewTokenizer creates a new tokenizer
func NewTokenizer(inputString string) *Tokenizer {
	return &Tokenizer{in: inputString}
}

// Tokenize splits the given input string into tokens
func (t *Tokenizer) Tokenize() ([]*Token, error) {
	var err error
	res := make([]*Token, 0, 10)

	l := len(t.in)
	i := 0
	mode := modeNorm
	curToken := &Token{tkType: typeStaticPart}
	bracketCounter := 0

	for {
		skip1 := false

		if i >= l {
			if len(curToken.part) > 0 {
				res = append(res, curToken)
			}

			break
		}

		switch mode {
		case modeNorm:
			bracketCounter = 0
			if t.in[i] == '$' {
				// peek one ahead
				if i < l-1 {
					if t.in[i+1] == '{' {
						mode = modeBeginParam
						skip1 = true
					}
				}
			}

		case modeBeginParam:
			if t.in[i] == '{' {
				if len(curToken.part) > 0 {
					res = append(res, curToken)
				}
				curToken = &Token{tkType: typeParamPart}

				mode = modeInParam
				skip1 = true
			}

		case modeInParam:
			if t.in[i] == '{' {
				bracketCounter++

				curToken.withNestedParam = true
			}
			if t.in[i] == '}' {
				bracketCounter--
				if bracketCounter < 0 {
					if len(curToken.part) > 0 {
						res = append(res, curToken)
					} else {
						return res, &TokenizeError{what: "empty params not allowed", pos: i, token: curToken}
					}
					curToken = &Token{tkType: typeStaticPart}

					mode = modeNorm
					skip1 = true
				}
			}

		}

		if skip1 == false {
			curToken.part = fmt.Sprintf("%s%c", curToken.part, t.in[i])
		}

		//fmt.Printf("%d (%c) mode=%d bc=%d curTokenPart=%s\n", i, t.in[i], mode, bracketCounter, curToken.part)

		i++
	}

	if mode == modeInParam {
		return res, &TokenizeError{what: "invalid bracket balance", pos: i, token: curToken}
	}

	return res, err
}

