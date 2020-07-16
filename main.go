package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/suksest/commodity/api"
	"github.com/suksest/commodity/handler"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	api.AuthGroup(e.Group("/auth"))
	// api.CommoGroup(e.Group("/commo")) //not working yet
	e.GET("/version", handler.GetVersion)

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("SERVICE_PORT")))
}
