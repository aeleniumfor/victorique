package container

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/google/uuid"
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

// GetContainerIDList is
func (dockerCli *DockerClient) GetContainerIDList() []string {
	list, _ := dockerCli.cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	containerlist := []string{}
	for i := 0; i < len(list); i++ {
		containerlist = append(containerlist, list[i].ID)
	}
	return containerlist
}

// GetContainerNameList is
func (dockerCli *DockerClient) GetContainerNameList() []string {
	list, _ := dockerCli.cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	containerlist := []string{}
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i].Names)
		containerlist = append(containerlist, list[i].Names[0])
	}
	return containerlist
}

// GetContainerInfoFromID is
func (dockerCli *DockerClient) GetContainerInfoFromID(containerID string) *common.Container {
	inspect, _ := dockerCli.cli.ContainerInspect(context.Background(), containerID)
	c := new(common.Container)
	c.ID = inspect.ID
	c.Name = inspect.Name
	c.NetNamespace = inspect.NetworkSettings.SandboxKey
	c.Status = inspect.State.Status
	return c
}

// GetContainerInfoFromName is
func (dockerCli *DockerClient) GetContainerInfoFromName(containerName string) *common.Container {
	c := new(common.Container)
	return c
}

// CreateContainer is Greate Container
func (dockerCli *DockerClient) CreateContainer(containerName string) string {
	// containerを作るだけの機能
	config := &container.Config{
		Image: "alpine",
		Cmd:   []string{"echo", "hello world"},
	}
	id := uuid.New().String()
	name := containerName + "-" + id
	createdContainer, err := dockerCli.cli.ContainerCreate(context.Background(), config, nil, nil, name)
	if err != nil {
		log.Println(err)
		return ""
	}
	return createdContainer.ID
}

// StartContainer is start Created Container
func (dockerCli *DockerClient) StartContainer(containerID string) *common.Container {
	dockerCli.cli.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
	inspect, _ := dockerCli.cli.ContainerInspect(context.Background(), containerID)
	
	c := new(common.Container)
	c.ID = inspect.ID
	c.Name = inspect.Name
	c.NetNamespace = inspect.NetworkSettings.SandboxKey
	c.Status = inspect.State.Status
	return c
}

// GetContainerNameList is Get Containers Name
// func (c *Containers) GetContainerNameList() []string {
// 	containerNameList := []string{}
// 	for i := 0; i < len(c.Container); i++ {
// 		containerNameList = append(containerNameList, c.Container[i].Name)
// 	}
// 	return containerNameList
// }

// docker run \
//   -d \
//   -e ENV_DOCKER_REGISTRY_HOST=127.0.0.1 \
//   -e ENV_DOCKER_REGISTRY_PORT=5000 \
//   -p 8080:80 \
//   konradkleine/docker-registry-frontend:v2
