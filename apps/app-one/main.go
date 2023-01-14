package main

import (
	user "github.com/hillfolk/go-monorepo/domain/users"
	echo "github.com/labstack/echo/v4"
	"net/http"
)

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! App One")
	})
	e.GET("/users", func(c echo.Context) error {
		userList := []user.User{
			{
				Id:        1,
				Email:     "user1@email.com",
				Password:  "password",
				FirstName: "john",
				LastName:  "cena",
				Phone:     "+821012345678",
			},
			{
				Id:        2,
				Email:     "user2@email.com",
				Password:  "password",
				FirstName: "john2",
				LastName:  "cena2",
				Phone:     "+821012345678",
			},
		}
		userList = append(userList, *user.NewUser(3, "user3@email.com", "password", "john3", "cena3", "+821012345678"))
		for i := 0; i < len(userList); i++ {
			println(userList[i].ToString())
		}
		return c.JSON(http.StatusOK, userList)
	})
	e.Logger.Fatal(e.Start(":1323"))

}
