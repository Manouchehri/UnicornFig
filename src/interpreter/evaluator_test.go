package main

import (
  "testing"
)

// Shorthand functions for longer struct initializers

func NewString(str string) Value {
	return Value{StringT, str, 0, 0.0, Name{}, Function{}}
}

func NewInteger(n int64) Value {
	return Value{IntegerT, "", n, 0.0, Name{}, Function{}}
}

func NewFloat(n float64) Value {
	return Value{FloatT, "", o, n, Name{}, Function{}}
}

func NewName(identifier string) Value {
	return Value{NameT, "", 0, 0.0, Name{identifier}, Function{}}
}

func NewSExpression(formName string, values ...interface{}) SExpression {
  sexp := SExpression{NewName(formName), SExpressionT, []interface{}}
  for _, value := range values {
    sexp.Values = append(sexp.Values, value)
  }
  return sexp
}

func NewCallableFunction(name string, argNames []string, fn Builtin) Function {
  names := make([]Name, len(argNames))
  for i, arg := range argNames {
    names[i] = NewName(arg)
  }
  return Function{NewName(name), names, SExpression{}, true, fn}
}

func NewFunction(name string, argNames []string, body SExpression) Function {
  names := make([]Name, len(argNames))
  for i, arg := range argNames {
    names[i] = NewName(arg)
  }
  return Function{NewName(name), names, body, false, nil}
}

// Actual test code

func TestEvaluateValue(t *testing.T) {
  env := Environment{
    "name": NewString("Alice"),
  }
  // Test that names get mapped to values in an environment
  err1, value1, newEnv1 := EvaluateValue(NewName("name"), env)
  if err1 != nil {
    t.Error(err1.Error())
  }
  if value1.Type != StringT {
    t.Error("Expected to get a string after evaluating the name `name`")
  }
  if value1.String.Contained != "Alice" {
    t.Error("Expected to get the string `Alice`")
  }
  if len(newEnv1) != 2 {
    t.Error("Expected no items to be added to or removed from the environment after name evaluation")
  }
  // Test that evaluating regular values just gets us back the value
  err2, value2, newEnv2 := EvaluateValue(NewInteger(12), env)
  if err2 != nil {
    t.Error(err2.Error())
  }
  if value2.Type != IntegerT {
    t.Error("Expected to get the integer we evaluated.")
  }
  if value2.Integer.Contained != 12 {
    t.Errorf("Expected the integer we found to contain 3. Got %d\n", value2.Integer.Contained)
  }
  if len(newEnv2) != 2 {
    t.Error("Expected no items to be added to or removed from the environment after integer evaluation")
  }
  // Test that if we evaluate a name that isn't in the environment, we get an error
  err3, _, _ := EvaluateValue(NewName("pi"))
  if err3 == nil {
    t.Error("Expected to get an error when evaluating a name that isn't in the environment")
  }
}


func TestApply(t *testing.T) {
  mult := func(args ...interface{}) Value {
    value := args[0].(Value).Integer.Contained * args[1].(Value).Integer.Contained
    return NewInteger(value)
  }
  env := Environment{
    "a": NewInteger(10),
    "b": NewInteger(3),
    "mult": NewCallableFunction("mult", []string{"a", "b"}, mult),
    "square": NewFunction("square", []string{"a"}, NewSExpression("mult", NewName("a"), NewName("a"))),
  }
  // Test that builtin functions can be invoked to get us a computed result
  err1, value1, newEnv1 := Apply(env, env["mult"], NewInteger(3), NewInteger(2))
  if err1 != nil {
    t.Error(err1.Error())
  }
  if value1.Type != IntegerT {
    t.Error("Expected to get an integer as a result of calling mult")
  }
  if value1.Integer.Contained != 6 {
    t.Error("Expected 3 * 2 to be 6")
  }
  if len(newEnv1) != 4 {
    t.Error("Expected no items to be added or removed from the environment")
  }
  // Test that user-defined functions can be reached and a value computed
  err2, value2, newEnv2 := Apply(env, env["square"], NewInteger(5))
  if err2 != nil {
    t.Error(err2.Error())
  }
  if value2.Type != IntegerT {
    t.Error("Expected result of calling square to be an integer")
  }
  if value2.Integer.Contained != 25 {
    t.Error("Expected square(5) to be 25")
  }
  if len(newEnv2) != 4 {
    t.Error("Expected no items to be added or removed from the environment")
  }
}
