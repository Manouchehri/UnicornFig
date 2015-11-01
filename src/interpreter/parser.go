package interpreter

import (
	"errors"
	"strconv"
	"strings"
)

type SyntaxTree struct {
	Root SExpression
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
