package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	s.Scan()
	run(strings.Split(s.Text(), ","))
}

const days = 256

func run(initialState []string) {
	var fishes [7]int
	for _, f := range initialState {
		days, err := strconv.Atoi(f)
		check(err)
		fishes[days]++
	}

	var babies [7]int
	for i, j := 0, 0; i < days; i, j = i+1, (i+1)%7 {
		if j >= 5 {
			babies[j-5] += fishes[j]
		} else {
			babies[j+2] += fishes[j]
		}
		fishes[j] += babies[j]
		babies[j] = 0

		total := 0
		for i := 0; i < 7; i++ {
			total += fishes[i] + babies[i]
		}
		fmt.Printf("Total fishes after %d days: %d\n", i+1, total)
	}
}
