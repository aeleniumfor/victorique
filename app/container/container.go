package container

type Container struct {
	id     string
	name   string
	ip     string
	status string
}

type Containers struct {
	host      string
	container Container
}

func New() (c *Container) {
	return &Container{}
}
