package handler

import (
	"github.com/victorique/app/common"
	"github.com/victorique/app/container"
	"github.com/labstack/echo"
)



// HndlerGetContainerIDList is GetContainerIDList 
func HndlerGetContainerIDList() echo.HandlerFunc {
    return func(c echo.Context) error {
		cli := container.New("http://localhost:2376")
		idlist := new(common.ResCotainerIDList)
		idlist.ContainerID = cli.GetContainerIDList()
		return c.JSON(200,idlist)
    }
}