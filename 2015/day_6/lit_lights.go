package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
)

const FILENAME = "./day_6/lit_lights_input"

type Command string
const (
	Turn_on  Command = "turn on"
	Turn_off Command = "turn off"
	Toggle   Command = "toggle"
)

type Point struct {
	x int
	y int
}

type CommandSet struct {
	command Command
	point_1 Point
	point_2 Point
}

func read_file(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

type Light struct {
	val int
}

type Board struct {
	board [][]Light
}

func (l *Light) process (command Command) {
	if command == Turn_on {
		l.val++
	} else if command == Turn_off {
		l.val--
	} else if command == Toggle {
		l.val += 2
	}

	if (l.val < 0) {
		l.val = 0
	}
}

func (l *Light) process_old (command Command) {
	if command == Turn_on {
		l.val = 1
	} else if command == Turn_off {
		l.val = 0
	} else if command == Toggle {
		if l.val > 0 {
			l.val = 0
		} else {
			l.val = 1
		}
	}
}
func make_board(x int, y int) [][]Light {
	b := make([][]Light, x)
	
	for i := 0; i < x; i++ {
		b[i] = make([]Light, y)
		for j := 0; j < y; j++  {
			b[i][j] = Light{val: 0}
		}
	}

	return b
}

func (b *Board) count_on() int {
	count := 0
	for i := 0; i < len(b.board); i++ {
		for j := 0; j < len(b.board[i]); j++  {
			if b.board[i][j].val > 0 {
				count++
			}
		}
	}
	return count
}

func (b *Board) sum_brightness_value() int {
	count := 0
	for i := 0; i < len(b.board); i++ {
		for j := 0; j < len(b.board[i]); j++  {
			count += b.board[i][j].val
		}
	}
	return count
}

func parse_command_string(command_string string) CommandSet {
	input_array := strings.Split(command_string, " ")
	command := strings.Join(input_array[:len(input_array) - 3], " ")
	
	point_1_string := strings.Split(input_array[len(input_array) - 3], ",")
	x1, _ := strconv.Atoi(point_1_string[0])
	y1, _ := strconv.Atoi(point_1_string[1])

	point_2_string := strings.Split(input_array[len(input_array) - 1], ",")
	x2, _ := strconv.Atoi(point_2_string[0])
	y2, _ := strconv.Atoi(point_2_string[1])

	return CommandSet{
		command: Command(command), 
		point_1: Point{x: x1, y: y1},
		point_2: Point{x: x2, y: y2},
	}
}

func (b *Board) proccess_command (command CommandSet, process_old bool) {
	for i := command.point_1.x; i <= command.point_2.x; i++ {
		for j := command.point_1.y; j <= command.point_2.y; j++ {
			if process_old {
				b.board[i][j].process_old(command.command)
			} else {
				b.board[i][j].process(command.command)
			}
		}
	}
} 

func main () {
	b := Board{board: make_board(1000,1000)}
	b_old := Board{board: make_board(1000,1000)}

	input_lines := strings.Split(read_file(FILENAME), "\n")
	for _, input := range input_lines {
		command := parse_command_string(input)
		b_old.proccess_command(command, true)
		b.proccess_command(command, false)
	}

	fmt.Println("Count of lights on:", b_old.sum_brightness_value())
	fmt.Println("Sum of total brightness:", b.sum_brightness_value())
}