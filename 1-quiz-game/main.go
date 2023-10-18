package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Printf("Loading questions...\n\n")

	filePath := "./quiz.csv"

	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("You will asked up to 10 questions or as many as you can answer in 15 seconds")

	quizCh := make(chan bool)
	timerCh := make(chan interface{})

	go func() {
		time.Sleep(time.Second * 15)
		timerCh <- "Time is up"
	}()

	total := len(records)
	answered := 0
	correct := 0

	go func() {
		inputReader := bufio.NewReader(os.Stdin)

		for _, record := range records {
			question := record[0]
			answer := record[1]

			fmt.Printf("%s: ", question)

			input, _, _ := inputReader.ReadLine()

			quizCh <- (string(input) == answer)
		}
	}()

Quiz:
	for {
		select {
		case result := <-quizCh:
			answered++
			if result {
				correct++
			}
			if answered == total {
				break Quiz
			}
		case <-timerCh:
			fmt.Println("\n\nTime is up!")
			break Quiz
		}
	}

	fmt.Printf("\nYou answered %d questions out of %d and got %d correct\n", answered, total, correct)
}
