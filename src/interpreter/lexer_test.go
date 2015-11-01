package interpreter

/**
 * TODO
 * - End numbers and names when ) is read, appending END_NAME (or END_NUMBER) and END_SEXP
 * - Distinguish between ' and " so that one doesn't end the other
 */

import (
	"regexp"
	"testing"
)

func TestTransitionsFromOpen(t *testing.T) {
	tests := [...]string{"`", "(", "\"", ";", "3", "H", ")"}
	for i, transition := range TransitionsFromOpen {
		matched, err := regexp.MatchString(transition.ReadMatch, tests[i])
		if err != nil {
			t.Error(err.Error())
		} else if !matched {
			t.Errorf("%s did not match %s\n", transition.ReadMatch, tests[i])
		}
	}
}

func TestTransitionsFromList(t *testing.T) {
	tests := [...]string{"(", "9"}
	for i, transition := range TransitionsFromList {
		matched, err := regexp.MatchString(transition.ReadMatch, tests[i])
		if err != nil {
			t.Error(err.Error())
		} else if !matched {
			t.Errorf("%s did not match %s\n", transition.ReadMatch, tests[i])
		}
	}
}

func TestTransitionsFromString(t *testing.T) {
	tests := [...]string{"'", " "}
	for i, transition := range TransitionsFromString {
		matched, err := regexp.MatchString(transition.ReadMatch, tests[i])
		if err != nil {
			t.Error(err.Error())
		} else if !matched {
			t.Errorf("%s did not match %s\n", transition.ReadMatch, tests[i])
		}
	}
}

func TestTransitionsFromComment(t *testing.T) {
	tests := [...]string{"\n", ";"}
	for i, transition := range TransitionsFromComment {
		matched, err := regexp.MatchString(transition.ReadMatch, tests[i])
		if err != nil {
			t.Error(err.Error())
		} else if !matched {
			t.Errorf("%s did not match %s\n", transition.ReadMatch, tests[i])
		}
	}
}

func TestTransitionsFromNumber(t *testing.T) {
	tests := [...]string{"\n", "3"}
	for i, transition := range TransitionsFromNumber {
		matched, err := regexp.MatchString(transition.ReadMatch, tests[i])
		if err != nil {
			t.Error(err.Error())
		} else if !matched {
			t.Errorf("%s did not match %s\n", transition.ReadMatch, tests[i])
		}
	}
}

func TestTransitionsFromName(t *testing.T) {
	tests := [...]string{"\t", "_"}
	for i, transition := range TransitionsFromName {
		matched, err := regexp.MatchString(transition.ReadMatch, tests[i])
		if err != nil {
			t.Error(err.Error())
		} else if !matched {
			t.Errorf("%s did not match %s\n", transition.ReadMatch, tests[i])
		}
	}
}

func TestTransition(t *testing.T) {
	tests := [...]struct {
		From   State
		To     State
		Input  string
		Action RecursiveAction
		Tokens []Token
	}{
		{OPEN, STRING, "\"", DoNothing, []Token{START_STRING}},
		{OPEN, NUMBER, "0", DoNothing, []Token{START_NUMBER, "0"}},
		{OPEN, OPEN, ")", Return, []Token{END_SEXP}},
		{LIST, OPEN, "(", Recurse, []Token{}},
		{STRING, OPEN, "'", Return, []Token{END_STRING}},
		{COMMENT, OPEN, "\n", Return, []Token{END_COMMENT}},
		{NUMBER, OPEN, "\t", Return, []Token{END_NUMBER}},
		{NAME, NAME, "f", DoNothing, []Token{Token("f")}},
	}
	for _, test := range tests {
		err, newState, action, tokens := Transition(test.From, test.Input)
		if err != nil {
			t.Error(err.Error())
		}
		if newState != test.To {
			t.Errorf("Expected state %d got %d\n", test.To, newState)
		}
		if action != test.Action {
			t.Errorf("Expected action %d got %d\n", test.Action, action)
		}
		if len(tokens) != len(test.Tokens) {
			t.Log(test)
			t.Errorf("Expected %d tokens got %d\n", len(test.Tokens), len(tokens))
		}
		for i, _ := range tokens {
			if tokens[i] != test.Tokens[i] {
				t.Errorf("Expected token %s got %s\n", test.Tokens[i], tokens[i])
			}
		}
	}
}

func TestLex(t *testing.T) {
	tests := [...]struct {
		Program string
		Lexed   []Token
	}{
		{"3", []Token{START_NUMBER, "3"}},             // Not a valid program, but useful lex test
		{"2.2", []Token{START_NUMBER, "2", ".", "2"}}, // Also not valid
		{"'test'", []Token{START_STRING, "t", "e", "s", "t", END_STRING}},
		{"\"t\"", []Token{START_STRING, "t", END_STRING}},
		{"(hi)", []Token{START_SEXP, START_NAME, "h", "i", END_SEXP}},
		{"func", []Token{START_NAME, "f", "u", "n", "c", END_NAME}},
		{";comment\n", []Token{START_COMMENT, "c", "o", "m", "m", "e", "n", "t", END_COMMENT}},
		{"'';\n", []Token{START_STRING, END_STRING, START_COMMENT, END_COMMENT}},
		{";5\n", []Token{START_COMMENT, "5", END_COMMENT}},
		{"`()", []Token{START_LIST, END_SEXP}},
		{"`(test)", []Token{START_LIST, START_NAME, "t", "e", "s", "t", END_NAME, END_SEXP}},
		{"(t (e \"st\"))", []Token{START_SEXP, START_NAME, "t", END_NAME, START_SEXP, START_NAME, "e", END_NAME, START_STRING, "s", "t", END_STRING, END_SEXP, END_SEXP}},
		{"(t (e) (s 't'))", []Token{START_SEXP, START_NAME, "t", END_NAME, START_SEXP, START_NAME, "e", END_NAME, END_SEXP, START_SEXP, START_NAME, "s", END_NAME, START_STRING, "t", END_STRING, END_SEXP, END_SEXP}},
		{"(t `('e' 'st'))", []Token{START_SEXP, START_NAME, "t", END_NAME, START_LIST, START_STRING, "e", END_STRING, START_STRING, "s", "t", END_STRING, END_SEXP, END_SEXP}},
		{"`(53 't' (e 'st'))", []Token{START_LIST, START_NUMBER, "5", "3", END_NUMBER, START_STRING, "t", END_STRING, START_SEXP, START_NAME, "e", END_NAME, START_STRING, "s", "t", END_STRING, END_SEXP, END_SEXP}},
		{"`(3.14) ; test", []Token{START_LIST, START_NUMBER, "3", ".", "1", "4", END_NUMBER, END_SEXP, START_COMMENT, " ", "t", "e", "s", "t", END_COMMENT}},
		{"(if (x) `(3.14) `('test'))", []Token{START_SEXP, START_NAME, "i", "f", END_NAME, START_SEXP, START_NAME, "x", END_NAME, END_SEXP, START_LIST, START_NUMBER, "3", ".", "1", "4", END_NUMBER, END_SEXP, START_LIST, START_STRING, "t", "e", "s", "t", END_STRING, END_SEXP, END_SEXP}},
	}
	for _, test := range tests {
		lexed, _ := Lex(test.Program)
		if len(lexed) != len(test.Lexed) {
			t.Errorf("Expected %d tokens got %d\n", len(test.Lexed), len(lexed))
		}
		for i, _ := range lexed {
			if lexed[i] != test.Lexed[i] {
				t.Errorf("Expected token %s got %s\n", test.Lexed[i], lexed[i])
			}
		}
	}
}
