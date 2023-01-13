package main

import (
	"github.com/hillfolk/go-monorepo/domain/orders"
	echo "github.com/labstack/echo/v4"
	"net/http"
)

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! App Two")
	})
	e.GET("/orders", func(c echo.Context) error {
		orderList := []orders.Order{
			{
				Id:        1,
				UserId:    1,
				ProductId: 1,
				Quantity:  1,
			},
			{
				Id:        2,
				UserId:    2,
				ProductId: 2,
				Quantity:  2,
			},
		}
		return c.JSON(http.StatusOK, orderList)
	})
	e.Logger.Fatal(e.Start(":1324"))
}
