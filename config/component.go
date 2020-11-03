package config

type Component struct {
	project    *Project
	name       string
	containers []*Container
}

func (comp *Component) AddContainer(name *string, dockerfile *string, context *string) *Container {
	cont := new(Container)

	cont.component = comp
	cont.name = name
	cont.dockerfile = dockerfile
	cont.context = context
	comp.containers = append(comp.containers, cont)
	return cont
}

func (comp *Component) Containers() []*Container {
	return comp.containers
}

func (comp *Component) Name() string {
	return comp.name
}
