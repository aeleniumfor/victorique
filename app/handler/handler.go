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

// GetContainerIDList is GetContainerIDList 
func GetContainerIDList() echo.HandlerFunc {
    return func(c echo.Context) error {
		cli := container.New(ContainerRemote)
		idlist := new(common.ResCotainerIDList)
		idlist.ContainerIDList = cli.GetContainerIDList()
		return c.JSON(200,idlist)
    }
}

// GetContainerNameList is GetContainerIDList 
func GetContainerNameList() echo.HandlerFunc {
    return func(c echo.Context) error {
		cli := container.New(ContainerRemote)
		namelist := new(common.ResCotainerNameList)
		namelist.ContainerNameList = cli.GetContainerNameList()
		return c.JSON(200,namelist)
    }
}

// ContainerInfoFromID is ContainerID inspect handler
func ContainerInfoFromID() echo.HandlerFunc {
    return func(c echo.Context) error {
		cli := container.New(ContainerRemote)
		containerJSON :=new(common.Container)

		containerID := c.Param("id")
		containerInfo := cli.GetContainerInfoFromID(containerID)
		containerJSON.ID = containerInfo.ID
		containerJSON.Name = containerInfo.Name
		containerJSON.NetNamespace = containerInfo.NetNamespace

		return c.JSON(200,containerJSON)
    }
}

// ContainerInfoFromName is ContainerID inspect handler
func ContainerInfoFromName() echo.HandlerFunc {
    return func(c echo.Context) error {
		cli := container.New(ContainerRemote)
		containerJSON :=new(common.Container)
		containerName := c.Param("name")
		containerInfo := cli.GetContainerInfoFromName(containerName)
		containerJSON.ID = containerInfo.ID
		containerJSON.Name = containerInfo.Name
		containerJSON.NetNamespace = containerInfo.NetNamespace

		return c.JSON(200,containerJSON)
    }
}