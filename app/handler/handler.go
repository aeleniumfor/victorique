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

// HndlerGetContainerIDList is GetContainerIDList 
func HndlerGetContainerIDList() echo.HandlerFunc {
    return func(c echo.Context) error {
		cli := container.New(ContainerRemote)
		idlist := new(common.ResCotainerIDList)
		idlist.ContainerID = cli.GetContainerIDList()
		return c.JSON(200,idlist)
    }
}

// HndlerGetContainerNameList is GetContainerIDList 
func HndlerGetContainerNameList() echo.HandlerFunc {
    return func(c echo.Context) error {
		cli := container.New(ContainerRemote)
		namelist := new(common.ResCotainerNameList)
		namelist.ContainerName = cli.GetContainerNameList()
		return c.JSON(200,namelist)
    }
}