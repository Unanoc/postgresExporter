package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"psqlexport/config"
	"psqlexport/database"
	"sync"
	"syscall"
)

var (
	pathToConfig = flag.String("config", "", "Path to configuration file")
	numThreads   = flag.Int("threads", 1, "Threads number")
)

func main() {
	flag.Parse()

	configFile, err := ioutil.ReadFile(*pathToConfig)
	if err != nil {
		log.Panic(err)
	}

	config := config.Config{}
	config.UnmarshalJSON(configFile)

	db := database.DB{}
	if err := db.Connect(config.Connection); err != nil {
		log.Panic(err)
	}

	wg := &sync.WaitGroup{}
	for i := 0; i < *numThreads; i++ {
		wg.Add(1)
		// go CreateCSV(wg, config.OutDir, config.Name, config.Query, config.MaxLines)
	}

	syscallChan := make(chan os.Signal, 1)
	signal.Notify(syscallChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-syscallChan // goroutine will be frozed at here cause it will be wating until signal is received.
		log.Println("Shutting down...")
		db.Disconnect()
		os.Exit(0)
	}()
	wg.Wait()
}
