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
	var fileOpen bool
	var fileName string

	if output[len(output)-1:] == "/" {
		output = fmt.Sprintf("%s%s", output, tableName)
	} else {
		output = fmt.Sprintf("%s/%s", output, tableName)
	}

	if _, err := os.Stat(output); os.IsNotExist(err) {
		os.MkdirAll(output, 0777)
	}

	isRunning := true
	for isRunning {
		select {
		case <-ctx.Done():
			return
		case record, ok := <-recordChan:
			if ok {
				if fileOpen {
					if err := csvWriter.Write(record); err != nil {
						printError(err.Error(), tableName)
						return
					}
					recordCounter++

					if recordCounter == maxLines {
						fileOpen = false
						fileCounter++
						recordCounter = 0

						csvFile.Close()
						csvWriter.Flush()
						if err := csvWriter.Error(); err != nil {
							printError(err.Error(), tableName)
							return
						}
						fmt.Println(fileName)
					}

				} else {
					fileName = fmt.Sprintf("%s/%d.csv", output, fileCounter)
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
					fileOpen = true

					if recordCounter == maxLines {
						fileOpen = false
						fileCounter++
						recordCounter = 0

						csvWriter.Flush()
						if err := csvWriter.Error(); err != nil {
							printError(err.Error(), tableName)
							return
						}
						fmt.Println(fileName)
					}
				}
			} else {
				isRunning = false
				if fileOpen {
					fileOpen = false

					csvFile.Close()
					csvWriter.Flush()
					if err := csvWriter.Error(); err != nil {
						printError(err.Error(), tableName)
						return
					}
					fmt.Println(fileName)
				}
			}
		}
	}
}

func printError(errMsg, tableName string) {
	msg := fmt.Sprintf("%s (table '%s')", errMsg, tableName)
	log.Println(color.RedString(msg))
}

// переделать генератор
// разобраться с контекстом отмены
