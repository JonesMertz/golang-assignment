package boardGames

import (
	"errors"
	"golang-assignment/database"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type BoardGameRepository interface {
	GetBoardGames() []BoardGame
	GetBoardGame(id int) (BoardGame, error)
	CreateBoardGame(boardGame BoardGame) (BoardGame, error)
	UpdateBoardGame(id int, boardGame BoardGame) (BoardGame, error)
	DeleteBoardGame(id int) error
	InitializeBoardGameDB()
	SeedBoardGameDB()
}

type DefaultBoardGameRepository struct {
}

func (b *DefaultBoardGameRepository) GetBoardGames() []BoardGame {
	db := database.GetDBConnection()

	rows, err := db.Query(`
		SELECT * FROM board_games
		`)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	fetchedBoardGames := []BoardGame{}

	for rows.Next() {
		boardGame := BoardGame{}
		err = rows.Scan(&boardGame.ID, &boardGame.Name, &boardGame.MinPlayers, &boardGame.MaxPlayers, &boardGame.PlayTime, &boardGame.Age, &boardGame.Description, &boardGame.Price)
		if err != nil {
			log.Fatal(err)
		}

		fetchedBoardGames = append(fetchedBoardGames, boardGame)
	}
	err = rows.Err()

	if err != nil {
		log.Fatal(err)
	}
	return fetchedBoardGames
}

func (b *DefaultBoardGameRepository) GetBoardGame(id int) (BoardGame, error) {
	db := database.GetDBConnection()
	boardGameStatement, err := db.Prepare("select * from board_games where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer boardGameStatement.Close()

	boardGame := BoardGame{}
	err = boardGameStatement.QueryRow(id).Scan(&boardGame.ID, &boardGame.Name, &boardGame.MinPlayers, &boardGame.MaxPlayers, &boardGame.PlayTime, &boardGame.Age, &boardGame.Description, &boardGame.Price)
	if err != nil {
		return boardGame, nil
	}
	return boardGame, nil
}

func (b *DefaultBoardGameRepository) CreateBoardGame(boardGame BoardGame) (BoardGame, error) {
	db := database.GetDBConnection()

	insertSQLStatement, err := db.Prepare(
		`insert into board_games(
			Name,
			MinPlayers,
			MaxPlayers,
			PlayTime,
			Age,
			Description,
			Price		
		) values(?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return boardGame, err
	}
	defer insertSQLStatement.Close()

	result, err := insertSQLStatement.Exec(
		boardGame.Name,
		boardGame.MinPlayers,
		boardGame.MaxPlayers,
		boardGame.PlayTime,
		boardGame.Age,
		boardGame.Description,
		boardGame.Price)
	if err != nil {
		return boardGame, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return boardGame, err
	}
	boardGame.ID = int(id)

	return boardGame, nil
}

func (b *DefaultBoardGameRepository) UpdateBoardGame(id int, boardGame BoardGame) (BoardGame, error) {
	db := database.GetDBConnection()

	updateBoardGameStatement, err := db.Prepare(`UPDATE board_games SET
		Name = ?,
		MinPlayers = ?,
		MaxPlayers = ?,
		PlayTime = ?,
		Age = ?,
		Description = ?,
		Price = ? WHERE id = ?`)

	if err != nil {
		return boardGame, err
	}
	defer updateBoardGameStatement.Close()

	_, err = updateBoardGameStatement.Exec(
		boardGame.Name,
		boardGame.MinPlayers,
		boardGame.MaxPlayers,
		boardGame.PlayTime,
		boardGame.Age,
		boardGame.Description,
		boardGame.Price,
		id)

	if err != nil {
		return boardGame, err
	}

	return boardGame, nil
}

func (b *DefaultBoardGameRepository) DeleteBoardGame(id int) error {
	db := database.GetDBConnection()
	deleteBoardGameStatement, err := db.Prepare("DELETE FROM board_games where id = ?")
	if err != nil {
		return errors.New("error deleting board game")
	}

	defer deleteBoardGameStatement.Close()

	res, err := deleteBoardGameStatement.Exec(id)
	if err != nil {
		return errors.New("error deleting board game")
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return errors.New("error deleting board game")
	}

	if affected == 1 {
		return nil
	}

	return errors.New("unable to delete board game. It may not exist")
}

func (b *DefaultBoardGameRepository) InitializeBoardGameDB() {
	var err error
	db := database.GetDBConnection()
	defer db.Close()

	createTableSQLStatement := `CREATE TABLE IF NOT EXISTS board_games (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		Name TEXT NOT NULL,
		MinPlayers INTEGER NOT NULL,
		MaxPlayers INTEGER NOT NULL,
		PlayTime INTEGER NOT NULL,
		Age INTEGER NOT NULL,
		Description TEXT NOT NULL,
		Price INTEGER NOT NULL
		)`

	_, err = db.Exec(createTableSQLStatement)
	if err != nil {
		log.Printf("%q: %s\n", err, createTableSQLStatement)
		return
	}
}

func (b *DefaultBoardGameRepository) SeedBoardGameDB() {

	// assume that if there are board games in the database, then the database has been seeded
	if len(b.GetBoardGames()) > 0 {
		return
	}

	dummyBoardGames := GetDummyBoardGames()
	for i := 0; i < len(dummyBoardGames); i++ {
		_, err := b.CreateBoardGame(dummyBoardGames[i])
		if err != nil {
			log.Fatal(err)
		}
	}
}
