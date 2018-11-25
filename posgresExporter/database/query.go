package database

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/fatih/color"
	"github.com/jackc/pgx"
)

// Query selects data from table and then sends record by record in chanel.
func Query(conn *pgx.ConnPool, queryString, tableName string, recordChan chan<- []string) {
	rows, err := conn.Query(queryString)
	defer rows.Close()

	if err != nil {
		msg := fmt.Sprintf("%s (table '%s')", err.Error(), tableName)
		log.Println(color.RedString(msg))
		return
	}

	fields := rows.FieldDescriptions()

	titles := make([]string, 1)
	for _, value := range fields {
		titles = append(titles, value.Name)
	}

	recordChan <- titles[1:]

	for rows.Next() {
		rowsRecord := make([]string, 1)
		values, err := rows.Values()
		if err != nil {
			log.Println(color.RedString(err.Error()))
		}
		for _, value := range values {
			if reflect.TypeOf(value).String() == "time.Time" {
				time := value.(time.Time)
				rowsRecord = append(rowsRecord, fmt.Sprintf("%d", time.Unix()))
			} else {
				row := fmt.Sprintln(value)
				rowsRecord = append(rowsRecord, row[:len(row)-1]) // remove "\n"
			}
		}
		recordChan <- rowsRecord[1:]
	}

	if err != nil {
		log.Println(color.RedString(err.Error()), color.BlueString(tableName))
	} else {
		msg := fmt.Sprintf("Table '%s' is successfully exported to CSV.", tableName)
		log.Println(color.GreenString(msg))
	}
}
