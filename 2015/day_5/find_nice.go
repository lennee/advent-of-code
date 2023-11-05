package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const FILENAME = "./day_5/find_nice_input"
var NAUGHTY_STRINGS = [4]string{"ab", "cd", "pq", "xy"}
const VOWEL_STRING = "aeiou"

func read_file(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func contains_naughty_substring(str string) bool {
	for _, s := range NAUGHTY_STRINGS {
		if strings.Contains(str, s) {
			return true
		}
	}
	return false
}

func check_if_nice_old_way(str string) bool {
	if (contains_naughty_substring(str)) {
		return false
	}
	has_double := false
	vowel_count := 0

	for i, c := range str {
		if (strings.Contains(VOWEL_STRING, string(c))) {
			vowel_count++
		}
		if (i < len(str) - 1 && string(c) == string(str[i + 1])) {
			has_double = true
		}
	}

	return has_double && vowel_count > 2
}

func check_if_nice(str string) bool {
	has_skipped_double := false
	has_pair := false
	for i, c := range str {
		if (i < len(str) - 2 && string(c) == string(str[i + 2])) {
			has_skipped_double = true
		}
		if (i < len(str) - 2 && strings.Contains(string(str[i+2:]), str[i:i+2])) {
			has_pair = true
		}

	}
	
	return has_pair && has_skipped_double
}


func main() {
	result := strings.Split(read_file(FILENAME), "\n")
	count_old := 0
	count := 0
	for _, s := range result {
		if (check_if_nice_old_way(s)) {
			count_old++
		}
		if (check_if_nice(s)) {
			count++
		}
	}
	fmt.Println("Number of Nice Strings (Old Way):", count_old)
	fmt.Println("Number of Nice Strings (New Way):", count)
}