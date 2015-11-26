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
		mapping.Map.Data[key.(string)] = value.(uni.Value)
	}
	return nil, mapping, env
}
