package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	crabs := []int{}
	start, end := 0, 0
	for _, p := range initialState {
		c, err := strconv.Atoi(p)
		check(err)
		crabs = append(crabs, c)
		if c < start {
			start = c
		}
		if c > end {
			end = c
		}
	}

	min := end * end
	pos := 0
	fmt.Println(crabs, min, start, end)
	for i := 0; i < end; i++ {
		cost := 0
		for _, crab := range crabs {
			cost += int(math.Abs(float64(i - crab)))
		}
		if cost < min {
			min = cost
			pos = i
		}
	}
	fmt.Printf("Least fuel %d at position %d\n", min, pos)
}
