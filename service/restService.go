package service

import (
	"fmt"
	"golang-assignment/boardGames"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

func Run() {
	restService := echo.New()

	restService.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Try the /boardGames endpoint.")
	})

	boardGames.Routes(restService)

	err := restService.Start(":" + os.Getenv("PORT"))
	if err != nil {
		fmt.Println(err.Error(), time.Now())
	}
}
