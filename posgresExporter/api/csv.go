package api

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

// CreateCSV gets records from chanel and then creates CSV files.
func CreateCSV(ctx context.Context, maxLines int, output, tableName string, recordChan <-chan []string) {
	var recordCounter, fileCounter int
	titles := <-recordChan

	var csvFile *os.File
	var csvWriter *csv.Writer
	var fileExist bool

	if output[len(output)-1:] == "/" {
		output = fmt.Sprintf("%s%s", output, tableName)
	} else {
		output = fmt.Sprintf("%s/%s", output, tableName)
	}

	if _, err := os.Stat(output); os.IsNotExist(err) {
		os.MkdirAll(output, 0777)
	}

	for {
		select {
		case <-ctx.Done():
			return
		case record := <-recordChan:
			fmt.Println(record)

			if fileExist {
				if err := csvWriter.Write(record); err != nil {
					printError(err.Error(), tableName)
					return
				}

				if recordCounter < maxLines-1 {
					recordCounter++
				} else {
					fileExist = false
					fileCounter++
					recordCounter = 0

					csvFile.Close()
					csvWriter.Flush()
					if err := csvWriter.Error(); err != nil {
						printError(err.Error(), tableName)
						return
					}
				}
			} else {
				fileName := fmt.Sprintf("%s/%d.csv", output, fileCounter)
				fmt.Println(fileName)
				csvFile, err := os.Create(fileName)
				if err != nil {
					printError(err.Error(), tableName)
					return
				}

				csvWriter = csv.NewWriter(csvFile)

				if err = csvWriter.Write(titles); err != nil {
					printError(err.Error(), tableName)
					return
				}

				if err = csvWriter.Write(record); err != nil {
					printError(err.Error(), tableName)
					return
				}
				recordCounter++
				fileExist = true
			}
		}
	}
}

func printError(errMsg, tableName string) {
	msg := fmt.Sprintf("%s (table '%s')", errMsg, tableName)
	log.Println(color.RedString(msg))
}

// TODO починить момент, когда указано maxLines больше, чем записей в таблице и когда
// файл не полностью заполнен последний

// и почему people таблица создает по maxLines = 2, хотя я указал 1

// переделать генератор

// разобраться с контекстом отмены
