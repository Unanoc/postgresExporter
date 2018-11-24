package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/icrowley/fake"
	"github.com/jackc/pgx"
)

func main() {
	var conn *pgx.ConnPool
	pgxConfig, err := pgx.ParseURI("postgres://testing:@localhost:5432/testing")
	if err != nil {
		log.Panic(err)
	}
	if conn, err = pgx.NewConnPool(
		pgx.ConnPoolConfig{
			ConnConfig: pgxConfig,
		}); err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	fmt.Println("-- TABLE PEOPLE --")
	for i := 0; i < 100; i++ {
		firstname := fake.FirstName()
		lastname := fake.LastName()
		email := fake.EmailAddress()
		phone := fake.Phone()
		text := fake.Paragraph()
		time := time.Now()

		_, err = conn.Exec(`
			INSERT
			INTO people ("firstname", "lastname", "email", "phone", "about", "birthday")
			VALUES ($1, $2, $3, $4, $5, $6)`,
			&firstname, &lastname, &email, &phone, &text, &time)

		if err != nil {
			log.Print(err)
		}
		fmt.Printf("#%d generated...\n", i)
	}

	fmt.Println("-- TABLE CITIES --")
	for i := 0; i < 100; i++ {
		cityName := fake.City()
		population := rand.Int31()

		_, err = conn.Exec(`
			INSERT
			INTO cities ("name", "population")
			VALUES ($1, $2)`,
			&cityName, &population)

		if err != nil {
			log.Print(err)
		}
		fmt.Printf("#%d generated...\n", i)
	}
}
