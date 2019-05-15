package handler

import (
	"github.com/victorique/app/common"
	"github.com/victorique/app/container"
	"github.com/labstack/echo"
)

const(
	// ContainerRemote is docker remote api url
	ContainerRemote string = "http://localhost:2376"
)

// HandlerGetContainerIDList is GetContainerIDList 
func HandlerGetContainerIDList() echo.HandlerFunc {
    return func(c echo.Context) error {
		cli := container.New(ContainerRemote)
		idlist := new(common.ResCotainerIDList)
		idlist.ContainerID = cli.GetContainerIDList()
		return c.JSON(200,idlist)
    }
}

// HandlerGetContainerNameList is GetContainerIDList 
func HandlerGetContainerNameList() echo.HandlerFunc {
    return func(c echo.Context) error {
		cli := container.New(ContainerRemote)
		namelist := new(common.ResCotainerNameList)
		namelist.ContainerName = cli.GetContainerNameList()
		return c.JSON(200,namelist)
    }
}

// HandlerContainerID is ContainerID inspect handler
func HandlerContainerID() echo.HandlerFunc {
    return func(c echo.Context) error {
		cli := container.New(ContainerRemote)
		namelist := new(common.ResCotainerNameList)
		namelist.ContainerName = cli.
		return c.JSON(200,namelist)
    }
}