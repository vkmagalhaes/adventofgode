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

	sum1 := 0
	sum1Counter := 0

	sum2 := 0
	sum2Counter := 0

	sum3 := 0
	sum3Counter := 0

	s := bufio.NewScanner(f)

	s.Scan()
	m1, err := strconv.Atoi(s.Text())
	check(err)
	sum1 += m1

	s.Scan()
	m2, err := strconv.Atoi(s.Text())
	check(err)
	sum1 += m2
	sum2 += m2

	sum1Counter = 2
	sum2Counter = 1

	previousSum := 0
	increases := -1 // filter out first comparinson

	for s.Scan() {
		measurement, err := strconv.Atoi(s.Text())
		check(err)

		sum1 += measurement
		sum2 += measurement
		sum3 += measurement

		sum1Counter++
		sum2Counter++
		sum3Counter++

		if sum1Counter == 3 {
			if sum1 > previousSum {
				increases++
			}
			previousSum = sum1
			sum1 = 0
			sum1Counter = 0
		}
		if sum2Counter == 3 {
			if sum2 > previousSum {
				increases++
			}
			previousSum = sum2
			sum2 = 0
			sum2Counter = 0
		}
		if sum3Counter == 3 {
			if sum3 > previousSum {
				increases++
			}
			previousSum = sum3
			sum3 = 0
			sum3Counter = 0
		}
	}

	fmt.Printf("There are %d measurements that are larger than the previous measurement\n", increases)
}
