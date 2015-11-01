package main

import (
	"./interpreter"
	"fmt"
)

func main() {
	program := "123 dill 3.124"
	tokens := interpreter.Parse(program)
	fmt.Println(tokens)
}
