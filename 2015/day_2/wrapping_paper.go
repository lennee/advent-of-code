package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

const FILENAME = "./paper_input"

func read_file(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

type present struct {
	x int
	y int
	z int
}

func (p *present) paper() int {
	sides := []int{p.x * p.y, p.x * p.z, p.y * p.z}
	sort.Ints(sides)
	return 3*sides[0] + 2*sides[1] + 2*sides[2]
}

func (p *present) ribbon() int {
	dimensions := []int{p.x, p.y, p.z}
	sort.Ints(dimensions)
	return 2*dimensions[0] + 2*dimensions[1] + dimensions[0]*dimensions[1]*dimensions[2]
}

func process(presents []string) []int {
	ribbon := 0
	paper := 0
	for _, gift := range presents {
		dims := []int{}
		for _, item := range strings.Split(gift, "x") {
			dim, _ := strconv.Atoi(item)
			dims = append(dims, dim)
		}
		p := present{x: dims[0], y: dims[1], z: dims[2]}
		ribbon += p.ribbon()
		paper += p.paper()
	}
	return []int{ribbon, paper}
}

func main() {
	response := process(strings.Split(read_file(FILENAME), "\n"))

	fmt.Println("Paper Amount:", response[0], "\nRibbon Amount:", response[1])
}
