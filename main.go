package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Init()
	fmt.Println("Running API...")

	r := router.Run()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
