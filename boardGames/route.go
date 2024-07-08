package boardGames

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) *echo.Echo {
	fmt.Println("Setting up boardGame routes")
	boardGameGroup := e.Group("/boardGames")

	handler := Handler{BoardGameRepository: &DefaultBoardGameRepository{}}
	handler.BoardGameRepository.InitializeBoardGameDB()
	handler.BoardGameRepository.SeedBoardGameDB()

	boardGameGroup.GET("", handler.GetBoardGames)
	boardGameGroup.GET("/:id", handler.GetBoardGame)
	boardGameGroup.POST("", handler.CreateBoardGame)
	boardGameGroup.PUT("/:id", handler.UpdateBoardGame)
	boardGameGroup.DELETE("/:id", handler.DeleteBoardGame)

	return e
}
