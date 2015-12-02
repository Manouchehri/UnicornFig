package stdlib

import (
	uni "../interpreter"
	"errors"
)

func SLIB_Map(arguments ...interface{}) (uni.Value, error) {
	mapping := uni.NewMap()
	if len(arguments) == 0 {
		return mapping, nil
	}
	if len(arguments)%2 == 1 {
		return mapping, errors.New("Must have an even number of arguments to create a map.")
	}
	for i := 0; i < len(arguments); i += 2 {
		key := arguments[i]
		value := arguments[i+1]
		switch key.(type) {
		case string:
			break
		default:
			return mapping, errors.New("All keys must be strings.")
		}
		wrapped, err := uni.Wrap(value)
		if err != nil {
			return mapping, err
		}
		mapping.Map.Data[key.(string)] = wrapped
	}
	return mapping, nil
}

func SLIB_Associate(arguments ...interface{}) (uni.Value, error) {
	mapping := uni.NewMap()
	if len(arguments) <= 1 || len(arguments)%2 == 0 {
		return mapping, errors.New("Associate function expects a map and at least one key-value pair of arguments.")
	}
	// Recreate the exsting map
	switch arguments[0].(type) {
	case map[string]interface{}:
		break
	default:
		return mapping, errors.New("Associate function expects its first argument to be a map.")
	}
	for k, v := range arguments[0].(map[string]interface{}) {
		wrapped, err := uni.Wrap(v)
		if err != nil {
			return mapping, err
		}
		mapping.Map.Data[k] = wrapped
	}
	// Add all the new key-value pairs
	for i := 1; i < len(arguments); i += 2 {
		k := arguments[i]
		v := arguments[i+1]
		switch k.(type) {
		case string:
			break
		default:
			return mapping, errors.New("Associate function expects all the new keys to be strings.")
		}
		wrapped, err := uni.Wrap(v)
		if err != nil {
			return mapping, err
		}
		mapping.Map.Data[k.(string)] = wrapped
	}
	return mapping, nil
}

func SLIB_GetMap(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) != 2 {
		return uni.Value{}, errors.New("Get function expects a map and a key argument.")
	}
	switch arguments[0].(type) {
	case map[string]interface{}:
		break
	default:
		return uni.Value{}, errors.New("Get function expects first argument to be a map.")
	}
	switch arguments[1].(type) {
	case string:
		break
	default:
		return uni.Value{}, errors.New("Get function expects second argument to be a string key.")
	}
	mapping := arguments[0].(map[string]interface{})
	key := arguments[1].(string)
	wrapped, err := uni.Wrap(mapping[key])
	return wrapped, err
}

func SLIB_Keys(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) != 1 {
		return uni.Value{}, errors.New("Keys function expects a single map argument.")
	}
	switch arguments[0].(type) {
	case map[string]interface{}:
		break
	default:
		return uni.Value{}, errors.New("Keys function expects its argument to be a map.")
	}
	mapping := arguments[0].(map[string]interface{})
	list := uni.NewList()
	for k, _ := range mapping {
		wrapped, _ := uni.Wrap(k)
		list.List.Data = append(list.List.Data, wrapped)
	}
	return list, nil

}
