package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
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
	log.Printf("processing bingo winner from %s\n", input)

	f, err := os.Open(input)
	check(err)
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Scan()

	numbers := []int{}
	for _, n := range strings.Split(s.Text(), ",") {
		number, err := strconv.Atoi(n)
		check(err)
		numbers = append(numbers, number)
	}

	// fmt.Println(numbers)

	space := regexp.MustCompile(`\s+`)
	winner := make(chan int)
	draws := []chan int{}
	sem := sync.WaitGroup{}

	s.Scan()
	var board [5][5]int
	boards := 0
	i := 0
	for s.Scan() {
		line := s.Text()

		if line == "" {
			i = 0
			board = [5][5]int{}
			// fmt.Println("new board")
			continue
		}

		normalizedLine := space.ReplaceAllString(line, " ")
		normalizedLine = strings.TrimSpace(normalizedLine)
		// fmt.Printf("line %v\n", normalizedLine)
		for j, n := range strings.Split(normalizedLine, " ") {
			number, err := strconv.Atoi(n)
			check(err)
			board[i][j] = number
		}

		i++

		if i == 5 {
			boards++
			draw := make(chan int)
			draws = append(draws, draw)
			go bingo(board, draw, winner, &sem)
		}
	}

	// draw numbers
	go func() {
		for i, number := range numbers {
			fmt.Printf("Round %d: Draw %d\n", i, number)
			sem.Add(boards)
			for _, draw := range draws {
				draw <- number
			}
			sem.Wait()
		}
	}()

	winnerScore := <-winner
	fmt.Printf("The winner scored: %d points!\n", winnerScore)
}

func bingo(board [5][5]int, draw, winner chan int, sem *sync.WaitGroup) {
	var marks [5][5]int
	for {
		select {
		case number := <-draw:
		out:
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if board[i][j] == number {
						marks[i][j] = 1
						if checkWin(marks) {
							fmt.Println("winner board", board)
							winner <- calculateScore(number, board, marks)
						}
						break out
					}
				}
			}
			sem.Done()
		}
	}
}

func checkWin(marks [5][5]int) bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if marks[i][j] == 0 {
				break
			}
			// win?
			if j == 4 {
				return true
			}
		}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if marks[j][i] == 0 {
				break
			}
			// win?
			if j == 4 {
				return true
			}
		}
	}

	return false
}

func calculateScore(number int, board, marks [5][5]int) int {
	score := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if marks[i][j] == 0 {
				score += board[i][j]
			}
		}
	}
	return score * number
}
