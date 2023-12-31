package handler

import (
	"suchicodes-go/cmd/cookie"
	"suchicodes-go/view/userViews"

	"github.com/labstack/echo/v4"
)


type UserHandler struct {
}

func (h *UserHandler) HandleUserShow(c echo.Context) error {
    return render(c, userViews.UserPage())
}

func (h *UserHandler) HandleUserLogin(c echo.Context) error {
    cookie.SetUserSession("suchi", c)    
    return c.HTML(200, "bye?")
}
