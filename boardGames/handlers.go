package boardGames

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	BoardGameRepository BoardGameRepository
}

func (h *Handler) GetBoardGames(c echo.Context) error {
	return c.JSON(http.StatusOK, h.BoardGameRepository.GetBoardGames())
}

func (h *Handler) GetBoardGame(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid ID")
	}

	boardGame, err := h.BoardGameRepository.GetBoardGame(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	if boardGame.ID == 0 {
		return c.JSON(http.StatusNotFound, "board game not found")
	}

	return c.JSON(http.StatusOK, boardGame)
}

func (h *Handler) CreateBoardGame(c echo.Context) error {
	boardGame := new(BoardGame)

	err := c.Bind(boardGame)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	createdBoardGame, creationErr := h.BoardGameRepository.CreateBoardGame(*boardGame)

	if creationErr != nil {
		return c.JSON(http.StatusInternalServerError, "error creating board game")
	}

	return c.JSON(http.StatusCreated, createdBoardGame)
}

func (h *Handler) UpdateBoardGame(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid ID")
	}

	boardGame := new(BoardGame)

	err = c.Bind(boardGame)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "board game not found")
	}

	boardGameResult, updateErr := h.BoardGameRepository.UpdateBoardGame(id, *boardGame)

	if updateErr != nil {
		return c.JSON(http.StatusInternalServerError, "error updating board game")
	}

	return c.JSON(http.StatusOK, boardGameResult)
}

func (h *Handler) DeleteBoardGame(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid ID")
	}

	err = h.BoardGameRepository.DeleteBoardGame(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, "board game deleted")
}
