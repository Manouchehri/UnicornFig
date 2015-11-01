package interpreter

import (
	"testing"
)

func TestParseName(t *testing.T) {
	tokens := []Token{START_NAME, "t", "e", "s", "t", END_NAME}
	err, value, newStart := ParseName(tokens, 0)
	if err != nil {
		t.Error(err.Error())
	}
	if value.Type == UnassignedT {
		t.Error("Expected to parse a name, got unassigned type.")
	}
	if newStart != len(tokens) {
		t.Error("Expected parser to move past end of name")
	}
	if value.Name.Contained != "test" {
		t.Errorf("Expected name to contain 'test'. Got %s\n", value.Name.Contained)
	}
}

func TestParseNumber(t *testing.T) {
	fTokens := []Token{START_NUMBER, "3", ".", "1", "4", END_NUMBER}
	iTokens := []Token{START_NUMBER, "3", "2", "1", END_NUMBER}
	eTokens := []Token{START_NUMBER, "3", "2", "1", END_SEXP, END_NUMBER}
	err, value, newStart := ParseNumber(fTokens, 0)
	if err != nil {
		t.Error(err.Error())
	}
	if value.Type != FloatT {
		t.Errorf("Expected to parse a Float. Got type %d\n", value.Type)
	}
	if newStart != len(fTokens) {
		t.Error("Expected parser to move past end of number")
	}
	if value.Float.Contained != 3.14 {
		t.Errorf("Expected float to contain 3.14. Got %f\n", value.Float.Contained)
	}
	err, value, newStart = ParseNumber(iTokens, 0)
	if err != nil {
		t.Error(err.Error())
	}
	if value.Type != IntegerT {
		t.Errorf("Expected to parse an Integer. Got type %d\n", value.Type)
	}
	if newStart != len(iTokens) {
		t.Error("Expected parser to move past end of number")
	}
	if value.Integer.Contained != 321 {
		t.Errorf("Expected float to contain 321. Got %d\n", value.Integer.Contained)
	}
	err, value, newStart = ParseNumber(eTokens, 0)
	if err == nil {
		t.Error("Expected to get an error parsing an invalid number")
	}
	if value.Type != UnassignedT {
		t.Errorf("Erroneous parsings should result in an unassigned value. Got type %d\n", value.Type)
	}
}
