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

		containerID := c.Param("id")
		containerJSON :=new(common.Container)		
		containerJSON = cli.GetContainerInfoFromID(containerID)
		return c.JSON(200,containerJSON)
    }
}

// ContainerInfoFromName is ContainerID inspect handler
func ContainerInfoFromName() echo.HandlerFunc {
    return func(c echo.Context) error {
		cli := container.New(ContainerRemote)
		
		containerName := c.Param("name")
		containerJSON :=new(common.Container)
		containerJSON = cli.GetContainerInfoFromName(containerName)
		return c.JSON(200,containerJSON)
    }
}

// CreateContainer is Create container handler
func CreateContainer() echo.HandlerFunc {
	return func(c echo.Context) error {
		containerName := c.Param("name")
		cli := container.New(ContainerRemote)
		
		containerID := cli.CreateContainer(containerName) 
		containerJSON :=new(common.Container)
		containerJSON = cli.GetContainerInfoFromID(containerID)
		return c.JSON(200,containerJSON)
	}
}

// StartContainer is Create container handler
func StartContainer() echo.HandlerFunc {
	return func(c echo.Context) error {
		cli := container.New(ContainerRemote)

		containerID := c.Param("id")
		containerJSON :=new(common.Container)
		containerJSON = cli.StartContainer(containerID)
		return c.JSON(200,containerJSON)
	}
}