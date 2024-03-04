package main

import (
	"log"

	"github.com/shailesh-shenoy/allcoinz/api"
	"github.com/shailesh-shenoy/allcoinz/db"
)

func main() {

	s := api.ApiServer{
		ListenAddr: ":8080",
	}

	dataStore := db.NewDataStore("user=postgres password=allcoinz dbname=postgres host=localhost sslmode=disable")

	log.Print("Connecting to postgres ...")
	if err := dataStore.Open(); err != nil {
		log.Fatal(err)
	}
	log.Print("Starting API Server")
	log.Fatal(s.Run())
}
