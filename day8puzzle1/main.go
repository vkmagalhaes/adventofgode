package main

import (
	"bufio"
	"fmt"

	// "fmt"
	"log"
	"math"
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

type Digit struct {
	signals []string
}

func newDigit(pattern string) *Digit {
	return &Digit{
		signals: strings.Split(pattern, ""),
	}
}

func (d *Digit) Len() int {
	return len(d.signals)
}

func (d *Digit) Include(other *Digit) bool {
out:
	for _, s := range other.signals {
		for _, sig := range d.signals {
			if sig == s {
				continue out
			}
		}
		return false
	}
	return true
}

func (d *Digit) IncludeStr(signal string) bool {
	for _, s := range d.signals {
		if s == signal {
			return true
		}
	}
	return false
}

func (d *Digit) Minus(other *Digit) string {
out:
	for _, s := range d.signals {
		for _, sig := range other.signals {
			if sig == s {
				continue out
			}
		}
		return s
	}
	panic("couldn't find signal")
}

func (d *Digit) Minuss(other *Digit, pattern string) string {
	// fmt.Println("minuss", d, other, pattern)
out:
	for _, s := range d.signals {
		for _, sig := range other.signals {
			if sig == s || s == pattern {
				continue out
			}
		}
		return s
	}
	panic("couldn't find signal")
}

func (d *Digit) MinusStr(signals ...string) string {
out:
	for _, s := range d.signals {
		for _, sig := range signals {
			if sig == s {
				continue out
			}
		}
		return s
	}
	panic("couldn't minus signals")
}

func (d *Digit) Diff(other *Digit) string {
	for _, s := range d.signals {
		if !other.IncludeStr(s) {
			return s
		}
	}
	panic("couldn't find remaining")
}

type Reading struct {
	zero  *Digit
	one   *Digit
	two   *Digit
	three *Digit
	four  *Digit
	five  *Digit
	six   *Digit
	seven *Digit
	eight *Digit
	nine  *Digit

	top  string
	topR string
	topL string
	mid  string
	botR string
	botL string
	bot  string
}

func newReading() *Reading {
	return &Reading{}
}

func (r *Reading) extractTop() {
	if r.seven == nil {
		return
	}

	// fmt.Println("top", r.botR, r.topR, r.seven, r.one)
	if r.botR != "" && r.topR != "" {
		r.top = r.seven.MinusStr(r.botR, r.topR)
	} else if r.one != nil {
		r.top = r.seven.Minus(r.one)
	}
}

func (r *Reading) extractTopR() {
	if r.one == nil || r.botR == "" {
		return
	}
	r.topR = r.one.MinusStr(r.botR)
	r.extractTop()
}

func (r *Reading) extractBotR() {
	if r.two == nil || r.seven == nil {
		return
	}
	r.botR = r.seven.Diff(r.two)
	r.extractTopR()
}

// func (r *Reading) extractTopL() {
// 	if r.one == nil && r.seven == nil {
// 		return
// 	}
// }

func (r *Reading) extractBotL() {
	if r.eight == nil || r.nine == nil {
		return
	}
	r.botL = r.eight.Minus(r.nine)
}

func (r *Reading) extractBot() {
	if r.five == nil || r.four == nil || r.top == "" {
		return
	}
	// fmt.Println("bot", r.five, r.four, r.top)
	r.bot = r.five.Minuss(r.four, r.top)
}

func (r *Reading) extractMid() {
	if r.seven == nil || r.three == nil || r.bot == "" {
		return
	}
	// fmt.Println("mid", r.three, r.seven, r.bot)
	r.mid = r.three.Minuss(r.seven, r.bot)
}

func (r *Reading) isZero(d *Digit) bool {
	// fmt.Println("isZero", r.mid)
	return r.mid != "" && !d.IncludeStr(r.mid)
}

func (r *Reading) isNine(d *Digit) bool {
	return r.four != nil && d.Include(r.four)
}

func (r *Reading) isFive(d *Digit) bool {
	if r.two != nil && r.three != nil {
		return true
	}
	return r.topR != "" && !d.IncludeStr(r.topR)
}

func (r *Reading) isTwo(d *Digit) bool {
	return r.botL != "" && d.IncludeStr(r.botL)
}

func (r *Reading) isThree(d *Digit) bool {
	return r.seven != nil && d.Include(r.seven)
}

func (r *Reading) isSix(d *Digit) bool {
	return r.zero != nil && r.nine != nil
}

func (r *Reading) isSolved() bool {
	return r.one != nil && r.two != nil && r.three != nil &&
		r.four != nil && r.five != nil && r.six != nil &&
		r.seven != nil && r.eight != nil && r.nine != nil && r.zero != nil
}

func (r *Reading) print() {
	fmt.Println(0, r.zero)
	fmt.Println(1, r.one)
	fmt.Println(2, r.two)
	fmt.Println(3, r.three)
	fmt.Println(4, r.four)
	fmt.Println(5, r.five)
	fmt.Println(6, r.six)
	fmt.Println(7, r.seven)
	fmt.Println(8, r.eight)
	fmt.Println(9, r.nine)
	fmt.Println("top", r.top)
	fmt.Println("topR", r.topR)
	fmt.Println("topL", r.topL)
	fmt.Println("mid", r.mid)
	fmt.Println("bot", r.bot)
	fmt.Println("botR", r.botR)
	fmt.Println("botL", r.botL)
}

func (r *Reading) Process(pattern string) int {
	d := newDigit(pattern)
	switch d.Len() {
	case 2:
		r.one = d
		r.extractTop()
		// fmt.Println(1, d)
		return 1
	case 3:
		r.seven = d
		r.extractTop()
		// fmt.Println(7, d)
		return 7
	case 4:
		r.four = d
		// fmt.Println(4, d)
		return 4
	case 7:
		r.eight = d
		// fmt.Println(8, d)
		return 8
	case 5:
		// 5,2,3
		if r.isTwo(d) {
			r.two = d
			// fmt.Println(2, d)
			return 2
		}
		if r.isThree(d) {
			r.three = d
			// fmt.Println(3, d)
			return 3
		}
		if r.isFive(d) {
			r.five = d
			// fmt.Println(5, d)
			r.extractBot()
			r.extractMid()
			return 5
		}
	case 6:
		// 9,6,0
		if r.isZero(d) {
			r.zero = d
			// fmt.Println(0, d)
			return 0
		}
		if r.isNine(d) {
			r.nine = d
			// fmt.Println(9, d)
			r.extractBotL()
			return 9
		}
		if r.isSix(d) {
			// fmt.Println(6, d)
			r.six = d
			return 6
		}
	}
	return -1
}

func run(initialState [][]string) {
	counter := 0
	for _, entry := range initialState {
		// fmt.Println(entry)
		r := newReading()
		output := entry[11:15]

		for i := 0; i < 5; i++ {
			for _, pattern := range entry {
				if pattern == "|" {
					continue
				}
				r.Process(pattern)
			}
			// fmt.Printf("%+v\n", r)
		}
		number := 0
		for i, pattern := range output {
			number += r.Process(pattern) * int(math.Pow10(3-i))
		}
		// r.print()
		fmt.Println(output, number)
		counter += number
	}

	fmt.Printf("Count: %d\n", counter)
}
