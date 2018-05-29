package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type question struct {
	q string
	a string
}

func main() {

	csvFile := flag.String("csv", "problems.csv", "a csv file in the format 'question, answer'")
	f, err := os.Open(*csvFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	rows, err := reader.ReadAll()
	if err != nil {
		handleErr(err)
	}

	questions := parseRows(rows)

	var counter int
	for idx, question := range questions {
		var resp string
		fmt.Printf("Problem #%d: %s = ", (idx + 1), question.q)
		fmt.Scanf("%s\n", &resp)

		if err != nil {
			handleErr(err)
		}

		if resp == question.a {
			counter++
		}
	}

	fmt.Printf("You scored %d out of %d", counter, len(questions))

}

func parseRows(rows [][]string) []question {
	questions := make([]question, len(rows))
	for i, row := range rows {
		questions[i] = question{
			q: row[0],
			a: strings.TrimSpace(row[1]),
		}
	}
	return questions
}

func handleErr(err error) {
	fmt.Errorf("An error occurred %v", err)
	os.Exit(1)
}
