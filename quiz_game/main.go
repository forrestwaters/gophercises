package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func exit(msg string) {
	fmt.Printf(msg)
	os.Exit(1)
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: strings.TrimSpace(line[0]),
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func askUser(problems []problem) int {
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}
	return correct
}

func main() {

	var fn = flag.String("f", "problems.csv", "Custom filename containing problem questions in the format 'question, answer'.")
	flag.Parse()
	content, err := os.Open(*fn)

	if err != nil {
		exit(fmt.Sprintf("Failed to open file: %s\n", *fn))
	}

	r := csv.NewReader(content)

	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse csv file")
	}

	problems := parseLines(lines)

	correct := askUser(problems)

	content.Close()

	fmt.Printf("You got %#v right out of %#v total!\n", correct, len(problems))
}
