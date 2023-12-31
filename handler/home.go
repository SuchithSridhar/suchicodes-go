package handler

import (
	"suchicodes-go/view/home"

	"github.com/labstack/echo/v4"
)


type HomeHandler struct {
}

func (h *HomeHandler) HandleIndexShow(c echo.Context) error {
    return render(c, home.Show())
}
