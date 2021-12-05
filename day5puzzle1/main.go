package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

type Point struct {
	x int
	y int
}

type Vec struct {
	p1 Point
	p2 Point
}

func run(f *os.File) {
	s := bufio.NewScanner(f)

	vectors := []Vec{}
	for s.Scan() {
		var x1, x2, y1, y2 int
		fmt.Sscanf(s.Text(), "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)

		if x1 != x2 && y1 != y2 {
			continue
		}

		if x2 < x1 || y2 < y1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		vectors = append(vectors, Vec{
			p1: Point{x: x1, y: y1},
			p2: Point{x: x2, y: y2},
		})
	}
	// fmt.Println(vectors)

	scan := scanVents(vectors)
	countOverlaps(scan)
}

func scanVents(vectors []Vec) [1000][1000]int {
	var scan [1000][1000]int
	for _, v := range vectors {
		x, y := v.p1.x, v.p1.y
		for {
			scan[y][x]++
			if x == v.p2.x && y == v.p2.y {
				break
			}
			if x != v.p2.x {
				x++
			}
			if y != v.p2.y {
				y++
			}
		}
	}
	return scan
}

func countOverlaps(scan [1000][1000]int) int {
	counter := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if scan[i][j] > 1 {
				counter++
			}
		}
	}
	fmt.Printf("%d points\n", counter)
	return counter
}

func printMap(scan [1000][1000]int) {
	fmt.Println("Map")
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if scan[i][j] == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", scan[i][j])
			}
		}
		fmt.Println()
	}
}
