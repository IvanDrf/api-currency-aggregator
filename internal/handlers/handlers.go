package handlers

import (
	"net/http"

	"github.com/IvanDrf/currency-aggregator/internal/service"
	"github.com/labstack/echo/v4"
)

var currencyList = map[string]bool{
	"cbr":     true,
	"binance": true,
}

func GetListHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, currencyList)
}

func GetCurrencyHandler(ctx echo.Context) error {
	currency := ctx.QueryParam("currency")

	switch currency {
	case "USD":
		return ctx.JSON(http.StatusOK, service.GetCurrency("USD"))
	case "EUR":
		return ctx.JSON(http.StatusOK, service.GetCurrency("EUR"))

	default:
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid query param"})
	}

}
