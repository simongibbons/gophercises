package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

type Question struct {
	question string
	answer   string
}

type Quiz []Question

func main() {
	quiz, err := readQuizFromCSV("problems.csv")
	if err != nil {
		log.Fatalf("Error parsing quiz csv: %v", err)
	}

	fmt.Printf("Got %d questions for quiz\n", len(quiz))
}

func readQuizFromCSV(filename string) (q Quiz, err error) {
	csvFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		if len(line) != 2 {
			return nil, errors.New("expected 2 columns in csv record")
		}

		q = append(q, Question{
			question: line[0],
			answer:   line[1],
		})
	}

	return q, nil
}
