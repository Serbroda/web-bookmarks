package handlers

import (
	"fmt"

	"github.com/Serbroda/ragbag/gen"
	"github.com/labstack/echo/v4"
)

type ServerInterfaceImpl struct {
}

func RegisterHandlers(router gen.EchoRouter, si gen.ServerInterface) {
	wrapper := gen.ServerInterfaceWrapper{
		Handler: si,
	}
	router.GET("/groups", wrapper.GetGroups)
}

func (si *ServerInterfaceImpl) GetGroups(ctx echo.Context, params gen.GetGroupsParams) error {
	fmt.Println("GetGroups")
	return nil
}
