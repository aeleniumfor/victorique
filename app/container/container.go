package container

import (
	"context"
	"log"
	"net/http"

	"github.com/docker/docker/api/types"

	"github.com/docker/docker/client"
	"github.com/victorique/app/common"
)

// DockerClient is docker client
type DockerClient struct {
	cli *client.Client
}

// Container is に代入するすべを知らないので辛い
type Container struct {
	Name string
	ID   string
	IP   string
}


// New is manager
func New(hostname string) (dockerClient *DockerClient) {
	tr := &http.Transport{}
	cli, err := client.NewClient(hostname, client.DefaultVersion, &http.Client{Transport: tr}, map[string]string{})
	if err != nil {
		log.Println(err)
	}
	// &common.Containers{Host: hostname}
	return &DockerClient{cli: cli}
}

// GetContainerList is multi host get container list
func (dockerClient *DockerClient) GetContainerList(c *common.Containers) {
	list, _ := dockerClient.cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	for i := 0; i < len(list); i++ {
		con := Container{
			Name: list[i].Names[0],
			ID: list[i].ID,
			IP: list[i].NetworkSettings.Networks["bridge"].IPAddress, // 中身はmap
		}
		c.ContainerStructs = append(c.ContainerStructs,con)
	}
}
// GetContainerNameList is Get Containers Name
// func (c *Containers) GetContainerNameList() []string {
// 	containerNameList := []string{}
// 	for i := 0; i < len(c.Container); i++ {
// 		containerNameList = append(containerNameList, c.Container[i].Name)
// 	}
// 	return containerNameList
// }
