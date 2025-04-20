package model

import (
	"car_inventory/config"
	"context"
	"database/sql"
	"fmt"
)

type Car struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Model string  `json:"model"`
	Brand string  `json:"brand"`
	Year  int64   `json:"year"`
	Price float64 `json:"price"`
}

func (c *Car) InsertCar() {
	query := "INSERT INTO cars (name, model, brand, year, price) VALUES (?, ?, ?, ?, ?)"

	result, err := config.DB.ExecContext(context.Background(), query, c.Name, c.Model, c.Brand, c.Year, c.Price)
	if err != nil {
		fmt.Printf("Error inserting the car record: %v\n", err)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("Error fetching last insert ID: %v\n", err)
		return
	}

	c.ID = id
	fmt.Printf("✅ Car inserted with ID: %d\n", c.ID)
}

func (c *Car) GetCar(id int64) {
	query := "SELECT id, name, model, brand, year, price FROM cars WHERE id=?"

	row := config.DB.QueryRowContext(context.Background(), query, id)

	err := row.Scan(&c.ID, &c.Name, &c.Model, &c.Brand, &c.Year, &c.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Errorf("Not found for id: %v: err: %v", id, err)
		}

		fmt.Errorf("Error retrieving the data: %v", err)
	}
	fmt.Print("Successfull retrived data")
}

func (c *Car) DeleteCar(id int64) {
	query := "DELETE FROM cars WHERE id=?"

	_, err := config.DB.Exec(query, id)
	if err != nil {
		fmt.Errorf("Error deleting the car record: %v\n", err)
		return
	}

	fmt.Println("Successfully delete the car record with id: %v", id)
}
