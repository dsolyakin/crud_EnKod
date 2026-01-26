package main

import (
	"crud/internal/logic"
	"crud/internal/repository/postgresql"
	"crud/internal/rest"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	var err error

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	personRepo, err := postgresql.NewPersonRepository(db)
	if err != nil {
		log.Fatal(err)
	}

	err = personRepo.TableCreateQuery()
	if err != nil {
		log.Fatal(err)
	}

	personHandler := &logic.PersonHandler{
		Repo: personRepo,
	}

	e := echo.New()
	rest.RegisterPersonRoutes(e, personHandler)
	e.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
