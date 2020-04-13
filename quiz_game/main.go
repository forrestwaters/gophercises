package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getInput(question string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(question)
	var answer string
	answer, err := reader.ReadString('\n')

	check(err)

	return answer
}

func main() {

	content, err := os.Open("problems.csv")

	check(err)

	r := csv.NewReader(bufio.NewReader(content))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		check(err)

		answer := getInput(record[0])

		fmt.Println(answer)
	}
}
