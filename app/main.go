package main

import (
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
	var c = container.New("hoda.cong")
	c.GetContainerList(containers)
}
