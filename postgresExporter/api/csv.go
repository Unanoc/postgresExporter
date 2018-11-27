package api

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/fatih/color"
)

// CreateCSV gets records from chanel and then creates CSV files.
func CreateCSV(ctx context.Context, maxLines int, output, tableName string, recordChan <-chan []string) {
	titles := <-recordChan
	var recordCounter, fileCounter int
	var csvFile *os.File
	var csvWriter *csv.Writer
	var fileOpen bool
	var fileName string

	output = getCorrectDirPath(output, tableName)
	if _, err := os.Stat(output); os.IsNotExist(err) {
		os.MkdirAll(output, 0777)
	}
	fmt.Println(color.BlueString(output))

	checkEqual := func() bool {
		if recordCounter == maxLines {
			fileOpen = false
			fileCounter++
			recordCounter = 0

			csvFile.Close()
			csvWriter.Flush()
			if err := csvWriter.Error(); err != nil {
				printError(err.Error(), tableName)
				return true
			}
			runtime.Gosched()
		}
		return false
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
					if checkEqual() {
						return
					}
				} else {
					fileName = fmt.Sprintf("%s/%d.csv", output, fileCounter)
					csvFile, err := os.Create(fileName)
					if err != nil {
						printError(err.Error(), tableName)
						return
					}

					csvWriter = csv.NewWriter(csvFile)
					csvWriter.WriteAll([][]string{titles, record}) // добавление заголовка в начало файла + добавление первой записи
					if err := csvWriter.Error(); err != nil {
						printError(err.Error(), tableName)
						return
					}

					recordCounter++
					fileOpen = true
					if checkEqual() {
						return
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
				}
			}
		}
	}
}

func printError(errMsg, tableName string) {
	msg := fmt.Sprintf("%s (table '%s')", errMsg, tableName)
	log.Println(color.RedString(msg))
}

func getCorrectDirPath(path, tableName string) string {
	if path[len(path)-1:] == "/" {
		path = fmt.Sprintf("%s%s", path, tableName)
	} else {
		path = fmt.Sprintf("%s/%s", path, tableName)
	}
	return path
}
