package main

import (
	"gora/config"
	"log"
)

func main() {
	conf, err := config.New()
	checkFatalError(err)
	conf.Print()

	server := NewServer(conf)

	err = server.Start()
	checkFatalError(err)
}

func checkFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
