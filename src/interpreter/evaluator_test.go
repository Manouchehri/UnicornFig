package interpreter

import (
  "testing"
)

// Shorthand functions for longer struct initializers

var (
  zeroi = IntegerLiteral{0}
  zerof = FloatLiteral{0.0}
  emptys = StringLiteral{""}
)

func NewString(str string) Value {
	return Value{StringT, StringLiteral{str}, zeroi, zerof, Name{}, Function{}}
}

func NewInteger(n int64) Value {
	return Value{IntegerT, emptys, IntegerLiteral{n}, zerof, Name{}, Function{}}
}

func NewFloat(n float64) Value {
	return Value{FloatT, emptys, zeroi, FloatLiteral{n}, Name{}, Function{}}
}

func NewName(identifier string) Value {
	return Value{NameT, emptys, zeroi, zerof, Name{identifier}, Function{}}
}

func NewSExpression(formName string, values ...interface{}) SExpression {
  emptyArray := make([]interface{}, 0)
  sexp := SExpression{Name{formName}, SExpressionT, emptyArray}
  for _, value := range values {
    sexp.Values = append(sexp.Values, value)
  }
  return sexp
}

func NewCallableFunction(name string, argNames []string, fn Builtin) Value {
  names := make([]Name, len(argNames))
  for i, arg := range argNames {
    names[i] = Name{arg}
  }
  return Value{FunctionT, emptys, zeroi, zerof, Name{}, Function{Name{name}, names, SExpression{}, true, fn}}
}

func NewFunction(name string, argNames []string, body SExpression) Value {
  names := make([]Name, len(argNames))
  for i, arg := range argNames {
    names[i] = Name{arg}
  }
  return Value{FunctionT, emptys, zeroi, zerof, Name{}, Function{Name{name}, names, body, false, nil}}
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
  err3, _, _ := EvaluateValue(NewName("pi"), env)
  if err3 == nil {
    t.Error("Expected to get an error when evaluating a name that isn't in the environment")
  }
}

func TestApply(t *testing.T) {
  mult := func(args ...interface{}) (error, Value, Environment) {
    value := args[0].(Value).Integer.Contained * args[1].(Value).Integer.Contained
    return nil, NewInteger(value), Environment{}
  }
  env := Environment{
    "mult": NewCallableFunction("mult", []string{"a", "b"}, mult),
    "square": NewFunction("square", []string{"a"}, NewSExpression("mult", NewName("a"), NewName("a"))),
  }
  // Test that builtin functions can be invoked to get us a computed result
  err1, value1, newEnv1 := Apply(env, env["mult"].Function, NewInteger(10), NewInteger(3))
  if err1 != nil {
    t.Error(err1.Error())
  }
  if value1.Type != IntegerT {
    t.Error("Expected to get an integer as a result of calling mult")
  }
  if value1.Integer.Contained != 30 {
    t.Error("Expected 10 * 3 to be 30")
  }
  if len(newEnv1) != 2 {
    t.Error("Expected no items to be added or removed from the environment")
  }
  // Test that user-defined functions can be reached and a value computed
  err2, value2, newEnv2 := Apply(env, env["square"].Function, NewInteger(5))
  if err2 != nil {
    t.Error(err2.Error())
  }
  if value2.Type != IntegerT {
    t.Error("Expected result of calling square to be an integer")
  }
  if value2.Integer.Contained != 25 {
    t.Error("Expected square(5) to be 25")
  }
  if len(newEnv2) != 2 {
    t.Error("Expected no items to be added or removed from the environment")
  }
}

func TestEvaluateSexp(t *testing.T) {
  mult := func(args ...interface{}) (error, Value, Environment) {
    value := args[0].(Value).Integer.Contained * args[1].(Value).Integer.Contained
    return nil, NewInteger(value), Environment{}
  }
  env := Environment{
    "a": NewInteger(4),
    "b": NewInteger(2),
    "mult": NewCallableFunction("mult", []string{"a", "b"}, mult),
  }
  // (set a 4)
  // (set b 2)
  // (mult a b)
  err1, value1, newEnv1 := EvaluateSexp(NewSExpression("mult", NewName("a"), NewName("b")), env)
  if err1 != nil {
    t.Error(err1)
  }
  if value1.Type != IntegerT {
    t.Error("Expected to get an integer after calling mult")
  }
  if value1.Integer.Contained != 8 {
    t.Error("Expected 4 * 2 to be 8")
  }
  if len(newEnv1) != 3 {
    t.Error("Expected no items to be added to or removed from the environment")
  }
  err2, _, _ := EvaluateSexp(NewSExpression("add", NewName("a"), NewName("b")), env)
  if err2 == nil {
    t.Error("Expected to get an error trying to evaluate a function name that don't exist")
  }
}

func TestEvaluate(t *testing.T) {
  mult := func(args ...interface{}) (error, Value, Environment) {
    value := args[0].(Value).Integer.Contained * args[1].(Value).Integer.Contained
    return nil, NewInteger(value), Environment{}
  }
  env := Environment{
    "a": NewInteger(-4),
    "b": NewInteger(3),
    "mult": NewCallableFunction("mult", []string{"a", "b"}, mult),
  }
  err1, value1, _ := Evaluate(NewName("a"), env)
  if err1 != nil {
    t.Error(err1.Error())
  }
  if value1.Type != IntegerT {
    t.Error("Expected to evaluate a value to an integer")
  }
  if value1.Integer.Contained != 4 {
    t.Error("Expected the evaluated name a to contain the value 4")
  }
  err2, value2, _ := Evaluate(NewSExpression("mult", NewName("a"), NewName("b")), env)
  if err2 != nil {
    t.Error(err2.Error())
  }
  if value2.Type != IntegerT {
    t.Error("Expected result of calling mult to be an integer")
  }
  if value2.Integer.Contained != -12 {
    t.Errorf("Expected -4 * 3 to be -12. Got %d\n", value2.Integer.Contained)
  }
  err3, _, _ := Evaluate(NewName("x"), env)
  if err3 == nil {
    t.Error("Expected to get an error evaluating name that isn'tin the environment")
  }
}

func TestUnwrap(t *testing.T) {
  a := NewInteger(10)
  pi := NewFloat(3.14)
  name := NewString("Alice")
  if Unwrap(a).(int64) != 10 {
    t.Error("Expected unwrapped integer to have value 10")
  }
  if Unwrap(pi).(float64) != 3.14 {
    t.Error("Expected unwrapped float to have value 3.14")
  }
  if Unwrap(name).(string) != "Alice" {
    t.Error("Expected unwrapped string to have value 'Alice'")
  }
}
