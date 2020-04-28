package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getInput(question string) string {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println(strings.TrimSuffix(question, "\n"))

	answer, err := reader.ReadString('\n')

	check(err)

	return strings.TrimSuffix(answer, "\n")
}

func main() {

	var fn = flag.String("f", "problems.csv", "Custom filename containing problem questions.")
	flag.Parse()
	content, err := os.Open(*fn)

	check(err)

	r := csv.NewReader(bufio.NewReader(content))

	var correct, totalQuestions int = 0, 0

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		check(err)

		answer := getInput(record[0])

		if answer == record[1] {
			correct++
		}

		totalQuestions++
	}

	content.Close()

	fmt.Println(fmt.Sprintf("You got %#v right out of %#v total!", correct, totalQuestions))
}
