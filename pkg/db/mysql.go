package db

import (
	"database/sql"
	"fanchann/Go-APIWithArchitecture2/utils"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func MysqlConnect() *sql.DB {
	err := godotenv.Load("./config/.env")
	utils.ErrorWithPanic(err)

	driverName := os.Getenv("DB_DRIVER_NAME")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	url := os.Getenv("DB_URL")
	port := os.Getenv("DB_PORT")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, url, port, database)

	db, _ := sql.Open(driverName, dataSource)

	db.SetConnMaxIdleTime(20 * time.Minute)
	db.SetConnMaxLifetime(11 * time.Minute)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(100)

	return db
}
