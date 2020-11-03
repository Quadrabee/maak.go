package config

import (
	"errors"
	"fmt"
	"os"
)

var p *Project

func Load() (*Project, error) {
	var err error
	p, err = generateProject()
	return p, err
}

func EnsureLoaded() {
	if p == nil {
		var err error
		_, err = Load()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func Loaded() bool {
	return p != nil
}

func EnsureNotExists() error {
	if _, err := os.Stat("maak.yaml"); err == nil {
		return errors.New("maak.yaml already exists")
	}
	return nil
}

func GetComponent(name string) *Component {
	for _, comp := range p.Components() {
		if comp.Name() == name {
			return comp
		}
	}
	return nil
}

func ComponentNames() []string {
	result := make([]string, 0, len(p.Components()))
	for _, comp := range p.Components() {
		result = append(result, comp.Name())
	}
	return result
}

func GetContainer(name string) *Container {
	for _, comp := range p.Components() {
		for _, cont := range comp.Containers() {
			if cont.Name() == name {
				return cont
			}
		}
	}
	return nil
}

func ContainerNames() []string {
	var names []string
	for _, comp := range p.Components() {
		for _, cont := range comp.Containers() {
			names = append(names, cont.Name())
		}
	}
	return names
}
