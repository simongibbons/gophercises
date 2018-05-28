package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Question struct {
	question, answer string
}

type Quiz []Question

func main() {
	quiz, err := readQuizFromCSV("problems.csv")
	if err != nil {
		log.Fatalf("Error parsing quiz csv: %v", err)
	}

	fmt.Printf("Got %d questions for quiz\n", len(quiz))
	quiz.Run(30 * time.Second)
}

func (q Quiz) Run(duration time.Duration) {
	correctCount := 0
	incorrectCount := 0

	outOfQuestions := make(chan bool)

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for _, question := range q {
			fmt.Printf("%s > ", question.question)
			scanner.Scan()
			if scanner.Text() == question.answer {
				correctCount += 1
			} else {
				incorrectCount += 1
			}
		}
		outOfQuestions <- true
	}()

	select {
	case <-outOfQuestions:
		fmt.Println("\nOut of questions")
	case <-time.After(duration):
		fmt.Println("\nOut of time")
	}

	totalAsked := correctCount + incorrectCount
	fmt.Printf("Got (%d/%d) correct.\n", correctCount, totalAsked)
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
		if err != nil {
			if err == io.EOF {
				break
			}
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
