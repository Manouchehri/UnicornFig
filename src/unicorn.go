package main

import (
	codegen "./codegen"
	uni "./interpreter"
	stdlib "./stdlib"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

const HelpMessage = `Run this program as ./unicorn [format file [format file [...]]] program.fig [program2.fig program3.fig ... programN.fig]
Currently the supported format flags are
	-json - Output program state to a JSON file
	-yaml - Output program state to a YAML file
	-go   - Output a Go source code file containing a Configuration struct and parser functions

At least one Fig program must be provided.

When more than one Fig program is provided, each will be run one after the other, and the
environment (global scope) produced by each will be made the environment of successive programs.
Therefore, one can run multiple Fig programs to effectively combine their outputs into a single
configuration.
`

var SupportedFormatHandlers = map[string]func(map[string]interface{}, string) error{
	"json": WriteJSON,
	"yaml": WriteYAML,
	"go":   codegen.GenerateConfigCodeFile,
}

func WriteJSON(env map[string]interface{}, fileName string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	bytes, encodeErr := json.MarshalIndent(env, "", "    ")
	if encodeErr != nil {
		return encodeErr
	}
	_, writeErr := f.Write(bytes)
	return writeErr
}

func WriteYAML(env map[string]interface{}, fileName string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	bytes, encodeErr := yaml.Marshal(env)
	if encodeErr != nil {
		return encodeErr
	}
	_, writeErr := f.Write(bytes)
	return writeErr
}

func WriteOutputFiles(formats map[string]string, data uni.Environment) error {
	// Strip out values that we can't encode, like functions, as well as constants defined in Unicorn.
	toWrite := make(map[string]interface{})
	for k, v := range data {
		if v.Ignored {
			continue
		}
		shouldContinue := false
		for _, constantName := range stdlib.ConstantNames {
			if k == constantName {
				shouldContinue = true
				break
			}
		}
		if shouldContinue {
			continue
		}
		unwrapped := uni.Unwrap(v)
		if unwrapped != nil {
			toWrite[k] = unwrapped
		}
	}
	for format, writer := range SupportedFormatHandlers {
		if fileName, shouldWrite := formats[format]; shouldWrite {
			if err := writer(toWrite, fileName); len(fileName) > 0 && err != nil {
				return err
			}
		}
	}
	return nil
}

func Interpret(program string, env uni.Environment) (uni.Environment, error) {
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
	//var value uni.Value
	//value := uni.Value{}
	for _, form := range parsedForms {
		err, _, env = uni.Evaluate(form, env)
		//err, value, env = uni.Evaluate(form, env)
		if err != nil {
			return uni.Environment{}, err
		}
	}
	return env, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(HelpMessage)
		return
	}
	// Maps are supported output file formats and values are names of files to write to if any.
	outputFormats := map[string]string{}
	for format, _ := range SupportedFormatHandlers {
		outputFormats[format] = ""
	}
	// Parse arguments in any form such as "--json output.json -YAML data.yaml myprogram.fig"
	i := 1
	for ; i < len(os.Args)-1; i++ {
		if !strings.HasPrefix(os.Args[i], "-") {
			break
		}
		format := strings.ToLower(strings.Replace(os.Args[i], "-", "", -1))
		_, isSupported := outputFormats[format]
		if isSupported {
			outputFormats[format] = os.Args[i+1]
			i++
		}
	}
	// Treat all arguments after the flags as source files
	env := uni.Environment{}
	if i == len(os.Args)-1 {
		fmt.Println("No input program file provided.")
		fmt.Println(HelpMessage)
		return
	}
	for ; i < len(os.Args); i++ {
		// Open and interpret the program file
		programFile := os.Args[i]
		fmt.Println("Trying to execute " + programFile)
		file, err := os.Open(programFile)
		if err != nil {
			fmt.Println("Couldn't open program file " + programFile)
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
		env, err = Interpret(program, env)
		if err != nil {
			fmt.Println("ERROR\n  ", err.Error())
		}
	}
	// Produce the desired output files
	err := WriteOutputFiles(outputFormats, env)
	if err != nil {
		fmt.Println("ERROR\n  ", err.Error())
	}
}
