package main

import (
	"fmt"
	"sync"
)

type singleton struct {
	name  string // keep the fields unexported (Starts with small case) to restrict the modification of this variable
	value int
}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{
			name:  "Singleton pattern",
			value: 5,
		}
	})
	return instance
}

func main() {
	s1 := GetInstance()
	s2 := GetInstance()

	if s1 == s2 {
		fmt.Println("Same instance..")
	}
}
