package handler

import (
	"suchicodes-go/cmd/auth"
	"suchicodes-go/config"
	"suchicodes-go/view/home"

	"github.com/labstack/echo/v4"
)


type HomeHandler struct {
    Config *config.GlobalConfig
}

func (h *HomeHandler) HandleIndexShow(c echo.Context) error {
    username, ok := auth.IsLoggedIn(c)
    return render(c, home.Show())
}
