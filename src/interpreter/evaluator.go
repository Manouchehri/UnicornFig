package interpreter

import (
	"errors"
	"fmt"
)

// A type that contains information about values in a given scope
type Environment map[string]Value

func isSpecialForm(formName string) bool {
	return formName == "define" || formName == "if" || formName == "function"
}

/**
 * Creates an environment consisting of only keys (and their corresponding values)
 * that are in env1 and not in env2.
 */
func environmentDifference(env1, env2 Environment) Environment {
	diff := Environment{}
	for key, value := range env1 {
		_, found := env2[key]
		if !found {
			diff[key] = value
		}
	}
	return diff
}

func EvaluateDefine(sexp SExpression, env Environment) (error, Value, Environment) {
	// Each value is (or should be) an S-Expression with a name to assign to and a value to evalute
	var lastValue Value
	for _, definition := range sexp.Values {
		switch definition.(type) {
		case SExpression:
			def := definition.(SExpression)
			if len(def.Values) != 1 {
				errMsg := "Definitions must be S-Expressions of the form (name <thing-to-evaluate>)."
				return errors.New(errMsg), Value{}, env
			}
			evalErr, value, newEnv := Evaluate(def.Values[0], env)
			if evalErr != nil {
				return evalErr, value, newEnv
			}
			lastValue = value
			newEnv[def.FormName.Contained] = value
			env = newEnv
		default:
			errMsg := "Pairs of names to assign to and their corresponding values must be contained in S-Expressions."
			return errors.New(errMsg), Value{}, env
		}
	}
	return nil, lastValue, env
}

func EvaluateIf(sexp SExpression, env Environment) (error, Value, Environment) {
	return nil, Value{}, env
}

func EvaluateFunction(sexp SExpression, env Environment) (error, Value, Environment) {
	return nil, Value{}, env
}

func EvaluateSpecialForm(sexp SExpression, env Environment) (error, Value, Environment) {
	switch sexp.FormName.Contained {
	case "define":
		return EvaluateDefine(sexp, env)
	case "if":
		return EvaluateIf(sexp, env)
	case "function":
		return EvaluateFunction(sexp, env)
	}
	return errors.New("Unrecognized special form " + sexp.FormName.Contained), Value{}, env
}

func EvaluateValue(value Value, env Environment) (error, Value, Environment) {
	if value.Type == NameT {
		varName := value.Name.Contained
		actual, found := env[varName]
		if !found {
			return errors.New("Variable " + varName + " not assigned."), Value{}, env
		} else {
			return nil, actual, env
		}
	} else {
		// Already a value
		return nil, value, env
	}
}

func EvaluateSexp(sexp SExpression, env Environment) (error, Value, Environment) {
	fnName := sexp.FormName.Contained
	function, found := env[fnName]
	if !found {
		return errors.New("No such function " + fnName), Value{}, env
	}
	arguments := make([]Value, 0)
	for _, arg := range sexp.Values {
		evalErr, value, newEnv := Evaluate(arg, env)
		if evalErr != nil {
			return evalErr, Value{}, newEnv
		}
		arguments = append(arguments, value)
	}
	return Apply(env, function.Function, arguments...)
}

func Evaluate(thing interface{}, env Environment) (error, Value, Environment) {
	switch thing.(type) {
	case Value:
		return EvaluateValue(thing.(Value), env)
	case SExpression:
		sexp := thing.(SExpression)
		if isSpecialForm(sexp.FormName.Contained) {
			return EvaluateSpecialForm(sexp, env)
		} else {
			return EvaluateSexp(thing.(SExpression), env)
		}
	default:
		return errors.New(fmt.Sprintf("No way to evaluate %v\n", thing)), Value{}, env
	}
}

func Apply(env Environment, fn Function, arguments ...Value) (error, Value, Environment) {
	// Check if the function maps to a builtin that can be executed as Go code.
	// Create a local scope
	localScope := Environment{}
	for key, value := range env {
		localScope[key] = value
	}
	for i := 0; i < len(fn.ArgumentNames); i++ {
		if i >= len(arguments) {
			return errors.New("Not enough arguments passed to " + fn.FunctionName.Contained), Value{}, env
		}
		localScope[fn.ArgumentNames[i].Contained] = arguments[i]
	}
	var err error
	var computedValue Value
	var newEnv Environment
	if fn.IsCallable {
		goValues := make([]interface{}, len(arguments))
		for i, arg := range arguments {
			goValues[i] = Unwrap(arg)
		}
		err, computedValue, newEnv = fn.Call(localScope, goValues...)
	} else {
		err, computedValue, newEnv = EvaluateSexp(fn.Body, localScope)
	}
	envDiff := environmentDifference(newEnv, localScope)
	for newKey, newValue := range envDiff {
		env[newKey] = newValue
	}
	if err != nil {
		return err, Value{}, env
	}
	err, computedValue, env = EvaluateValue(computedValue, env)
	return err, computedValue, env
}

func Unwrap(value Value) interface{} {
	switch value.Type {
	case StringT:
		return value.String.Contained
	case IntegerT:
		return value.Integer.Contained
	case FloatT:
		return value.Float.Contained
	case NameT:
		return value.Name.Contained
	case BooleanT:
		return value.Boolean.Contained
	case FunctionT:
		// TODO - This is definitely NOT going to work in most places
		return value.Function.Callable
	}
	return nil
}
