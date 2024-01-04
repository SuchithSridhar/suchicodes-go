package middleware

import (
	"log/slog"
	"net/http"
	"strings"
	"suchicodes-go/cmd/cookie"
	"suchicodes-go/config"

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

func SetCachePolicy() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            var cacheOption string
            if (config.Config.ENVIRONMENT == "development") {
                cacheOption = "no-cache"
            } else {
                cacheOption = "max-age=1200, must-revalidate"
            }
            c.Response().Header().Set("Cache-Control", cacheOption)
            return next(c)
        }
    }
}

func getIPAddress(r *http.Request) string {
    ipHeaders := []string{"X-Forwarded-For", "X-Real-IP", "X-Client-IP", "CF-Connecting-IP"}
    for _, header := range ipHeaders {
        ip := r.Header.Get(header)
        if ip != "" && ip != "unknown" {
            ip = strings.TrimSpace(strings.Split(ip, ",")[0])
            return ip
        }
    }
    ip := r.RemoteAddr
    return ip
}

func AccessLog() echo.MiddlewareFunc {
    return func (next echo.HandlerFunc) echo.HandlerFunc {
        return func (c echo.Context) error {
            var request *http.Request = c.Request()
            url := request.URL.String()
            method := request.Method
            ipAddr := getIPAddress(request)

            slog.Info("Access Log", "url", url, "method", method, "ip-addr", ipAddr)
            return next(c)
        }
    }
}
