package config

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Cache *redis.Client

func ConnectDB() {
	const (
		host     = "localhost"
		port     = 3306
		user     = "root"
		password = "password"
		dbname   = "cars_inventory"
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error opening MySQL: %v\n", err)
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("Error getting database instance: %v\n", err)
		panic(err)
	}

	if err := sqlDB.Ping(); err != nil {
		fmt.Printf("Error connecting to MySQL: %v\n", err)
		panic(err)
	}

	DB = db // <- this is the correct assignment
	fmt.Println("Successfully connected to the MySQL database")
}

func ConnectCache() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	cmd := rdb.Ping(ctx)
	if cmd.Err() != nil {
		fmt.Printf("Error connecting to redis cache: %v\n", cmd.Err())
		panic(cmd.Err())
	}

	Cache = rdb
	fmt.Println("Successfully connected to the redis cache")

}
