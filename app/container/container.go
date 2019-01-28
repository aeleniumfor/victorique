package container

import (
	"net/http"

	"github.com/docker/docker/client"
)

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

func (c *Container) GetContainer() {
	tr := &http.Transport{}
	cli, err := client.NewClient("http://localhost:2375", client.DefaultVersion, &http.Client{Transport: tr}, map[string]string{})
	if err != nil {
		panic(err)
	}
}
