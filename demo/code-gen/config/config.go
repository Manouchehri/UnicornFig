
package config

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"os"
	"io/ioutil"
)

// A structure that contains parsed configuration data
type Configuration struct {

	Routes map[string]interface{} `json:"routes", yaml:"routes"`

	Port string `json:"port", yaml:"port"`

	Address string `json:"address", yaml:"address"`

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
