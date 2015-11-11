package interpreter

import (
	"errors"
	"strconv"
	"strings"
)

type SyntaxTree struct {
	Root SExpression
}

type SimpleParser func([]Token, int) (error, Value, int)

var SimpleParsersTable = map[Token]SimpleParser{
	START_STRING: ParseString,
	START_NUMBER: ParseNumber,
	START_NAME:   ParseName,
}

func ParseName(tokens []Token, i int) (error, Value, int) {
	value := Value{}
	value.Type = UnassignedT
	if tokens[i] != START_NAME {
		errMsg := "Expected START_NAME, got " + string(tokens[i])
		return errors.New(errMsg), value, i
	}
	name := Name{""}
	i++
	for tokens[i] != END_NAME {
		if len(tokens[i]) > 1 {
			errMsg := "Expected token or END_NAME. Found " + string(tokens[i])
			return errors.New(errMsg), value, i
		}
		name.Contained += string(tokens[i])
		i++
	}
	value.Type = NameT
	value.Name = name
	return nil, value, i + 1
}

func ParseNumber(tokens []Token, i int) (error, Value, int) {
	value := Value{}
	value.Type = UnassignedT
	if tokens[i] != START_NUMBER {
		errMsg := "Expected START_NUMBER, got " + string(tokens[i])
		return errors.New(errMsg), value, i
	}
	numberStr := ""
	i++
	for tokens[i] != END_NUMBER {
		if len(tokens[i]) > 1 {
			errMsg := "Expected token or END_NUMBER. Found " + string(tokens[i])
			return errors.New(errMsg), value, i
		}
		numberStr += string(tokens[i])
		i++
	}
	if strings.Contains(numberStr, ".") {
		value.Type = FloatT
		f, _ := strconv.ParseFloat(numberStr, 64)
		value.Float = FloatLiteral{f}
	} else {
		value.Type = IntegerT
		i, _ := strconv.ParseInt(numberStr, 10, 64)
		value.Integer = IntegerLiteral{i}
	}
	return nil, value, i + 1
}

func ParseComment(tokens []Token, i int) (error, Value, int) {
	value := Value{}
	value.Type = UnassignedT
	if tokens[i] != START_COMMENT {
		errMsg := "Expected START_COMMENT, got " + string(tokens[i])
		return errors.New(errMsg), value, i
	}
	for tokens[i] != END_COMMENT {
		i++
	}
	return nil, value, i + 1
}

func ParseString(tokens []Token, i int) (error, Value, int) {
	value := Value{}
	value.Type = UnassignedT
	if tokens[i] != START_STRING {
		errMsg := "Expected START_STRING, got " + string(tokens[i])
		return errors.New(errMsg), value, i
	}
	str := ""
	i++
	for tokens[i] != END_STRING {
		if len(tokens[i]) > 1 {
			errMsg := "Expected token or END_STRING. Found " + string(tokens[i])
			return errors.New(errMsg), value, i
		}
		str += string(tokens[i])
		i++
	}
	value.Type = StringT
	value.String = StringLiteral{str}
	return nil, value, i + 1
}

func ParseSExpression(tokens []Token, i int) (error, SExpression, int) {
	sexp := SExpression{}
	sexp.ContainedType = SExpressionT
	if tokens[i] != START_SEXP {
		errMsg := "Expected START_SEXP, got " + string(tokens[i])
		return errors.New(errMsg), sexp, i
	}
	i++
	formErr, formName, newStart := ParseName(tokens, i)
	if formErr != nil {
		return formErr, sexp, i
	}
	sexp.FormName = formName.Name
	i = newStart
	for tokens[i] != END_SEXP {
		simpleParser, found := SimpleParsersTable[tokens[i]]
		var parseErr error = nil
		if found {
			err, value, nextIndex := simpleParser(tokens, i)
			sexp.Values = append(sexp.Values, value)
			parseErr = err
			i = nextIndex - 1 // We'll increment i at the end of the loop
		} else if tokens[i] == START_COMMENT {
			err, _, nextIndex := ParseComment(tokens, i)
			parseErr = err
			i = nextIndex - 1
		} else {
			err, innerSexp, nextIndex := ParseSExpression(tokens, i)
			parseErr = err
			i = nextIndex - 1
			sexp.Values = append(sexp.Values, innerSexp)
		}
		if parseErr != nil {
			return parseErr, sexp, i
		}
		i++
	}
	return nil, sexp, i + 1
}
