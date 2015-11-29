package stdlib

import (
	uni "../interpreter"
	"errors"
)

func SLIB_List(arguments ...interface{}) (uni.Value, error) {
	list := uni.NewList()
	if len(arguments) == 0 {
		return list, nil
	}
	for _, value := range arguments {
		wrapped, err := uni.Wrap(value)
		if err != nil {
			return list, err
		}
		list.List.Data = append(list.List.Data, wrapped)
	}
	return list, nil
}

func SLIB_First(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) != 1 {
		return uni.Value{}, errors.New("First function expects one list argument.")
	}
	switch arguments[0].(type) {
	case []interface{}:
		break
	default:
		return uni.Value{}, errors.New("First expects a list of values.")
	}
	values := arguments[0].([]interface{})
	if len(values) == 0 {
		return uni.Value{}, errors.New("First expects a list with at least one value.")
	}
	wrapped, err := uni.Wrap(values[0])
	if err != nil {
		return uni.Value{}, err
	}
	return wrapped, nil
}

func SLIB_Tail(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) != 1 {
		return uni.Value{}, errors.New("Tail function expects one list argument.")
	}
	switch arguments[0].(type) {
	case []interface{}:
		break
	default:
		return uni.Value{}, errors.New("Tail function expects a list of values.")
	}
	values := arguments[0].([]interface{})
	if len(values) == 0 {
		return uni.Value{}, errors.New("Tail expects a list with at least one value.")
	}
	if len(values) == 1 {
		return uni.NewList(), nil
	}
	list := uni.NewList()
	for i := 1; i < len(values); i++ {
		wrapped, err := uni.Wrap(values[i])
		if err != nil {
			return uni.Value{}, err
		}
		list.List.Data = append(list.List.Data, wrapped)
	}
	return list, nil
}

func SLIB_Append(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) < 2 {
		return uni.Value{}, errors.New("Append function expects a list and at least one value to append.")
	}
	switch arguments[0].(type) {
	case []interface{}:
		break
	default:
		return uni.Value{}, errors.New("Append function expects first argument to be a list.")
	}
	values := arguments[0].([]interface{})
	list := uni.NewList()
	wrappedValues := make([]uni.Value, len(values)+len(arguments)-1)
	for i := 0; i < len(values); i++ {
		wrapped, err := uni.Wrap(values[i])
		if err != nil {
			return list, err
		}
		wrappedValues[i] = wrapped
	}
	for i := 1; i < len(arguments); i++ {
		wrapped, err := uni.Wrap(arguments[i])
		if err != nil {
			return list, err
		}
		wrappedValues[len(values)+i-1] = wrapped
	}
	list.List.Data = append(list.List.Data, wrappedValues...)
	return list, nil
}

func SLIB_Size(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) != 1 {
		return uni.Value{}, errors.New("List Size function expects exactly one list argument.")
	}
	switch arguments[0].(type) {
	case []interface{}:
		break
	default:
		return uni.Value{}, errors.New("List Size function expects first argument to be a list.")
	}
	length := int64(len(arguments[0].([]interface{})))
	return uni.NewInteger(length), nil
}
