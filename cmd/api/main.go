package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

}

func main() {
	fmt.Println("main")
}
