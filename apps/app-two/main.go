package main

import (
	"github.com/hillfolk/go-monorepo/domain/orders"
	echo "github.com/labstack/echo/v4"
	zap "go.uber.org/zap"
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
		orderList = append(orderList, *orders.NewOrder(3, 3, 3, 3, 3))
		for i := 0; i < len(orderList); i++ {
			logger, _ := zap.NewProduction()
			defer logger.Sync()
			logger.Info("failed to fetch URL",
				// Structured context as strongly typed Field values.
				zap.Int("orderId", int(orderList[i].Id)),
				zap.Int("userId", int(orderList[i].UserId)),
				zap.Int("productId", int(orderList[i].ProductId)),
				zap.Int("quantity", int(orderList[i].Quantity)),
				zap.String("orderTime", orderList[i].OrderTime.Format("2006-01-02 15:04:05")),
			)
		}
		return c.JSON(http.StatusOK, orderList)
	})
	e.Logger.Fatal(e.Start(":1324"))
}
