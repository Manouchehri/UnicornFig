package stdlib

import (
	uni "../interpreter"
	"errors"
)

func SLIB_List(env uni.Environment, arguments ...interface{}) (error, uni.Value, uni.Environment) {
	list := uni.NewList()
	if len(arguments) == 0 {
		return nil, list, env
	}
	for _, value := range arguments {
		wrapped, err := uni.Wrap(value)
		if err != nil {
			return err, list, env
		}
		list.List.Data = append(list.List.Data, wrapped)
	}
	return nil, list, env
}

func SLIB_First(env uni.Environment, arguments ...interface{}) (error, uni.Value, uni.Environment) {
	if len(arguments) != 1 {
		return errors.New("First function expects one list argument."), uni.Value{}, env
	}
	switch arguments[0].(type) {
	case []interface{}:
		break
	default:
		return errors.New("First expects a list of values."), uni.Value{}, env
	}
	values := arguments[0].([]interface{})
	if len(values) == 0 {
		return errors.New("First expects a list with at least one value."), uni.Value{}, env
	}
	wrapped, err := uni.Wrap(values[0])
	if err != nil {
		return err, uni.Value{}, env
	}
	return nil, wrapped, env
}

func SLIB_Tail(env uni.Environment, arguments ...interface{}) (error, uni.Value, uni.Environment) {
	if len(arguments) != 1 {
		return errors.New("Tail function expects one list argument."), uni.Value{}, env
	}
	switch arguments[0].(type) {
	case []interface{}:
		break
	default:
		return errors.New("Tail function expects a list of values."), uni.Value{}, env
	}
	values := arguments[0].([]interface{})
	if len(values) == 0 {
		return errors.New("Tail expects a list with at least one value."), uni.Value{}, env
	}
	if len(values) == 1 {
		return nil, uni.NewList(), env
	}
	list := uni.NewList()
	for i := 1; i < len(values); i++ {
		wrapped, err := uni.Wrap(values[i])
		if err != nil {
			return err, uni.Value{}, env
		}
		list.List.Data = append(list.List.Data, wrapped)
	}
	return nil, list, env
}

func SLIB_Append(env uni.Environment, arguments ...interface{}) (error, uni.Value, uni.Environment) {
	if len(arguments) < 2 {
		return errors.New("Append function expects a list and at least one value to append."), uni.Value{}, env
	}
	switch arguments[0].(type) {
	case []interface{}:
		break
	default:
		return errors.New("Append function expects first argument to be a list."), uni.Value{}, env
	}
	values := arguments[0].([]interface{})
	list := uni.NewList()
	wrappedValues := make([]uni.Value, len(values)+len(arguments)-1)
	for i := 0; i < len(values); i++ {
		wrapped, err := uni.Wrap(values[i])
		if err != nil {
			return err, list, env
		}
		wrappedValues[i] = wrapped
	}
	for i := 1; i < len(arguments); i++ {
		wrapped, err := uni.Wrap(arguments[i])
		if err != nil {
			return err, list, env
		}
		wrappedValues[len(values)+i-1] = wrapped
	}
	list.List.Data = append(list.List.Data, wrappedValues...)
	return nil, list, env
}
