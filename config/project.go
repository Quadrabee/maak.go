package config

type Project struct {
	name       string
	rootPath   string
	components map[string]*Component
}

func CreateProject(name string, rootPath string) *Project {
	p := new(Project)
	p.name = name
	p.rootPath = rootPath
	p.components = make(map[string]*Component)
	return p
}

func (proj *Project) CreateComponent(name string) *Component {
	comp := new(Component)
	comp.name = name
	comp.project = proj
	proj.components[name] = comp
	return comp
}

func (proj *Project) Components() []*Component {
	var comps []*Component
	for _, comp := range proj.components {
		comps = append(comps, comp)
	}
	return comps
}

func (proj *Project) Containers() []*Container {
	var containers []*Container
	for _, comp := range proj.Components() {
		containers = append(containers, comp.Containers()...)
	}
	return containers
}

func (proj *Project) DockerPrefix() string {
	return proj.name + "/"
}

func (proj *Project) Name() string {
	return proj.name
}

func (proj *Project) RootPath() string {
	return proj.rootPath
}
