package main

import (
	v1 "go-simple-web/routes/v1"
	"log"
	"net/http"
)

func main() {

	log.Println("Init route")

	v1.InitRoute()

	log.Println("Listen 8000")

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
