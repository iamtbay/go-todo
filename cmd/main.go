package main

import (
	"log"

	"github.com/iamtbay/todo/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//connect db
	err = database.ConnectDB()
	if err != nil {
		panic(err)
	}
	if err != nil {
		log.Println("hello init")
		panic(err)
	}
	//start server
	err = StartServer()
	if err != nil {
		panic(err)
	}

}
