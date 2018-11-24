package main

import (
	"flag"
	"io/ioutil"
	"log"
	"psqlexport/config"
	"psqlexport/database"
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

	err = database.Query(db.Conn, "select firstname, lasname, birthday, email, phone, about, nickname from people")
	if err != nil {
		log.Print(err)
	}
	// wg := &sync.WaitGroup{}
	// stopChan := make(chan struct{})
	// for i := 0; i < *numThreads; i++ {
	// 	wg.Add(1)
	// 	go api.Export(wg, stopChan, config.OutDir, config.Name, config.Query, config.MaxLines)
	// }

	// syscallChan := make(chan os.Signal, 1)
	// signal.Notify(syscallChan, syscall.SIGINT, syscall.SIGTERM)

	// go func() {
	// 	<-syscallChan // goroutine will be frozed at here cause it will be wating until signal is received.
	// 	log.Println("Shutting down...")
	// 	db.Disconnect()
	// 	close(stopChan)
	// 	os.Exit(0)
	// }()
	// wg.Wait()
}
