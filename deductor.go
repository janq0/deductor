package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// TODO: clean this func up
	reader := bufio.NewReader(os.Stdin)
	reading := true
	var vars, conclusion []string
	var postfixPremises [][]string
	for reading {
		line, _ := reader.ReadString('\n')
		toks := tokens(line)
		postf := postfix(toks)
		if toks[0] == "therefore" {
			reading = false
			postf = postf[1:]
			conclusion = postf
		} else {
			postfixPremises = append(postfixPremises, postf)
			vars = varsUpdate(vars, variables(toks))
		}
	}
	conclusionTable := line_table(conclusion, vars)
	var premisesTables [][]bool
	for _, pre := range postfixPremises {
		premisesTables = append(premisesTables, line_table(pre, vars))
	}
	fmt.Println(isValid(premisesTables, conclusionTable))
}

func varsUpdate(old []string, new []string) []string {
	for _, v := range new {
		if !contains(v, old) {
			old = append(old, v)
		}
	}
	return old
}

func contains(value string, values []string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

func isValid(premiseTables [][]bool, conclusion []bool) bool {
	for row, concl := range conclusion {
		if truePremises(premiseTables, row) && !concl {
			return false
		}
	}
	return true
}

func truePremises(premiseTables [][]bool, row int) bool {
	for _, tab := range premiseTables {
		if !tab[row] {
			return false
		}
	}
	return true
}

var operatorTable = map[string][4]bool{
	"not":     {true, false, true, false},
	"and":     {false, false, false, true},
	"or":      {false, true, true, true},
	"if":      {true, true, false, true},
	"only if": {true, false, true, true},
}

func line_table(postfix []string, variables []string) []bool {
	loaded := make([][]bool, 0)
	for _, tok := range postfix {
		if !isOperator(tok) {
			table := variable_table(tok, variables)
			loaded = append(loaded, table)
			continue
		}
		optable := operatorTable[tok]
		var left, right []bool
		if tok == "not" {
			left = loaded[len(loaded)-1]
			right = loaded[len(loaded)-1]
			loaded = loaded[:len(loaded)-1]
		} else {
			right = loaded[len(loaded)-1]
			left = loaded[len(loaded)-2]
			loaded = loaded[:len(loaded)-2]
		}
		loaded = append(loaded, combine_tables(optable, left, right))
	}
	return loaded[0]
}

func combine_tables(operator [4]bool, left []bool, right []bool) []bool {
	result := make([]bool, len(left))
	for i := range left {
		oprow := intOfBool(left[i]) + 2*intOfBool(right[i])
		result[i] = operator[oprow]
	}
	return result
}

func intOfBool(b bool) int {
	if b {
		return 1
	}
	return 0
}

func variable_table(variable string, variables []string) []bool {
	length := 1 << len(variables)
	table := make([]bool, length)
	consec := 0
	for i, v := range variables {
		if v == variable {
			consec = i + 1
			break
		}
	}
	for i := range table {
		table[i] = (i/consec)%2 == 1
	}
	return table
}

func variables(tokens []string) []string {
	result := make([]string, 0, len(tokens))
	for _, tok := range tokens {
		if !isOperator(tok) {
			result = append(result, tok)
		}
	}
	return result
}

var operators = [...]string{
	"(", ")", "if", "only if", "and", "or", "not",
} // Order by precedence

func isOperator(operator string) bool {
	for _, op := range operators {
		if op == operator {
			return true
		}
	}
	return false
}

func postfix(infix []string) []string {
	// Shunting yard algorithm
	result := make([]string, 0, len(infix))
	opstack := make([]string, 0, len(infix))
	for _, tok := range infix {
		prec := precedence(tok)
		if prec == -1 { // Is a variable
			result = append(result, tok)
			continue
		}
		if tok == "(" {
			opstack = append(opstack, tok)
			continue
		}
		if tok == ")" {
			for opstack[len(opstack)-1] != "(" {
				result = append(result, opstack[len(opstack)-1])
				opstack = opstack[:len(opstack)-1]
			}
			opstack = opstack[:len(opstack)-1]
			continue
		}
		for len(opstack) > 0 && prec < precedence(opstack[len(opstack)-1]) {
			result = append(result, opstack[len(opstack)-1])
			opstack = opstack[:len(opstack)-1]
		}
		opstack = append(opstack, tok)
	}
	for len(opstack) > 0 {
		result = append(result, opstack[len(opstack)-1])
		opstack = opstack[:len(opstack)-1]
	}
	return result
}

func precedence(operator string) int {
	for i, op := range operators {
		if op == operator {
			return i
		}
	}
	return -1
}

var matchDelimiters = regexp.MustCompile(` if | only if | and | or |not |\(|\)|therefore `)

func tokens(line string) []string {
	line = strings.ToLower(line)
	delimiters := matchDelimiters.FindAllString(line, -1)
	nondelimiters := matchDelimiters.Split(line, -1)
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
