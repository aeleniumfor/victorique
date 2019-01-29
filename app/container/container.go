package main

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

// GetContainer method is docker ps
func (c *Container) GetContainer() {
	connect := make(chan *client.Client)
	connectList := []func(){
		func() {
			tr := &http.Transport{}
			cli, err := client.NewClient("http://localhost:2375", client.DefaultVersion, &http.Client{Transport: tr}, map[string]string{})
			if err != nil {
				panic(err)
			}
			connect <- cli
		},
	}
	for _, i := range connectList {
		go i()
	}
	for i := 0; i < len(connectList); i++ {
		fmt.Println(<-connect)
	}
}

func main() {
	s := New()
	s.GetContainer()
}
