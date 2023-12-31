package main

import (
	"suchicodes-go/cmd/cookie"
	"suchicodes-go/config"
	"suchicodes-go/handler"
	suchi_middleware "suchicodes-go/middleware"

	"github.com/gorilla/sessions"
	echo_session "github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	echo_middleware "github.com/labstack/echo/v4/middleware"
)


func main() {

    config.InitConfig()

    app := echo.New()

    app.Use(echo_session.Middleware(sessions.NewCookieStore(
	[]byte(config.SECRET_KEY),
    )))
    app.Use(echo_middleware.Gzip())
    app.Use(suchi_middleware.ContextMiddleware())

    app.Static("/static", "static")

    homeHandler := handler.HomeHandler{}
    app.GET("/", homeHandler.HandleIndexShow)

    userHandler := handler.UserHandler{}
    app.GET("/user", userHandler.HandleUserShow)
    app.GET("/login", userHandler.HandleUserLogin)
    app.GET("/dark-theme", func (c echo.Context) error {
	cookie.SetThemeCookie("dark", c)
	return c.HTML(200, "set dark theme")
    })

    app.Start(":3000")
}
