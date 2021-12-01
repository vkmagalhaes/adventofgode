package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
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

	channels := []chan int{
		make(chan int),
		make(chan int),
		make(chan int),
	}
	sums := make(chan int)

	for _, c := range channels {
		go func(measurement chan int) {
			sumCounter := 0
			sum := 0
			for m := range measurement {
				sum += m
				sumCounter++
				if sumCounter == 3 {
					sums <- sum
					sumCounter = 0
					sum = 0
				}
			}
			sums <- sum // flush any remaining counter
		}(c)
	}

	go func() {
		previous := 0
		counter := -1 // start with -1 since the first will always be gt

		for new := range sums {
			fmt.Println(new)
			if new > previous {
				counter++
			}
			previous = new
		}

		fmt.Printf("There are %d measurements that are larger than the previous measurement\n", counter)
	}()

	s := bufio.NewScanner(f)

	s.Scan()
	m1, err := strconv.Atoi(s.Text())
	check(err)
	channels[0] <- m1

	s.Scan()
	m2, err := strconv.Atoi(s.Text())
	check(err)
	channels[0] <- m2
	channels[1] <- m2

	for s.Scan() {
		measurement, err := strconv.Atoi(s.Text())
		check(err)

		for _, c := range channels {
			c <- measurement
		}
	}

	for _, c := range channels {
		close(c)
	}

	time.Sleep(time.Second)
	close(sums)
	time.Sleep(time.Second)
}
