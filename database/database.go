package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	DB *sql.DB
}

type User struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Admin     bool      `json:"admin"`
	CreatedAt time.Time `json:"createdAt"`
}

type Todo struct {
	Id          string    `json:"id"`
	UserId      string    `json:"userId"`
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	IsCompleted bool      `json:"isCompleted"`
	QuestTime   time.Time `json:"questTime"`
	CreatedAt   time.Time `json:"createdAt"`
}

var dbConn = &DB{}

// SESSIONS

func ConnectDB() error {

	conn, err := sql.Open("postgres", os.Getenv("POSTGRE_URI"))
	if err != nil {
		panic(err)
	}
	err = conn.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Pinged!")
	dbConn.DB = conn
	return nil
}
