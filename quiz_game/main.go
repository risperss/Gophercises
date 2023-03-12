package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv file %s.", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to parse the csv file %s.", *csvFilename))
	}

	problems := parseLines(lines)
	correct := 0

	for i, p := range problems {
		fmt.Printf("Question #%d: %s = ", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)

		if answer == p.a {
			fmt.Println("Correct!")
			correct++
		} else {
			fmt.Println("Incorrect")
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

type problem struct {
	q, a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret
}
