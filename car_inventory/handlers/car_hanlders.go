package handlers

import (
	"car_inventory/config"
	model "car_inventory/models"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

var (
	sm sync.Mutex
)

/*
net/http handlers

	func CarHandler(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetCarHandler(w, r)
		case http.MethodDelete:
			deleteCarHandler(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}

	func AddCarHandler(w http.ResponseWriter, r *http.Request) {
		sm.Lock()
		defer sm.Unlock()

		car := &model.Car{}

		if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
			fmt.Errorf("Invalid request body: %v", err)
			http.Error(w, "Invalid request body:", http.StatusBadRequest)
			return
		}

		car.InsertCar()
		if car.ID != 0 {
			fmt.Println("Successfully created with id: %v", car.ID)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(car)
		}

		w.WriteHeader(http.StatusInternalServerError)
	}

	func GetCarHandler(w http.ResponseWriter, r *http.Request) {
		sm.Lock()
		defer sm.Unlock()
		car := &model.Car{}
		idStr := strings.TrimPrefix(r.URL.Path, "/cars/")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		car.GetCar(id)

		json.NewEncoder(w).Encode(car)
	}

	func deleteCarHandler(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/cars/")
		id, err := strconv.ParseInt(idStr, 10, 64)
		car := &model.Car{}
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		car.DeleteCar(id)

		w.WriteHeader(http.StatusNoContent) // No content to return
	}
*/
func AddCarHandler(c *fiber.Ctx) error {
	sm.Lock()
	defer sm.Unlock()

	car := &model.Car{}

	if err := c.BodyParser(car); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "invalid body",
			"details": err.Error(),
		})
	}

	car.InsertCar()
	fmt.Println("Successfully created the car recorded with an id: %v", car.ID)
	return c.Status(fiber.StatusCreated).JSON(car)
}

func GetCarHandler(c *fiber.Ctx) error {
	sm.Lock()
	defer sm.Unlock()
	car := &model.Car{}

	// Check if the car already present in cache
	// if not then only goto mysql db

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "invalid id param",
			"detail": err.Error(),
		})
	}

	val, err := config.Cache.Get(c.Context(), strconv.FormatInt(int64(id), 10)).Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Printf("Key is not added in the cache: %v \n", err)
		}
		fmt.Printf("Unable to get the key from redis: %v \n", err)
	} else {
		if err := json.Unmarshal([]byte(val), &car); err != nil {
			fmt.Printf("Unable to unmarshall the value from redis: %v \n", err)
		}
		return c.Status(fiber.StatusOK).JSON(val)
	}

	b, _ := json.Marshal(car)
	config.Cache.Set(c.Context(), strconv.FormatInt(int64(id), 10), b, 60*time.Minute)

	car.GetCar(int64(id))
	fmt.Println("Successfull retrieved car record: id: %v", id)
	return c.Status(fiber.StatusOK).JSON(car)
}

func DeleteCarHandler(c *fiber.Ctx) error {

	car := &model.Car{}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "invalid id param",
			"detail": err.Error(),
		})
	}

	car.DeleteCar(int64(id))

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"msg": "Successfully deleted the record",
		"id":  int64(id),
	})
}
