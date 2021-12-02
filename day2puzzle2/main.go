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

func main() {
	if len(os.Args) == 1 {
		log.Fatalln("expected input file but received none")
	}
	input := os.Args[1]
	log.Printf("processing coordinates from %s\n", input)

	f, err := os.Open(input)
	check(err)
	defer f.Close()

	s := bufio.NewScanner(f)
	position := 0
	depth := 0
	aim := 0

	for s.Scan() {
		command := strings.Split(s.Text(), " ")
		value, err := strconv.Atoi(command[1])
		check(err)

		switch action := command[0]; action {
		case "forward":
			position += value
			depth += aim * value
		case "down":
			aim += value
		case "up":
			aim -= value
		default:
			check(fmt.Errorf("unknown command"))
		}
	}

	fmt.Printf("depth (%d) x position (%d) = %d\n", depth, position, depth*position)
}
