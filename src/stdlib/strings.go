package stdlib

import (
	uni "../interpreter"
	"errors"
	"fmt"
	"strings"
)

func SLIB_Concatenate(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) < 2 {
		return uni.Value{}, errors.New("Concatenate function expects two or more arguments.")
	}
	result := arguments[0].(string)
	for i := 1; i < len(arguments); i++ {
		result += arguments[i].(string)
	}
	return uni.NewString(result), nil
}

func SLIB_Substring(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) != 3 {
		return uni.Value{}, errors.New("Susbtring function expects three arguments.")
	}
	str := arguments[0].(string)
	start := arguments[1].(int64)
	end := arguments[2].(int64)
	if start < 0 {
		return uni.Value{}, errors.New("Cannot start a substring at a negative index.")
	}
	if end > int64(len(str)) {
		return uni.Value{}, errors.New("Cannot end a substring past the end of the string's length.")
	}
	result := str[start:end]
	return uni.NewString(result), nil
}

func SLIB_Index(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) != 2 {
		return uni.Value{}, errors.New("Index function expects two arguments.")
	}
	first := arguments[0].(string)
	second := arguments[1].(string)
	index := strings.Index(first, second)
	return uni.NewInteger(int64(index)), nil
}

func SLIB_Length(arguments ...interface{}) (uni.Value, error) {
	length := len(arguments[0].(string))
	return uni.NewInteger(int64(length)), nil
}

func SLIB_Upcase(arguments ...interface{}) (uni.Value, error) {
	str := arguments[0].(string)
	str = strings.ToUpper(str)
	return uni.NewString(str), nil
}

func SLIB_Downcase(arguments ...interface{}) (uni.Value, error) {
	str := arguments[0].(string)
	str = strings.ToLower(str)
	return uni.NewString(str), nil
}

func SLIB_Split(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) != 2 {
		return uni.Value{}, errors.New("Split function expects two string arguments.")
	}
	for i := 0; i < 2; i++ {
		switch arguments[i].(type) {
		case string:
			break
		default:
			return uni.Value{}, errors.New("Split function expects oth arguments to be strings.")
		}
	}
	splitStrings := strings.Split(arguments[0].(string), arguments[1].(string))
	list := uni.NewList()
	for _, str := range splitStrings {
		wrapped, _ := uni.Wrap(str)
		list.List.Data = append(list.List.Data, wrapped)
	}
	return list, nil
}

func SLIB_AtIndex(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) != 2 {
		return uni.Value{}, errors.New("At function expects two arguments.")
	}
	switch arguments[0].(type) {
	case string:
		break
	default:
		return uni.Value{}, errors.New("At function expects first argument to be a string.")
	}
	switch arguments[1].(type) {
	case int64:
		break
	default:
		return uni.Value{}, errors.New("At function expects second argument to be an integer.")
	}
	str := arguments[0].(string)
	index := arguments[1].(int64)
	if index < int64(0) || index >= int64(len(str)) {
		errMsg := fmt.Sprintf("String index out of range. Cannot get index %d of \"%s\"", index, str)
		return uni.Value{}, errors.New(errMsg)
	}
	return uni.NewString(string(str[index])), nil
}
