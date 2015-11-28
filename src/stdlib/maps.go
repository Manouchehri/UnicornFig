package stdlib

import (
	uni "../interpreter"
	"errors"
)

func SLIB_Map(env uni.Environment, arguments ...interface{}) (error, uni.Value, uni.Environment) {
	mapping := uni.NewMap()
	if len(arguments) == 0 {
		return nil, mapping, env
	}
	if len(arguments)%2 == 1 {
		return errors.New("Must have an even number of arguments to create a map."), mapping, env
	}
	for i := 0; i < len(arguments); i += 2 {
		key := arguments[i]
		value := arguments[i+1]
		switch key.(type) {
		case string:
			break
		default:
			return errors.New("All keys must be strings."), mapping, env
		}
		wrapped, err := uni.Wrap(value)
		if err != nil {
			return err, mapping, env
		}
		mapping.Map.Data[key.(string)] = wrapped
	}
	return nil, mapping, env
}

func SLIB_Associate(env uni.Environment, arguments ...interface{}) (error, uni.Value, uni.Environment) {
	mapping := uni.NewMap()
	if len(arguments) <= 1 || len(arguments)%2 == 0 {
		return errors.New("Associate function expects a map and at least one key-value pair of arguments."), mapping, env
	}
	// Recreate the exsting map
	switch arguments[0].(type) {
	case map[string]interface{}:
		break
	default:
		return errors.New("Associate function expects its first argument to be a map."), mapping, env
	}
	for k, v := range arguments[0].(map[string]interface{}) {
		wrapped, err := uni.Wrap(v)
		if err != nil {
			return err, mapping, env
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
			return errors.New("Associate function expects all the new keys to be strings."), mapping, env
		}
		wrapped, err := uni.Wrap(v)
		if err != nil {
			return err, mapping, env
		}
		mapping.Map.Data[k.(string)] = wrapped
	}
	return nil, mapping, env
}

func SLIB_GetMap(env uni.Environment, arguments ...interface{}) (error, uni.Value, uni.Environment) {
	if len(arguments) != 2 {
		return errors.New("Get function expects a map and a key argument."), uni.Value{}, env
	}
	switch arguments[0].(type) {
	case map[string]interface{}:
		break
	default:
		return errors.New("Get function expects first argument to be a map."), uni.Value{}, env
	}
	switch arguments[1].(type) {
	case string:
		break
	default:
		return errors.New("Get function expects second argument to be a string key."), uni.Value{}, env
	}
	mapping := arguments[0].(map[string]interface{})
	key := arguments[1].(string)
	wrapped, err := uni.Wrap(mapping[key])
	return err, wrapped, env
}
