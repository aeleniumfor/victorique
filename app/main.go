package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/victorique/app/handler"
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


	
	s := echo.New()
	s.Use(middleware.Logger())
	s.GET("/container/list/id", handler.HndlerGetContainerIDList())
	s.GET("/container/list/name", handler.HndlerGetContainerNameList())
	s.Start(":1323")
	// cli.CreateContainer(co)
	// cli.StartContainer(co)
	// statestore.SetKey()
}
