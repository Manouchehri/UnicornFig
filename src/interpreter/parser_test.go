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
	etokens := []Token{START_NAME, "t", "e", "s", "t", END_NUMBER, END_NAME}
	err, value, newStart = ParseName(etokens, 0)
	if err == nil {
		t.Error("Expected to get an error parsing an invalid name")
	}
	if value.Type != UnassignedT {
		t.Error("Expected to parse a name, got unassigned type.")
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

func TestParseComment(t *testing.T) {
	tokens := []Token{START_COMMENT, END_COMMENT} // The only way it appears
	err, value, newStart := ParseComment(tokens)
	if err != nil {
		t.Error(err.Error())
	}
	if value.Type != UnassignedT {
		t.Errorf("Expected parsed comment value to have no type. Got %d\n", value.Type)
	}
	if newStart != len(tokens) {
		t.Error("Expected parsing comment to take us past comment tokens")
	}
}

func TestParseString(t *testing.T) {
	tokens := []Token{START_STRING, "t", "e", "s", "t", END_STRING}
	err, value, newStart := ParseString(tokens, 0)
	if err != nil {
		t.Error(err.Error())
	}
	if value.Type != StringT {
		t.Errorf("Expected to parse a string. Got %d\n", value.Type)
	}
	if value.String.Contained != "test" {
		t.Errorf("Expected to parse the string 'test'. Got %s\n", value.String.Contained)
	}
	etokens := []Token{START_STRING, "h", "i", END_NUMBER, END_STRING}
	err, value, newStart = ParseString(etokens, 0)
	if err == nil {
		t.Error("Expected to get an error parsing a malformed string.")
	}
	if value.Type != UnassignedT {
		t.Errorf("Failed parsings should result in an unassigned value. Got %d\n", value.Type)
	}
}

func TestParseSExpression(t *testing.T) {
	tokens := []Token{START_SEXP, START_NAME, "s", "q", END_NAME, START_NUMBER, "3", END_NUMBER, END_SEXP}
	err, sexp, newStart := ParseSExpression(tokens, 0)
	if err != nil {
		t.Error(err.Error())
	}
	if newStart != len(tokens) {
		t.Error("Parsing S-Expressions should take us past the list of tokens")
	}
	if sexp.FormName.Contained != "sq" {
		t.Errorf("Expected sexp to start with the form name 'sq'. Got %s\n", sexp.FormName.Contained)
	}
	if len(sexp.Values) != 1 {
		t.Errorf("Expected 1 value in sexp. Got len = %d\n", len(sexp.Values))
	}
	if (sexp.Values[0].(Value)).Contained != 3 {
		value := (sexp.Values[0].(Value)).Contained
		t.Errorf("Expected first value in sexp to be 3. Got %d\n", value)
	}
	etokens := []Token{START_SEXP, START_STRING, "s", "q", END_STRING, END_SEXP}
	err, sexp, newStart = ParseSExpression(etokens, 0)
	if err == nil {
		t.Error("Expected to get an error parsing an S-Expression that starts with a string")
	}
	e2tokens := []Token{START_SEXP, START_NAME, "3", "b", END_NAME, END_SEXP}
	err, sexp, newStart = ParseSExpression(e2tokens, 0)
	if err == nil {
		t.Error("Expeced to get an error parsing an S-Expression with invalid content")
	}
}
