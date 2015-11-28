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
	tests := [...]string{"\n", "(", "'", "\"", ";", "4", "H", ")"}
	for i, transition := range TransitionsFromOpen {
		matched, err := regexp.MatchString(transition.ReadMatch, tests[i])
		if err != nil {
			t.Error(err.Error())
		} else if !matched {
			t.Errorf("%s did not match %s\n", transition.ReadMatch, tests[i])
		}
	}
}

func TestTransitionsFromString1(t *testing.T) {
	tests := [...]string{"\"", "'", " "}
	for i, transition := range TransitionsFromString1 {
		matched, err := regexp.MatchString(transition.ReadMatch, tests[i])
		if err != nil {
			t.Error(err.Error())
		} else if !matched {
			t.Errorf("%s did not match %s\n", transition.ReadMatch, tests[i])
		}
	}
}

func TestTransitionsFromString2(t *testing.T) {
	tests := [...]string{"'", "\"", "f"}
	for i, transition := range TransitionsFromString2 {
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
	tests := [...]string{"\n", ")", "3"}
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
	tests := [...]string{"\t", ")", "_"}
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
		{OPEN, STRING2, "\"", DoNothing, []Token{START_STRING}},
		{OPEN, STRING1, "'", DoNothing, []Token{START_STRING}},
		{OPEN, NUMBER, "0", DoNothing, []Token{START_NUMBER, "0"}},
		{OPEN, OPEN, ")", Return, []Token{END_SEXP}},
		{OPEN, NAME, "+", DoNothing, []Token{START_NAME, "+"}},
		{STRING1, OPEN, "'", DoNothing, []Token{END_STRING}},
		{STRING2, OPEN, "\"", DoNothing, []Token{END_STRING}},
		{STRING1, STRING1, "\"", DoNothing, []Token{"\""}},
		{STRING2, STRING2, "'", DoNothing, []Token{"'"}},
		{COMMENT, OPEN, "\n", DoNothing, []Token{END_COMMENT}},
		{NUMBER, OPEN, "\t", DoNothing, []Token{END_NUMBER}},
		{NAME, NAME, "f", DoNothing, []Token{Token("f")}},
		{NAME, NAME, "?", DoNothing, []Token{Token("?")}},
		{NAME, OPEN, " ", DoNothing, []Token{END_NAME}},
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
		{"(3)", []Token{START_SEXP, START_NUMBER, "3", END_NUMBER, END_SEXP}},
		{"(2.2)", []Token{START_SEXP, START_NUMBER, "2", ".", "2", END_NUMBER, END_SEXP}},
		{"('test')", []Token{START_SEXP, START_STRING, "t", "e", "s", "t", END_STRING, END_SEXP}},
		{"\"'\"", []Token{START_STRING, "'", END_STRING}},
		{"'\"'", []Token{START_STRING, "\"", END_STRING}},
		{"(hi)", []Token{START_SEXP, START_NAME, "h", "i", END_NAME, END_SEXP}},
		{";comment\n", []Token{START_COMMENT, END_COMMENT}},
		{"'';\n", []Token{START_STRING, END_STRING, START_COMMENT, END_COMMENT}},
		{";5\n", []Token{START_COMMENT, END_COMMENT}},
		{"(t (e \"st\"))", []Token{START_SEXP, START_NAME, "t", END_NAME, START_SEXP, START_NAME, "e", END_NAME, START_STRING, "s", "t", END_STRING, END_SEXP, END_SEXP}},
		{"(t (e) (s 't'))", []Token{START_SEXP, START_NAME, "t", END_NAME, START_SEXP, START_NAME, "e", END_NAME, END_SEXP, START_SEXP, START_NAME, "s", END_NAME, START_STRING, "t", END_STRING, END_SEXP, END_SEXP}},
		{"(hi 't' (e 'st'))", []Token{START_SEXP, START_NAME, "h", "i", END_NAME, START_STRING, "t", END_STRING, START_SEXP, START_NAME, "e", END_NAME, START_STRING, "s", "t", END_STRING, END_SEXP, END_SEXP}},
		{"(x) ; test\n", []Token{START_SEXP, START_NAME, "x", END_NAME, END_SEXP, START_COMMENT, END_COMMENT}},
		{"(if (x) 3.14 'test')", []Token{START_SEXP, START_NAME, "i", "f", END_NAME, START_SEXP, START_NAME, "x", END_NAME, END_SEXP, START_NUMBER, "3", ".", "1", "4", END_NUMBER, START_STRING, "t", "e", "s", "t", END_STRING, END_SEXP}},
	}
	for _, test := range tests {
		lexed, _ := Lex(test.Program, 0)
		if len(lexed) != len(test.Lexed) {
			t.Log("Lexing program: " + test.Program)
			t.Log(lexed)
			t.Errorf("Expected %d tokens got %d\n", len(test.Lexed), len(lexed))
			return
		}
		for i, _ := range lexed {
			if lexed[i] != test.Lexed[i] {
				t.Log("Lexing program: " + test.Program)
				t.Errorf("Expected token %s got %s\n", test.Lexed[i], lexed[i])
				return
			}
		}
	}
}
