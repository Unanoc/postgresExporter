package database

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/jackc/pgx"
)

// Query selects data from table and then sends record by record in chanel.
func Query(conn *pgx.ConnPool, queryString string, recordChan chan<- []string) {

	rows, err := conn.Query(queryString)
	defer rows.Close()

	if err != nil {
		log.Println(err)
	}

	fields := rows.FieldDescriptions()

	titles := make([]string, 1)
	for _, value := range fields {
		titles = append(titles, value.Name)
	}

	recordChan <- titles

	for rows.Next() {
		rowsRecord := make([]string, 1)
		values, _ := rows.Values()
		for _, value := range values {
			if reflect.TypeOf(value).String() == "time.Time" {
				time := value.(time.Time)
				rowsRecord = append(rowsRecord, fmt.Sprintf("%d", time.Unix()))
			} else {
				row := fmt.Sprintln(value)
				rowsRecord = append(rowsRecord, row[:len(row)-1]) // remove "\n"
			}
		}
		recordChan <- rowsRecord
	}

	// return nil
}
