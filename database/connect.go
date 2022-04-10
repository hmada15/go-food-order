package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hmada15/go-food-order/config"
)
var DB = Conn()

func Conn() *sql.DB {
	dbConfig := config.NewConfig()
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbConfig.MySQL.Username, dbConfig.MySQL.Password, dbConfig.MySQL.Host, dbConfig.MySQL.Port, dbConfig.MySQL.DatabaseName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
