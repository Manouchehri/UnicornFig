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
	if len(sexp.Values) != 3 {
		return errors.New("If expects one condition and two branches."), Value{}, env
	}
	conditionErr, conditionResult, newEnv := Evaluate(sexp.Values[0], env)
	if conditionErr != nil {
		return conditionErr, conditionResult, newEnv
	}
	if conditionResult.Type != BooleanT {
		return errors.New("Conditions for branching must evaluate to either true or false."), conditionResult, newEnv
	}
	if conditionResult.Boolean.Contained {
		return Evaluate(sexp.Values[1], newEnv)
	} else {
		return Evaluate(sexp.Values[2], newEnv)
	}
}

func EvaluateFunction(sexp SExpression, env Environment) (error, Value, Environment) {
	if len(sexp.Values) != 2 {
		errMsg := "Function declarations expect one S-Expression with a set of argument names and one with a body."
		return errors.New(errMsg), Value{}, env
	}
	switch sexp.Values[0].(type) {
	case SExpression:
		break
	default:
		return errors.New("Function argument names must be declared in an S-Expression."), Value{}, env
	}
	argumentNames := make([]string, 0)
	argumentList := sexp.Values[0].(SExpression)
	// Expect the character "_" to signify that a function takes no arguments, as opposed to an empty S-Expression
	if argumentList.FormName.Contained != "_" {
		argumentNames = append(argumentNames, argumentList.FormName.Contained)
		for i := 0; i < len(argumentList.Values); i++ {
			switch argumentList.Values[i].(type) {
			case Value:
				if argumentList.Values[i].(Value).Type != NameT {
					return errors.New("All items in a function argument list must be names."), Value{}, env
				}
			default:
				return errors.New("All items in a function argument list must be names."), Value{}, env
			}
			argumentNames = append(argumentNames, argumentList.Values[i].(Value).Name.Contained)
		}
	}
	newFn := NewFunction("tempname", argumentNames, sexp.Values[1])
	for k, v := range env {
		newFn.Function.Scope[k] = v
	}
	return nil, newFn, env
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
	for k, v := range env {
		function.Function.Scope[k] = v
	}
	value, err := Apply(function.Function, arguments...)
	return err, value, env
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

func Apply(fn Function, arguments ...Value) (Value, error) {
	for i := 0; i < len(fn.ArgumentNames); i++ {
		if i >= len(arguments) {
			return Value{}, errors.New("Not enough arguments passed to " + fn.FunctionName.Contained)
		}
		fn.Scope[fn.ArgumentNames[i].Contained] = arguments[i]
	}
	var err error
	var computedValue Value
	//var newEnv Environment
	if fn.IsCallable {
		goValues := make([]interface{}, len(arguments))
		for i, arg := range arguments {
			goValues[i] = Unwrap(arg)
		}
		computedValue, err = fn.Call(goValues...)
	} else {
		err, computedValue, _ = Evaluate(fn.Body, fn.Scope)
	}
	/*
		envDiff := environmentDifference(newEnv, localScope)
		for newKey, newValue := range envDiff {
			env[newKey] = newValue
		}
	*/
	if err != nil {
		return Value{}, err
	}
	err, computedValue, _ = EvaluateValue(computedValue, fn.Scope)
	return computedValue, err
}
