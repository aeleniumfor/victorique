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
	s.GET("/container/list/id", handler.GetContainerIDList())
	s.GET("/container/list/name", handler.GetContainerNameList())
	s.GET("/container/id/:id", handler.ContainerInfoFromID())
	s.GET("/container/name/:name", handler.ContainerInfoFromName())

	s.POST("/container/create/:name", handler.CreateContainer())
	s.POST("/container/start/:id", handler.StartContainer())
	s.POST("/container/stop/:id",handler.StopContainer())

	s.DELETE("/container/delete/:id",handler.DeleteContainer())
	
	s.Start(":8080")
	// cli.CreateContainer(co)
	// cli.StartContainer(co)
	// statestore.SetKey()
}
