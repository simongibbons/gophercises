package main

import (
	"github.com/simongibbons/gophercises/cyoa"
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
)

func main() {
	json, err := ioutil.ReadFile("gopher.json")
	if err != nil {
		log.Fatalf("cannot open file error: %s", err)
	}

	adventure , err := cyoa.ParseCYOA(json)
	if err != nil {
		log.Fatalf("cannot parse CYOA: %s", err)
	}

	h := cyoa.NewHandler(adventure)

	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", h))
}
