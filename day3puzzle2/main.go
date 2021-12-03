package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func countBits(readings []string) ([12]int, [12]int) {
	var zeros [12]int
	var ones [12]int
	for _, reading := range readings {
		for i := 0; i < len(reading); i++ {
			if string(reading[i]) == "1" {
				ones[i]++
			} else {
				zeros[i]++
			}
		}
	}
	return zeros, ones
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
	readings := []string{}

	for s.Scan() {
		readings = append(readings, s.Text())
	}

	var oxygen string
	set := readings
	for i := 0; i < 12; i++ {
		zeros, ones := countBits(set)

		most := "1"
		if zeros[i] > ones[i] {
			most = "0"
		}

		newSet := []string{}
		for _, reading := range set {
			if string(reading[i]) == most {
				newSet = append(newSet, reading)
			}
		}
		set = newSet
		if len(set) == 1 {
			oxygen = set[0]
			break
		}
	}

	var co2 string
	set = readings
	for i := 0; i < 12; i++ {
		zeros, ones := countBits(set)

		least := "0"
		if ones[i] < zeros[i] {
			least = "1"
		}

		newSet := []string{}
		for _, reading := range set {
			if string(reading[i]) == least {
				newSet = append(newSet, reading)
			}
		}
		set = newSet
		if len(set) == 1 {
			co2 = set[0]
			break
		}
	}

	oxyDec := 0.0
	for i := 0; i < 12; i++ {
		bit, _ := strconv.Atoi(string(oxygen[i]))
		oxyDec += math.Pow(2, float64(11-i)) * float64(bit)
	}
	co2Dec := 0.0
	for i := 0; i < 12; i++ {
		bit, _ := strconv.Atoi(string(co2[i]))
		co2Dec += math.Pow(2, float64(11-i)) * float64(bit)
	}

	fmt.Printf("oxygen (%s) = %f\n", oxygen, oxyDec)
	fmt.Printf("co2 (%s) = %f\n", co2, co2Dec)
	fmt.Printf("oxygen (%f) x  co2 (%f) = %f\n", oxyDec, co2Dec, oxyDec*co2Dec)
}
