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

type QuizResponse int

const (
	Correct QuizResponse = iota
	Incorrect
	OutOfQuestions
)

func main() {
	quiz, err := readQuizFromCSV("problems.csv")
	if err != nil {
		log.Fatalf("Error parsing quiz csv: %v", err)
	}

	fmt.Printf("Got %d questions for quiz\n", len(quiz))
	quiz.Run(30 * time.Second)
}

func (q Quiz) Run(duration time.Duration) {
	timer := time.NewTimer(duration)
	correctCount := 0
	incorrectCount := 0

	responses := make(chan QuizResponse)

	scanner := bufio.NewScanner(os.Stdin)

	go func() {
		for _, question := range q {
			fmt.Printf("%s > ", question.question)
			scanner.Scan()
			if scanner.Text() == question.answer {
				responses <- Correct
			} else {
				responses <- Incorrect
			}
		}
		responses <- OutOfQuestions
	}()

QuizLoop:
	for {
		select {
		case response := <-responses:
			{
				if response == Correct {
					correctCount += 1
				} else if response == Incorrect {
					incorrectCount += 1
				} else if response == OutOfQuestions {
					break QuizLoop
				} else {
					panic("Unknown response")
				}
			}
		case <-timer.C:
			fmt.Println("Out of time")
			break QuizLoop
		}
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
