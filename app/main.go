package main

import (
	"github.com/labstack/echo"
	"fmt"
	"github.com/victorique/app/container"
	"github.com/labstack/echo/middleware"
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


	cli := container.New("http://localhost:2376")
	fmt.Println(cli.GetContainerIDList())
	s := echo.New()
	s.Use(middleware.Logger())
    s.Use(middleware.Recover())
    s.GET("/hello", handler.MainPage())
    s.Start(":1323")
	// cli.CreateContainer(co)
	// cli.StartContainer(co)
	// statestore.SetKey()
}
