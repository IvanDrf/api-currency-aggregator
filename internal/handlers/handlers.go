package handlers

import (
	"net/http"

	"github.com/IvanDrf/currency-aggregator/internal/models"
	"github.com/labstack/echo/v4"
)

func PostHandler(ctx echo.Context) error {
	return nil
}

func GetHandler(ctx echo.Context) error {
	req := models.Request{}

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid body req"})
	}

	return ctx.JSON(http.StatusOK, nil)

}
