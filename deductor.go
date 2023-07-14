package main

import (
	"regexp"
	"strings"
)

type Token struct {
	type_ TokenType
	value string
}

type TokenType int

const (
	therefore TokenType = iota
	leftParen
	rightParen
	negation
	conjunction
	disjunction
	conditional
	biconditional
	variable
)

type Parser struct {
	matchDelimiters  regexp.Regexp
	OperatorByLexeme map[string]Operator
}

func (t *Parser) tokens(lexemes []string) []

func (t *Parser) lexemes(line string) []string {
	delimiters := t.matchDelimiters.FindAllString(line, -1)
	nondelimiters := t.matchDelimiters.Split(line, -1)
	result := make([]string, 0, len(delimiters)+len(nondelimiters))
	result = appendTrimSpaceIfNonempty(result, nondelimiters[0])
	for i := 0; i < len(delimiters); i++ {
		result = appendTrimSpaceIfNonempty(result, delimiters[i])
		result = appendTrimSpaceIfNonempty(result, nondelimiters[i+1])
	}
	return result
}

func appendTrimSpaceIfNonempty(slice []string, item string) []string {
	if trim := strings.TrimSpace(item); len(trim) > 0 {
		return append(slice, trim)
	}
	return slice
}

func main() {}
