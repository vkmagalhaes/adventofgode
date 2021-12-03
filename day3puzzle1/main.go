package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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
	var zeros [12]int
	var ones [12]int

	for s.Scan() {
		reading := s.Text()
		for i := 0; i < len(reading); i++ {
			if string(reading[i]) == "1" {
				ones[i]++
			} else {
				zeros[i]++
			}
		}
	}

	var gama [12]int
	var epsilon [12]int
	for i := 0; i < 12; i++ {
		if ones[i] > zeros[i] {
			gama[i] = 1
			epsilon[i] = 0
		} else {
			gama[i] = 0
			epsilon[i] = 1
		}
	}

	gamaDec := 0.0
	for i, bit := range gama {
		gamaDec += math.Pow(2, float64(11-i)) * float64(bit)
	}
	epsilonDec := 0.0
	for i, bit := range epsilon {
		epsilonDec += math.Pow(2, float64(11-i)) * float64(bit)
	}

	fmt.Printf("gama (%v) = %f\n", gama, gamaDec)
	fmt.Printf("epsilon (%v) = %f\n", epsilon, epsilonDec)
	fmt.Printf("gama (%f) x  epsilon (%f) = %f\n", gamaDec, epsilonDec, gamaDec*epsilonDec)
}
