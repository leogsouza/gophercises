package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type question struct {
	q string
	a string
}

func main() {

	csvFile := flag.String("csv", "problems.csv", "a csv file in the format 'question, answer'")
	limit := flag.Int("limit", 30, "time limit to finish the quiz")
	flag.Parse()
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

	timer := time.NewTimer(time.Duration(*limit) * time.Second)

	var counter int
loopquestions:
	for idx, question := range questions {

		fmt.Printf("Problem #%d: %s = ", (idx + 1), question.q)
		respch := make(chan string)

		go func() {
			var resp string
			fmt.Scanf("%s\n", &resp)
			respch <- resp
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break loopquestions
		case resp := <-respch:
			if resp == question.a {
				counter++
			}
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
