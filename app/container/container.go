package container

type Container struct {
	id     string
	name   string
	ip     string
	status string
}

func New() (c *Container) {
	return &Container{}
}
