package main

import (
	echo "github.com/labstack/echo/v4"
	"net/http"
)

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! App Two")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
