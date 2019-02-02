package container

import (
	"context"
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// Container type struct
type Container struct {
	name    string
	id      string
	inspect types.ContainerJSON
	cli     *client.Client
}

// New Container struct
func New(hostname string) (c *Container) {
	tr := &http.Transport{}
	hostname = "http://" + hostname + ":2375"
	fmt.Println(hostname)
	cli, err := client.NewClient(hostname, client.DefaultVersion, &http.Client{Transport: tr}, map[string]string{})
	if err != nil {
		panic(err)
	}
	return &Container{cli: cli}
}

// GetContainerNameList method is docker container name list
func (c *Container) GetContainerNameList() []string {
	containerlist := []string{}
	containers, _ := c.cli.ContainerList(context.Background(), types.ContainerListOptions{})
	for _, container := range containers {
		containerlist = append(containerlist, container.Names[0])
	}
	return containerlist
}

// GetContainerInspect is Container inspect
func (c *Container) GetContainerInspect() {
	i, _ := c.cli.ContainerInspect(context.Background(), c.id)
	c.inspect = i
}

// CreateContainer is Container create return Container ID
func (c *Container) CreateContainer(containerName string) (ConrainerID string) {
	container, _ := c.cli.ContainerCreate(context.Background(), &container.Config{
		Image: "alpine",
		Tty:   true,
	}, &container.HostConfig{
		AutoRemove: true,
	}, nil, containerName)
	c.name = containerName
	c.id = container.ID
	return c.id
}

// StopContainer is docker stop <id>
func (c *Container) StopContainer() {
	fmt.Println(c.id)
	c.cli.ContainerStop(context.Background(), c.id, nil)
}

// RunContainer is docker run
func (c *Container) RunContainer() {
	c.cli.ContainerStart(context.Background(), c.id, types.ContainerStartOptions{})
}

// ListContainer is docker ps
func (c *Container) ListContainer() (NameList []string) {
	list, _ := c.cli.ContainerList(context.Background(), types.ContainerListOptions{})
	ContainerNameList := []string{}
	for _, container := range list {
		ContainerNameList = append(ContainerNameList, container.Names[0])
	}
	return ContainerNameList
}

//テスト環境が動かないためコメントで削除したことにしておく
// func main() {
// 	s := New()
// 	s.CreateContainer()
// 	s.RunContainer()
// 	s.GetContainerInspect()
// 	s.StopContainer()
// }
