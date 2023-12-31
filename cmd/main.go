package main

import (
	"suchicodes-go/config"
	"suchicodes-go/handler"

	"github.com/gorilla/sessions"
	echo_session "github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {

    config := config.GetConfig()

    app := echo.New()

    app.Use(middleware.Gzip())
    app.Use(echo_session.Middleware(sessions.NewCookieStore(
	[]byte(config.SECRET_KEY),
    )))

    app.Static("/static", "static")

    homeHandler := handler.HomeHandler{
	Config: config,
    }
    app.GET("/", homeHandler.HandleIndexShow)

    userHandler := handler.UserHandler{
	Config: config,
    }
    app.GET("/user", userHandler.HandleUserShow)

    app.Start(":3000")
}
