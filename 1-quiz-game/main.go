package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Loading questions...")
	fmt.Println()

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

	total := len(records)
	correct := 0

	inputReader := bufio.NewReader(os.Stdin)

	for _, record := range records {
		question := record[0]
		answer := record[1]

		fmt.Printf("%s: ", question)

		input, _, _ := inputReader.ReadLine()

		if string(input) == answer {
			correct++
		}

	}

	fmt.Println()
	fmt.Printf("You got %d out of %d correct\n", correct, total)
}
