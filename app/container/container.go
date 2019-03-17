package container

import (
	"fmt"
	"context"
	"log"
	"net/http"

	"github.com/docker/docker/api/types/container"

	"github.com/docker/docker/api/types"

	"github.com/docker/docker/client"
	"github.com/victorique/app/common"
)

// DockerClient is docker client
type DockerClient struct {
	cli *client.Client
}

// New is manager
func New(hostname string) (dockerCli *DockerClient) {
	tr := &http.Transport{}
	cli, err := client.NewClient(hostname, client.DefaultVersion, &http.Client{Transport: tr}, map[string]string{})
	if err != nil {
		log.Println(err)
	}
	// &common.Containers{Host: hostname}
	return &DockerClient{cli: cli}
}

// GetContainerList is multi host get container list
func (dockerCli *DockerClient) GetContainerList(c *common.Containers) {
	list, _ := dockerCli.cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	for i := 0; i < len(list); i++ {
		container := common.Container{
			Name: list[i].Names[0],
			ID:   list[i].ID,
			IP:   list[i].NetworkSettings.Networks["bridge"].IPAddress,
		}
		c.ContainerStructs = append(c.ContainerStructs, container)
	}
}

// CreateContainer is Greate Container
func (dockerCli *DockerClient) CreateContainer() {
	// containerを作るだけの機能
	config := &container.Config{
		Image: "alpine",
		Cmd:   []string{"echo", "hello world"},
	}
	con, _ := dockerCli.cli.ContainerCreate(context.Background(), config, nil, nil, "ProjectName-ServiceName-ContainerName")
	fmt.Println(con)
}

// StartContainer is start Created Container
func (dockerCli *DockerClient) StartContainer() {
	
}

// GetContainerNameList is Get Containers Name
// func (c *Containers) GetContainerNameList() []string {
// 	containerNameList := []string{}
// 	for i := 0; i < len(c.Container); i++ {
// 		containerNameList = append(containerNameList, c.Container[i].Name)
// 	}
// 	return containerNameList
// }
