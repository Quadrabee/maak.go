package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/quadrabee/maak/utils"
	"gopkg.in/yaml.v2"
)

type YamlDeps struct {
	Glob       *[]string
	Components *[]string
}

type YamlContainer struct {
	Name         *string
	Dockerfile   *string
	Context      *string
	Dependencies *YamlDeps
}

type YamlComponent struct {
	Name       *string
	Containers []YamlContainer
}

type YamlProject struct {
	Name *string `yaml:"name,omitempty"`
}

type YamlConfig struct {
	RootPath   string `yaml:"-"`
	Project    YamlProject
	Components map[string]YamlComponent
}

var loaded = false
var yamlConfig YamlConfig

func loadYamlConfig() (*YamlConfig, error) {
	if loaded {
		return &yamlConfig, nil
	}
	path, err := utils.Find("maak.yaml")
	if err != nil {
		return nil, errors.New("No maak.yaml found")
	}

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Readfile: %v\n", err)
		return nil, errors.New("maak.yaml not readable")
	}

	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		fmt.Printf("Unmarshal: %v\n", err)
		return nil, errors.New("Invalid maak.yaml")
	}

	// Ensure the project root path is set
	yamlConfig.RootPath = filepath.Dir(path)

	loaded = true
	return &yamlConfig, nil
}
