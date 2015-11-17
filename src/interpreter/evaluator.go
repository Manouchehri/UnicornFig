package interpreter

import (
  "errors"
  "fmt"
)

// A type that contains information about values in a given scope
type Environment map[string]Value

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
    return EvaluateSexp(thing.(SExpression), env)
  default:
    return errors.New(fmt.Sprintf("No way to evaluate %v\n", thing)), Value{}, env
  }
}

func Apply(env Environment, fn Function, arguments ...Value) (error, Value, Environment) {
  // Check if the function maps to a builtin that can be executed as Go code.
  if fn.IsCallable {
    goValues := make([]interface{}, len(arguments))
    for i, arg := range arguments {
      goValues[i] = Unwrap(arg)
    }
    return fn.Call(goValues...)
  } else {
    return EvaluateSexp(fn.Body, env)
  }
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
  case FunctionT:
    // TODO - This is definitely NOT going to work in most places
    return value.Function.Callable
  }
  return nil
}
