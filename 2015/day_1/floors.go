package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

const FILENAME = "./floors_input"

func read_file(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func find_floor(input string, first_negative bool) int {
	out := 0
	for pos, char := range input {
		if char == '(' {
			out++
		} else if char == ')' {
			out--
		}
		if first_negative && out < 0 {
			return pos + 1
		}
	}
	return out
}

func main() {
	data := read_file(FILENAME)
	final_floor := find_floor(data, false)
	first_negative_floor := find_floor(data, true)
	fmt.Println("Final Floor:", final_floor, "\nFirst Negative Floor:", first_negative_floor)
}
