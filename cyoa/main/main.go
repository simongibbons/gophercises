package main

import (
	"fmt"
	"github.com/simongibbons/gophercises/cyoa"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	json, err := ioutil.ReadFile("gopher.json")
	if err != nil {
		log.Fatalf("cannot open file error: %s", err)
	}

	adventure, err := cyoa.ParseJSON(json)
	if err != nil {
		log.Fatalf("cannot parse CYOA: %s", err)
	}

	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", cyoa.NewHandler(adventure)))
}
