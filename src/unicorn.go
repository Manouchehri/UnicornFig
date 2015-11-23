package main

import (
	uni "./interpreter"
	stdlib "./stdlib"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func Interpret(program string) (uni.Environment, error) {
	env := uni.Environment{}
	// Copy the standard library into the local scope so we don't corrupt the former
	for key, value := range stdlib.StandardLibrary {
		env[key] = value
	}
	lexed, length := uni.Lex(program, 0)
	if length != len(program) {
		return env, errors.New("Could not lex to the end of your program. Check that it is properly formatted.")
	}
	parseErr, parsedForms := uni.Parse(lexed)
	if parseErr != nil {
		return env, parseErr
	}
	var err error = nil
	//value := uni.Value{}
	for _, form := range parsedForms {
		//err, value, env = uni.Evaluate(form, env)
		err, _, env = uni.Evaluate(form, env)
		if err != nil {
			return uni.Environment{}, err
		}
		//fmt.Println("sexp", i, "value", value, "\n")
	}
	return env, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Run this program with %s <program.fig>\n", os.Args[0])
	} else {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		programBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		program := string(programBytes)
		_, err = Interpret(program)
		if err != nil {
			fmt.Println("ERROR\n  ", err.Error())
		}
	}
}
