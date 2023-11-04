package main

import (
	"fmt"
	"os"
	"bufio"
	"io/ioutil"
	"strings"
)

type coordinates struct {
	x int
	y int
}

const FILENAME = "./day_3/present_delivery_input"

func read_file_to_scanner() *bufio.Scanner {
	filebuffer, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	input_data := string(filebuffer)
	data := bufio.NewScanner(strings.NewReader(input_data))
	data.Split(bufio.ScanRunes)
	return data
}

func update_index(index coordinates, move string) coordinates {
	if move == ">" {
		return coordinates{x: index.x + 1, y: index.y}
	} else if move == "<" {
		return coordinates{x: index.x - 1, y: index.y}
	} else if move == "^" {
		return coordinates{x: index.x, y: index.y + 1}
	} else if move == "v" {
		return coordinates{x: index.x, y: index.y - 1}
	}
	return index
}

func deliver_presents(data *bufio.Scanner) int {
	deliveries := make(map[string]int)
	santa_index := coordinates{x: 0, y: 0}
	deliveries[fmt.Sprintf("%d,%d", santa_index.x, santa_index.y)] += 2
	for data.Scan() {
		santa_index = update_index(santa_index, data.Text())
		deliveries[fmt.Sprintf("%d,%d", santa_index.x, santa_index.y)] += 1
	}
	return len(deliveries)
}

func deliver_presents_alternate(data *bufio.Scanner) int {
	deliveries := make(map[string]int)
	santa_index := coordinates{x: 0, y: 0}
	robo_index := coordinates{x: 0, y: 0}
	deliveries[fmt.Sprintf("%d,%d", santa_index.x, santa_index.y)] += 2
	robo := false
	var index coordinates
	for data.Scan() {
		if robo {
			index = robo_index
		} else {
			index = santa_index
		}

		index = update_index(index, data.Text())
		deliveries[fmt.Sprintf("%d,%d", index.x, index.y)] += 1
	
		if robo {
			robo_index = index
		} else {
			santa_index = index
		}

		robo = !robo
	}

	return len(deliveries)
}


func main() {
	data := read_file_to_scanner()
	fmt.Println("Santa Only Delivery:", deliver_presents(data))
	data = read_file_to_scanner()
	fmt.Println("Robo Santa Assisted Delivery:", deliver_presents_alternate(data))
}