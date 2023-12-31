package auth

import (
	"log/slog"
	"suchicodes-go/config"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const yearInSeconds = 365 * 24 * 60 * 60

func IsLoggedIn(c echo.Context) (ok bool, username string) {
    sesh, err := session.Get("session", c)
    if err != nil {
	return false, ""
    }

    user, ok := sesh.Values["user"]
    
    if !ok {
	return false, ""
    }

    if userStr, ok := user.(string); ok {
	return true, userStr
    } else {
	return false, ""
    }
}

func Login(username string, cfg config.GlobalConfig, c echo.Context) (ok bool) {
    sesh, err := session.Get("session", c)
    if err != nil {
	slog.Error("Unable to get session during login.")
	return false
    }

    sesh.Options = &sessions.Options {
	Path: "/",
	MaxAge: yearInSeconds,
	HttpOnly: true,
	Secure: cfg.ENVIRONMENT != "development",  // only when env is not dev
    }

    sesh.Values["user"] = username
    sesh.Save(c.Request(), c.Response())

    return true
}

func Logout(c echo.Context) (ok bool) {
    sesh, err := session.Get("session", c)
    if err != nil {
	slog.Error("Unable to get session during logout.")
	return false
    }

    delete(sesh.Values, "user")
    sesh.Save(c.Request(), c.Response())

    return true
}
