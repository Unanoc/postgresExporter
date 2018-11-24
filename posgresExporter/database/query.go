package database

import (
	"fmt"

	"github.com/jackc/pgx"
)

func Query(conn *pgx.ConnPool, queryString string) error {

	rows, err := conn.Query(queryString)
	defer rows.Close()

	if err != nil {
		return err
	}

	fields := rows.FieldDescriptions()

	for _, value := range fields {
		fmt.Println(value.Name, value.DataTypeName)
	}

	for rows.Next() {
		fmt.Println(rows.Values())
	}

	return nil
}
