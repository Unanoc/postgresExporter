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
	pgxConfig, err := pgx.ParseURI("postgres://testing:testing@localhost:5432/testing")
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

	fmt.Println("-- TABLE USERS --")
	for i := 0; i < 1000; i++ {
		firstname := fake.FirstName()
		lastname := fake.LastName()
		email := fake.EmailAddress()
		phone := fake.Phone()
		text := fake.Paragraph()
		time := time.Now()

		_, err = conn.Exec(`
			INSERT
			INTO users ("firstname", "lastname", "email", "phone", "about", "birthday")
			VALUES ($1, $2, $3, $4, $5, $6)`,
			&firstname, &lastname, &email, &phone, &text, &time)

		if err != nil {
			log.Print(err)
		}
	}
	fmt.Println("Table 'users' is generated...")

	fmt.Println("-- TABLE VILAGES --")
	for i := 0; i < 1000; i++ {
		cityName := fake.City()
		population := rand.Int31()

		_, err = conn.Exec(`
			INSERT
			INTO vilages ("name", "population")
			VALUES ($1, $2)`,
			&cityName, &population)

		if err != nil {
			log.Print(err)
		}
	}
	fmt.Println("Table 'vilages' is generated...")

	fmt.Println("-- TABLE PEOPLE --")
	for i := 0; i < 1000; i++ {
		firstname := fake.FirstName()
		lastname := fake.LastName()
		flag := rand.Int31()
		time := time.Now()

		_, err = conn.Exec(`
			INSERT
			INTO people ("name", "lastname", "birthday", "some_flag", "created")
			VALUES ($1, $2, current_timestamp, $3, $4)`,
			&firstname, &lastname, &flag, &time)

		if err != nil {
			log.Print(err)
		}
	}
	fmt.Println("Table 'people' is generated...")

	fmt.Println("-- TABLE CITIES --")
	for i := 0; i < 1000; i++ {
		cityName := fake.City()
		countryID := rand.Int31()
		time := time.Now()

		_, err = conn.Exec(`
			INSERT
			INTO cities ("name", "country_id", "created")
			VALUES ($1, $2, $3)`,
			&cityName, &countryID, &time)

		if err != nil {
			log.Print(err)
		}
	}
	fmt.Println("Table 'cities' is generated...")

	fmt.Println("-- TABLE COUNTRIES --")
	for i := 0; i < 1000; i++ {
		country := fake.Country()
		time := time.Now()

		_, err = conn.Exec(`
			INSERT
			INTO countries ("name", "created")
			VALUES ($1, $2)`,
			&country, &time)

		if err != nil {
			log.Print(err)
		}
	}
	fmt.Println("Table 'countries' is generated...")
}
