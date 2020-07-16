package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/suksest/commodity/handler"
)

// CommoGroup api group for fetch
func CommoGroup(g *echo.Group) {
	g.Use(middleware.Logger())

	g.GET("/fetch", handler.Fetch)
}
