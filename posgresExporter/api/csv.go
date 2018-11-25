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

	for {
		select {
		case <-ctx.Done():
			return
		case record := <-recordChan:

			if fileExist {
				err := csvWriter.Write(record)
				if err != nil {
					printError(err.Error(), tableName)
					return
				}

				if recordCounter < maxLines-1 {
					recordCounter++
				} else {
					fileExist = false
					fileCounter++
					recordCounter = 0

					csvWriter.Flush()
					csvFile.Close()
				}
			} else {
				if output[len(output)-1:] == "/" {
					output = fmt.Sprintf("%s%s/", output, tableName)
				} else {
					output = fmt.Sprintf("%s/%s/", output, tableName)
				}

				if _, err := os.Stat(output); os.IsNotExist(err) {
					os.MkdirAll(output, 0777)
				}

				fileName := fmt.Sprintf("%s%d.csv", output, fileCounter)
				csvFile, err := os.Create(fileName)
				if err != nil {
					printError(err.Error(), tableName)
					return
				}

				csvWriter = csv.NewWriter(csvFile)
				err = csvWriter.Write(titles)
				if err != nil {
					printError(err.Error(), tableName)
					return
				}

				err = csvWriter.Write(titles)
				if err != nil {
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
