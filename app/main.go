package main

import (
	"fmt"

	"github.com/victorique/app/common"
	"github.com/victorique/app/container"
)

func main() {

	// tr := &http.Transport{}
	// cli, err := client.NewClient("http://localhost:2375", client.DefaultVersion, &http.Client{Transport: tr}, map[string]string{})
	// if err != nil {
	// 	panic(err)
	// }

	// containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})

	// for i, container := range containers {
	// 	fmt.Println(container.Image)
	// 	fmt.Println(i)
	// }

	containers := &common.Containers{}
	co := &common.Container{}
	cli := container.New("http://localhost:4243")
	cli.GetContainerList(containers)
	cli.CreateContainer(co)
	cli.StartContainer(co)
	fmt.Println(co)
}
