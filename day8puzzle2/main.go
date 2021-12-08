package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func readInput() string {
	if len(os.Args) == 1 {
		log.Fatalln("expected input file but received none")
	}
	input := os.Args[1]
	log.Printf("processing bingo winner from %s\n", input)
	return input
}

func main() {
	input := readInput()
	f, err := os.Open(input)
	check(err)
	defer f.Close()

	s := bufio.NewScanner(f)
	initialState := [][]string{}
	for s.Scan() {
		initialState = append(initialState, strings.Split(s.Text(), " "))
	}
	run(initialState)
}

func run(initialState [][]string) {
	counter := 0
	for _, entry := range initialState {
		// signals := entry[0:10]
		output := entry[11:15]

		for _, digit := range output {
			switch len(digit) {
			//   1   7   4   8
			case 2, 3, 4, 7:
				counter++
			}
		}
	}

	fmt.Printf("Count: %d\n", counter)
}
