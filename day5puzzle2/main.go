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

	result := mapVents(vectors)
	countOverlaps(result)
	// printMap(result)
}

const mapSize = 1000

func mapVents(vectors []Vec) [mapSize][mapSize]int {
	var scan [mapSize][mapSize]int
	for _, v := range vectors {
		x, y := v.p1.x, v.p1.y
		for {
			// fmt.Printf("%d,%d\n", x, y)
			scan[y][x]++
			if x == v.p2.x && y == v.p2.y {
				break
			}
			if x != v.p2.x {
				if v.p2.x < v.p1.x {
					x--
				} else {
					x++
				}
			}
			if y != v.p2.y {
				if v.p2.y < v.p1.y {
					y--
				} else {
					y++
				}
			}
		}
	}
	return scan
}

func countOverlaps(scan [mapSize][mapSize]int) int {
	counter := 0
	for i := 0; i < mapSize; i++ {
		for j := 0; j < mapSize; j++ {
			if scan[i][j] > 1 {
				counter++
			}
		}
	}
	fmt.Printf("%d points\n", counter)
	return counter
}

func printMap(scan [mapSize][mapSize]int) {
	fmt.Println("Map")
	for i := 0; i < mapSize; i++ {
		for j := 0; j < mapSize; j++ {
			if scan[i][j] == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", scan[i][j])
			}
		}
		fmt.Println()
	}
}
