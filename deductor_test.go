package main

import (
	"regexp"
	"testing"
)

var simpleDelimParser Parser = Parser{*regexp.MustCompile(".|~ | & "), map[string]Operator{}}

func TestLexemesOnEmptyLine(t *testing.T) {
	got := simpleDelimParser.lexemes("")
	if len(got) > 0 {
		t.Errorf(`singleDelimParser.lexemes("") = %v; want []`, got)
	}
}

func TestLexemesOnSingleDelimiter(t *testing.T) {
	got := simpleDelimParser.lexemes(" & ")
	if len(got) != 1 || got[0] != "&" {
		t.Errorf(`singleDelimParser.lexemes(" & ") = "%v"; want [&]`, got)
	}
}

func TestLexemesOnSingleNondelimiter(t *testing.T) {
	got := simpleDelimParser.lexemes("a")
	if len(got) != 1 || got[0] != "a" {
		t.Errorf(`singleDelimParser.lexemes("a") = "%v"; want [a]`, got)
	}
}

func TestLexemesOnComplexExpression(t *testing.T) {
	got := simpleDelimParser.lexemes(".a & b. & ~ c")
	if len(got) != 8 || got[0] != "." || got[1] != "a" || got[2] != "&" || got[3] != "b" || got[4] != "." || got[5] != "&" || got[6] != "~" || got[7] != "c" {
		t.Errorf(`singleDelimParser.lexemes("a") = "%v"; want [. a & b . & ~ c]`, got)
	}
}
