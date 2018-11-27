package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"psqlexport/api"
	"psqlexport/config"
	"psqlexport/database"
	"sync"
	"syscall"
)

var (
	pathToConfig = flag.String("config", "", "Path to configuration JSON file")
	numThreads   = flag.Int("threads", 1, "Threads number")
)

func main() {
	flag.Parse()

	configFile, err := ioutil.ReadFile(*pathToConfig)
	if err != nil {
		log.Panic(err)
	}

	configInstance := config.Config{}
	configInstance.UnmarshalJSON(configFile)

	db := database.DB{}
	if err := db.Connect(configInstance.Connection); err != nil {
		log.Panic(err)
	}

	wg := &sync.WaitGroup{}
	ctx, finish := context.WithCancel(context.Background())

	if *numThreads > len(configInstance.Tables) {
		*numThreads = len(configInstance.Tables)
	}

	taskChan := make(chan config.Table, *numThreads)
	for i := 0; i < *numThreads; i++ {
		wg.Add(1)
		go api.WorkerExport(ctx, wg, db.Conn, configInstance.OutputDir, taskChan)
	}

	for _, table := range configInstance.Tables {
		taskChan <- *table
	}

	syscallChan := make(chan os.Signal, 1)
	signal.Notify(syscallChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-syscallChan
		log.Println("Shutting down...")
		finish()
		db.Disconnect()
		os.Exit(0)
	}()
	wg.Wait()
}
