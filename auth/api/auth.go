package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/suksest/commodity/handler"
)

// AuthGroup api group
func AuthGroup(g *echo.Group) {
	g.Use(middleware.Logger())

	g.POST("/login", handler.Login)
	g.POST("/signup", handler.Signup)
	g.GET("/check", handler.Check)
}
