package utils

import (
    "context"
    "suchicodes-go/cmd/cookie"
)

func IsLoggedIn(c context.Context) string {
    if username, ok := c.Value(cookie.UserKey).(string); ok {
        return username
    } else {
        return ""
    }
}

func GetCurrentTheme(c context.Context) string {
    if theme, ok := c.Value(cookie.ThemeKey).(string); ok {
        return theme
    } else {
        return "light"
    }
}

