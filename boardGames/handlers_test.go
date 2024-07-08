package boardGames

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetBoardGames(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/boardGames", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	mockRepository := MockBoardGameRepository{}
	endpoint := Handler{BoardGameRepository: &mockRepository}

	err := endpoint.GetBoardGames(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var boardGames []BoardGame
		err := json.Unmarshal(rec.Body.Bytes(), &boardGames)
		if assert.NoError(t, err) {
			mockedBoardGames := mockRepository.GetBoardGames()
			for i := range mockedBoardGames {
				boardGame := boardGames[i]
				mockedBoardGame := mockedBoardGames[i]

				assert.Equal(t, boardGame.Name, mockedBoardGame.Name)
				assert.Equal(t, boardGame.MinPlayers, mockedBoardGame.MinPlayers)
				assert.Equal(t, boardGame.MaxPlayers, mockedBoardGame.MaxPlayers)
				assert.Equal(t, boardGame.PlayTime, mockedBoardGame.PlayTime)
				assert.Equal(t, boardGame.Age, mockedBoardGame.Age)
				assert.Equal(t, boardGame.Description, mockedBoardGame.Description)
				assert.Equal(t, boardGame.Price, mockedBoardGame.Price)
			}
		}
	}
}

func TestGetBoardGame(t *testing.T) {
	t.Run("valid request should return 200", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/boardGames", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("3")
		mockRepository := MockBoardGameRepository{}
		endpoint := Handler{BoardGameRepository: &mockRepository}

		mockedBoardGame, _ := mockRepository.GetBoardGame(3)
		err := endpoint.GetBoardGame(c)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, rec.Code)

			var boardGame BoardGame
			err := json.Unmarshal(rec.Body.Bytes(), &boardGame)
			if assert.NoError(t, err) {
				assert.Equal(t, boardGame.Name, mockedBoardGame.Name)
				assert.Equal(t, boardGame.MinPlayers, mockedBoardGame.MinPlayers)
				assert.Equal(t, boardGame.MaxPlayers, mockedBoardGame.MaxPlayers)
				assert.Equal(t, boardGame.PlayTime, mockedBoardGame.PlayTime)
				assert.Equal(t, boardGame.Age, mockedBoardGame.Age)
				assert.Equal(t, boardGame.Description, mockedBoardGame.Description)
				assert.Equal(t, boardGame.Price, mockedBoardGame.Price)
			}
		}
	})

	t.Run("not found should return 404", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/boardGames", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("7654345654")
		mockRepository := MockBoardGameRepository{}
		endpoint := Handler{BoardGameRepository: &mockRepository}

		err := endpoint.GetBoardGame(c)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}
	})

	t.Run("bad request should return 400", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/boardGames", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("shrbfvefwq52")
		mockRepository := MockBoardGameRepository{}
		endpoint := Handler{BoardGameRepository: &mockRepository}

		err := endpoint.GetBoardGame(c)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}

func TestCreateBoardGame(t *testing.T) {
	t.Run("valid request should return 200", func(t *testing.T) {
		mockedBoardGame := BoardGame{Name: "Risk", MinPlayers: 2, MaxPlayers: 6, PlayTime: 120, Age: 10, Description: "A command and conquer game", Price: 40}
		mockedBoardGameJSON, _ := json.Marshal(mockedBoardGame)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/boardGames", strings.NewReader(string(mockedBoardGameJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockRepository := MockBoardGameRepository{}
		endpoint := Handler{BoardGameRepository: &mockRepository}

		err := endpoint.CreateBoardGame(c)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusCreated, rec.Code)

			var boardGame BoardGame
			err := json.Unmarshal(rec.Body.Bytes(), &boardGame)
			if assert.NoError(t, err) {
				assert.Equal(t, boardGame.Name, mockedBoardGame.Name)
				assert.Equal(t, boardGame.MinPlayers, mockedBoardGame.MinPlayers)
				assert.Equal(t, boardGame.MaxPlayers, mockedBoardGame.MaxPlayers)
				assert.Equal(t, boardGame.PlayTime, mockedBoardGame.PlayTime)
				assert.Equal(t, boardGame.Age, mockedBoardGame.Age)
				assert.Equal(t, boardGame.Description, mockedBoardGame.Description)
				assert.Equal(t, boardGame.Price, mockedBoardGame.Price)
			}
		}
	})
	t.Run("bad request should return 400", func(t *testing.T) {

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/boardGames", strings.NewReader("invalid json"))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockRepository := MockBoardGameRepository{}
		endpoint := Handler{BoardGameRepository: &mockRepository}

		err := endpoint.CreateBoardGame(c)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}

type MockBoardGameRepository struct {
}

func (m *MockBoardGameRepository) GetBoardGames() []BoardGame {
	boardGames := GetDummyBoardGames()
	for i := range boardGames {
		boardGames[i].ID = i
	}
	return boardGames
}

func (m *MockBoardGameRepository) GetBoardGame(id int) (BoardGame, error) {
	if id != 3 {
		return BoardGame{}, nil
	}
	boardGames := GetDummyBoardGames()
	boardGames[2].ID = 3
	return boardGames[2], nil
}

func (m *MockBoardGameRepository) CreateBoardGame(boardGame BoardGame) (BoardGame, error) {
	boardGame.ID = 3
	return boardGame, nil
}

func (m *MockBoardGameRepository) UpdateBoardGame(id int, boardGame BoardGame) (BoardGame, error) {
	return boardGame, nil
}

func (m *MockBoardGameRepository) DeleteBoardGame(id int) error {
	return nil
}

func (m *MockBoardGameRepository) InitializeBoardGameDB() {
}

func (m *MockBoardGameRepository) SeedBoardGameDB() {
}
