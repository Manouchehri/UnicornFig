package stdlib

import (
	uni "../interpreter"
	"errors"
	"strings"
)

func SLIB_Concatenate(env uni.Environment, arguments ...interface{}) (error, uni.Value, uni.Environment) {
	if len(arguments) < 2 {
		return errors.New("Concatenate function expects two or more arguments."), uni.Value{}, env
	}
	result := arguments[0].(string)
	for i := 1; i < len(arguments); i++ {
		result += arguments[i].(string)
	}
	return nil, uni.NewString(result), env
}

func SLIB_Substring(env uni.Environment, arguments ...interface{}) (error, uni.Value, uni.Environment) {
	if len(arguments) != 3 {
		return errors.New("Susbtring function expects three arguments."), uni.Value{}, env
	}
	str := arguments[0].(string)
	start := arguments[1].(int64)
	end := arguments[2].(int64)
	if start < 0 {
		return errors.New("Cannot start a substring at a negative index."), uni.Value{}, env
	}
	if end > int64(len(str)) {
		return errors.New("Cannot end a substring past the end of the string's length."), uni.Value{}, env
	}
	result := str[start:end]
	return nil, uni.NewString(result), env
}

func SLIB_Index(env uni.Environment, arguments ...interface{}) (error, uni.Value, uni.Environment) {
	if len(arguments) != 2 {
		return errors.New("Index function expects two arguments."), uni.Value{}, env
	}
	first := arguments[0].(string)
	second := arguments[1].(string)
	index := strings.Index(first, second)
	return nil, uni.NewInteger(int64(index)), env
}

func SLIB_Length(env uni.Environment, arguments ...interface{}) (error, uni.Value, uni.Environment) {
	length := len(arguments[0].(string))
	return nil, uni.NewInteger(int64(length)), env
}

func SLIB_Upcase(env uni.Environment, arguments ...interface{}) (error, uni.Value, uni.Environment) {
	str := arguments[0].(string)
	str = strings.ToUpper(str)
	return nil, uni.NewString(str), env
}

func SLIB_Downcase(env uni.Environment, arguments ...interface{}) (error, uni.Value, uni.Environment) {
	str := arguments[0].(string)
	str = strings.ToLower(str)
	return nil, uni.NewString(str), env
}
