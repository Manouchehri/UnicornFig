package interpreter

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
