package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	const (
		host     = "localhost"
		port     = 3306
		user     = "root"
		password = "password"
		dbname   = "cars_inventory"
	)

	// Data Source Name (DSN) format: host=localhost port=5432 user=postgres password=password dbname=mydb sslmode=disable
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Errorf("Error oppening the mysql: %v \n", err)
		panic(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Errorf("Error connecting to mysql: %v \n", err)
		panic(err)
	}
	DB = db
	fmt.Println("Successfully connected to the mysql db")

}
