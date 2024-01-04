package main

import (
	"log/slog"
	"os"
	"suchicodes-go/config"
	"suchicodes-go/handler"
	suchi_middleware "suchicodes-go/middleware"

	"github.com/gorilla/sessions"
	echo_session "github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	echo_middleware "github.com/labstack/echo/v4/middleware"
)

func configureLogger() {
    logLevel := new(slog.LevelVar)
    if config.Config.ENVIRONMENT == "development" {
	logLevel.Set(slog.LevelDebug)
    } else {
	logLevel.Set(slog.LevelInfo)
    }
    logger := slog.NewTextHandler(
	os.Stdout,
	&slog.HandlerOptions{AddSource: true, Level: logLevel},
    )
    slog.SetDefault(slog.New(logger))
}


func main() {

    config.InitConfig()

    configureLogger()

    app := echo.New()

    // Middleware installs
    app.Use(echo_session.Middleware(sessions.NewCookieStore(
	[]byte(config.SECRET_KEY),
    )))
    app.Use(echo_middleware.Gzip())
    app.Use(suchi_middleware.ContextMiddleware())
    app.Use(suchi_middleware.SetCachePolicy())
    app.Use(suchi_middleware.AccessLog())

    // General app settings
    app.Static("/static", "static")

    // Home page routing
    homeHandler := handler.HomeHandler{}
    app.GET("/", homeHandler.HandleIndexShow)
    app.GET("/about", homeHandler.HandleAboutShow)
    app.GET("/contact", homeHandler.HandleContactShow)
    app.File("/resume", "data/resume.pdf")
    
    // User page routing
    userHandler := handler.UserHandler{}
    app.GET("/user", userHandler.HandleUserShow)
    app.GET("/login", userHandler.HandleUserLogin)

    app.Start(":3000")
}
