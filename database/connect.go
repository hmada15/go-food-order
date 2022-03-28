package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	database "github.com/hmada15/go-food-order/database/sqlc"
)

func Conn() *database.Queries {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/go-food-order?parseTime=true")
	if err != nil {
		panic(err)
	}
	database := database.New(db)
	return database
}
