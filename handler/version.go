package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/suksest/commodity/model"
)

// GetVersion will return version from application
func GetVersion(c echo.Context) error {
	version := model.Version{Backend: "v0"}
	return c.JSON(http.StatusOK, version)
}
