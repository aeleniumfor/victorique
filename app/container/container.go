package container

import (
	"fmt"
	"net/http"

	"github.com/docker/docker/client"
)

// Container type struct
type Container struct {
	id     string
	name   string
	ip     string
	status string
}

// Containers type struct
type Containers struct {
	host      string
	container Container
}

// New Container struct
func New() (c *Container) {
	return &Container{}
}

// Connect docker server
func Connect() (cli *client.Client) {
	tr := &http.Transport{}
	cli, err := client.NewClient("http://localhost:2375", client.DefaultVersion, &http.Client{Transport: tr}, map[string]string{})
	if err != nil {
		panic(err)
	}
	return cli
}

func (c *Container) GetContainer() {
	connectList := 1
	fmt.Println(connectList)
}
