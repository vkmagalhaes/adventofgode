package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	log.Printf("processing measurements from %s\n", input)

	f, err := os.Open(input)
	check(err)
	defer f.Close()

	s := bufio.NewScanner(f)

	s.Scan()
	previous, err := strconv.Atoi(s.Text())
	check(err)
	new := 0
	counter := 0

	for s.Scan() {
		new, err = strconv.Atoi(s.Text())
		check(err)
		if new > previous {
			counter++
		}
		previous = new
	}

	fmt.Printf("There are %d measurements that are larger than the previous measurement\n", counter)

	f.Close()
}
