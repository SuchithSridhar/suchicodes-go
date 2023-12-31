package handler

import (
	"suchicodes-go/config"
	"suchicodes-go/view/user"

	"github.com/labstack/echo/v4"
)


type UserHandler struct {
    Config *config.GlobalConfig
}

func (h *UserHandler) HandleUserShow(c echo.Context) error {
    return render(c, user.Show())
}
