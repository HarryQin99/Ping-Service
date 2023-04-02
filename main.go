package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// Get url from env
	url := os.Getenv("URL")

	go pingUrl(url)

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	log.Println("Shut down")
}

func pingUrl(url string) {
	for {
		_, err := http.Get(url)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println(url)
		time.Sleep(10 * time.Second)
	}
}
