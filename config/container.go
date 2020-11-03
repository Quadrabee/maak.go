package config

import (
	"path"

	"gopkg.in/godo.v2/glob"
)

type Deps struct {
	glob       *[]string
	components *[]string
}

type Container struct {
	component    *Component
	name         *string
	dockerfile   *string
	context      *string
	dependencies *Deps
}

func (cont *Container) Name() string {
	if cont.name != nil {
		return cont.component.Name() + "." + *cont.name
	}
	return cont.component.Name()
}

func (cont *Container) ImageName() string {
	var prefix = cont.component.project.DockerPrefix()
	// Single container for a component
	if cont.name == nil {
		return prefix + cont.component.Name()
	}
	return prefix + cont.component.Name() + "." + *cont.name
}

func (cont *Container) Context() string {
	if cont.context == nil {
		return cont.component.Name()
	}
	return *cont.context
}

func (cont *Container) Dockerfile() string {
	if cont.dockerfile != nil {
		return path.Join(cont.Context(), *cont.dockerfile)
	} else if cont.name != nil {
		return path.Join(cont.Context(), "Dockerfile."+*cont.name)
	} else {
		return path.Join(cont.Context(), "Dockerfile")
	}
}

func (cont *Container) SetDependencies(glob *[]string, components *[]string) {
	cont.dependencies = new(Deps)
	cont.dependencies.glob = glob
	cont.dependencies.components = components
}

func (cont *Container) ComponentDependencies() []string {
	if cont.dependencies == nil {
		return make([]string, 0)
	}
	if cont.dependencies.components == nil {
		return make([]string, 0)
	}
	return *cont.dependencies.components
}

func (cont *Container) GlobDependencies() ([]string, error) {
	if cont.dependencies == nil {
		cont.dependencies = new(Deps)
	}
	if cont.dependencies.glob == nil {
		cont.dependencies.glob = &[]string{"**/*"}
	}

	// Glob dependency: list the files
	var globs []string
	for _, glob := range *cont.dependencies.glob {
		globs = append(globs, path.Join(cont.Context(), glob))
	}

	assets, _, err := glob.Glob(globs)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, asset := range assets {
		files = append(files, asset.Path)
	}
	return files, nil
}
