package config

import (
	"path/filepath"
)

func generateContainers(component *Component, conf []YamlContainer) error {
	if len(conf) <= 0 {
		component.AddContainer(nil, nil, nil)
	}
	for _, cont := range conf {
		container := component.AddContainer(cont.Name, cont.Dockerfile, cont.Context)
		if cont.Dependencies != nil {
			container.SetDependencies(cont.Dependencies.Glob, cont.Dependencies.Components)
		}
	}
	return nil
}

func generateComponent(project *Project, name string, conf YamlComponent) error {
	comp := project.CreateComponent(name)
	generateContainers(comp, conf.Containers)
	return nil
}

func generateProject() (*Project, error) {
	config, err := loadYamlConfig()
	if err != nil {
		return nil, err
	}
	if config.Project.Name == nil {
		var name = filepath.Base(config.RootPath)
		config.Project.Name = &name
	}

	project := CreateProject(*config.Project.Name, config.RootPath)

	for name, def := range config.Components {
		generateComponent(project, name, def)
	}
	return project, nil
}
