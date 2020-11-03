package config

import (
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

func Generate() error {
	err := EnsureNotExists()
	if err != nil {
		return err
	}

	var cfg YamlConfig

	dir, _ := os.Getwd()
	name := path.Base(dir)
	cfg.Project.Name = &name
	str, _ := yaml.Marshal(&cfg)

	err = ioutil.WriteFile("maak.yaml", str, 0644)
	if err != nil {
		return err
	}
	return nil
}
