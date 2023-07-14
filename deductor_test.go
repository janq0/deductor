package main

import (
	"regexp"
	"testing"
)

var simpleDelimParser Parser = Parser{*regexp.MustCompile(".|~ | & "), map[string]Operator{}}

func TestTokensOnEmptyLine(t *testing.T) {
	got := simpleDelimParser.tokens("")
	if len(got) > 0 {
		t.Errorf(`singleDelimParser.tokens("") = %v; want []`, got)
	}
}

func TestTokensOnSingleDelimiter(t *testing.T) {
	got := simpleDelimParser.tokens(" & ")
	if len(got) != 1 || got[0] != "&" {
		t.Errorf(`singleDelimParser.tokens(" & ") = "%v"; want [&]`, got)
	}
}

func TestTokensOnSingleNondelimiter(t *testing.T) {
	got := simpleDelimParser.tokens("a")
	if len(got) != 1 || got[0] != "a" {
		t.Errorf(`singleDelimParser.tokens("a") = "%v"; want [a]`, got)
	}
}

func TestTokensOnComplexExpression(t *testing.T) {
	got := simpleDelimParser.tokens(".a & b. & ~ c")
	if len(got) != 8 || got[0] != "." || got[1] != "a" || got[2] != "&" || got[3] != "b" || got[4] != "." || got[5] != "&" || got[6] != "~" || got[7] != "c" {
		t.Errorf(`singleDelimParser.tokens("a") = "%v"; want [. a & b . & ~ c]`, got)
	}
}
