package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
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

	run(f)
}

const runForDays = 256

func run(f *os.File) {
	s := bufio.NewScanner(f)
	s.Scan()
	fishes := []int{}
	for _, f := range strings.Split(s.Text(), ",") {
		fish, err := strconv.Atoi(f)
		check(err)
		fishes = append(fishes, fish)
	}

	fmt.Printf("Initial State: %v\n", fishes)
	fmt.Printf("Total fishes after 0 days: %d\n", len(fishes))
	result := make(chan int)
	sem := &sync.WaitGroup{}
	sem.Add(len(fishes))
	for _, fish := range fishes {
		go breed(uint8(fish), 0, result, sem)
	}
	// for i := 0; i < runForDays; i++ {
	// 	newBreed := []int{}
	// 	for d := 0; d < len(fishes); d++ {
	// 		fishes[d]--
	// 		if fishes[d] < 0 {
	// 			fishes[d] = 6
	// 			newBreed = append(newBreed, 8)
	// 		}
	// 	}
	// 	fishes = append(fishes, newBreed...)
	// 	// fmt.Printf("After %d days: %v\n", i+1, fishes)
	// }

	go func() {
		total := 0
		for {
			count := <-result
			total += count
			fmt.Printf("Total fishes so far: %d\n", total)
			sem.Done()
		}
	}()
	sem.Wait()
}

func breed(clock uint8, initialDay uint16, result chan int, sem *sync.WaitGroup) {
	counter := 1
	for i := initialDay; i < runForDays; i++ {
		if clock == 0 {
			clock = 6
			sem.Add(1)
			go breed(8, i+1, result, sem)
			continue
		}
		clock--
	}
	result <- counter
}
