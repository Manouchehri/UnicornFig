package main

import (
	uni "./interpreter"
	stdlib "./stdlib"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

var SupportedFormatHandlers = map[string]func(map[string]interface{}, string) error{
	"json": WriteJSON,
	"yaml": WriteYAML,
	"xml":  WriteXML,
}

func WriteJSON(env map[string]interface{}, fileName string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	bytes, encodeErr := json.Marshal(env)
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

// Setting becomes the name of the tag surrounding each <name> and <value> tag
type Setting struct {
	Name  string      `xml:"name"`
	Value interface{} `xml:"value"`
}

func WriteXML(env map[string]interface{}, fileName string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	data := make([]Setting, len(env))
	index := 0
	for k, v := range env {
		data[index] = Setting{k, v}
		index++
	}
	bytes, encodeErr := xml.Marshal(data)
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
		// TODO - Write a more complete help string
		fmt.Printf("Run this program with %s <program.fig>\n", os.Args[0])
		return
	}
	// Maps are supported output file formats and values are names of files to write to if any.
	outputFormats := map[string]string{}
	for format, _ := range SupportedFormatHandlers {
		outputFormats[format] = ""
	}
	programFile := os.Args[len(os.Args)-1]
	// Parse arguments in any form such as "--json output.json -YAML data.yaml xml encoded.xml myprogram.fig"
	for i := 1; i < len(os.Args)-1; i++ {
		format := strings.ToLower(strings.Replace(os.Args[i], "-", "", -1))
		_, isSupported := outputFormats[format]
		if isSupported {
			outputFormats[format] = os.Args[i+1]
			i++
		}
	}
	// Open and interpret the program file
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
	env, err := Interpret(program)
	if err != nil {
		fmt.Println("ERROR\n  ", err.Error())
	}
	// Produce the desired output files
	err = WriteOutputFiles(outputFormats, env)
	if err != nil {
		fmt.Println("ERROR\n  ", err.Error())
	}
}
