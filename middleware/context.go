package middleware

import (
	"suchicodes-go/cmd/cookie"

	"github.com/labstack/echo/v4"
)

func ContextMiddleware() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {

            theme := cookie.ReadThemeCookie(c)
            c.Set(cookie.ThemeKey, theme)

            username, _ := cookie.ReadUserSession(c)
            c.Set(cookie.UserKey, username)

            return next(c)
        }
    }
}
