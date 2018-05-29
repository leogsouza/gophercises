package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("problems.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	rows, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	var r string
	var counter int
	for idx, row := range rows {
		q, cr := row[0], row[1]
		fmt.Printf("Problem #%d: %s = ", (idx + 1), q)
		fmt.Scanf("%s\n", &r)

		if err != nil {
			panic(err)
		}

		if cr == r {
			counter++
		}
	}

	fmt.Printf("You scored %d out of %d", counter, len(rows))

}
