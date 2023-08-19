package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const createTableSQL = `CREATE TABLE IF NOT EXISTS sukiPlayers (username VARCHAR(255) PRIMARY KEY, kills INT, deaths INT, killstreak INT, rankID INT);	`

type Database struct {
	username string
	password string
	host     string
	port     int
	database string

	SqlHandler *sql.DB
}

var DB = Database{username: "db_608589", password: "7577d1545c", host: "na02-db.cus.mc-panel.net", port: 3306, database: "db_608589"}

func OpenDB() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", DB.username, DB.password, DB.host, DB.port, DB.database))
	DB.SetSQLHandler(db)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Practice Database Connected!")
	_, err = DB.SqlHandler.Exec(createTableSQL)
	if err != nil {
		panic(err.Error())
		return
	}
}

func PlayerExists(username string) bool {
	result, err := DB.SqlHandler.Query("SELECT * FROM sukiPlayers WHERE username = '" + username + "'")
	if err != nil {
		panic(err.Error())
		return false
	}

	if result == nil {
		return false
	}
	return true
}

func GetPlayerData(username string) *sql.Rows {
	result, err := DB.SqlHandler.Query("SELECT * FROM sukiPlayers WHERE username = '" + username + "'")
	if err != nil {
		panic(err.Error())
		return nil
	}

	return result
}

func (Database *Database) GetUsername() string {
	return Database.username
}

func (Database *Database) GetPassword() string {
	return Database.password
}

func (Database *Database) GetHost() string {
	return Database.host
}

func (Database *Database) GetPort() int {
	return Database.port
}

func (Database *Database) GetDatabaseName() string {
	return Database.database
}

func (Database *Database) GetSQLHandler() *sql.DB {
	return Database.SqlHandler
}

func (Database *Database) SetSQLHandler(sqlHandler *sql.DB) {
	Database.SqlHandler = sqlHandler
}
