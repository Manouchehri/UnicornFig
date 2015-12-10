package codegen

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

const CodeTemplate = `
package config

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"os"
	"io/ioutil"
)

// A structure that contains parsed configuration data
type Configuration struct {
	{{.Fields}}
}

func LoadConfigJson(fileName string) (Configuration, error) {
	config := Configuration{}
	file, openErr := os.Open(fileName)
	if openErr != nil {
		return config, openErr
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	decodeErr := decoder.Decode(&config)
	return config, decodeErr
}

func LoadConfigYaml(fileName string) (Configuration, error) {
	config := Configuration{}
	file, openErr := os.Open(fileName)
	if openErr != nil {
		return config, openErr
	}
	defer file.Close()
	bytes, readErr := ioutil.ReadAll(file)
	if readErr != nil {
		return config, readErr
	}
	decodeErr := yaml.Unmarshal(bytes, &config)
	return config, decodeErr
}
`

const FieldTemplate = "{{.FieldName}} {{.Type}} {{.Tags}}"

/**
 * Strip out symbols and capitalize the first character of a string to make it
 * a Go-specific public field name.
 */
func fieldName(varName string) string {
	if len(varName) == 0 {
		return ""
	}
	newName := strings.ToUpper(string(varName[0]))
	for i := 1; i < len(varName); i++ {
		char := varName[i]
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			newName += string(char)
		}
	}
	return newName
}

/**
 * Create the fields strings to be inserted into the code template.
 * These fields define the struct into which config file data can be parsed/unmarshalled.
 */
func createFields(env map[string]interface{}) string {
	fields := make([]string, len(env))
	index := 0
	for k, v := range env {
		field := strings.Replace(FieldTemplate, "{{.FieldName}}", fieldName(k), 1)
		tags := fmt.Sprintf("`json:\"%s\", yaml:\"%s\"`", k, k)
		field = strings.Replace(field, "{{.Tags}}", tags, 1)
		typeName := ""
		switch v.(type) {
		case string:
			typeName = "string"
		case int64:
			typeName = "int64"
		case float64:
			typeName = "float64"
		case bool:
			typeName = "bool"
		case []interface{}:
			typeName = "[]interface{}"
		case map[string]interface{}:
			typeName = "map[string]interface{}"
		}
		field = strings.Replace(field, "{{.Type}}", typeName, 1)
		fields[index] = "\t" + field
		index++
	}
	return strings.Join(fields, "\n")
}

func GenerateConfigCodeFile(env map[string]interface{}, fileName string) error {
	file, openErr := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if openErr != nil {
		return openErr
	}
	defer file.Close()
	t, templateErr := template.New("code").Parse(CodeTemplate)
	if templateErr != nil {
		return templateErr
	}
	fields := createFields(env)
	return t.Execute(file, map[string]string{"Fields": fields})
}
