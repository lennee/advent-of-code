package main

import (
	"fmt"
	"crypto/md5"
	"strings"
)

const PUZZLE_INPUT = "bgvyzdsv"

func generate_md5(num int) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", PUZZLE_INPUT, num))))
}

func find_lowest_int(zero_count int) int {
	prefix := strings.Repeat("0", zero_count)
	num_val := 1
	hash_val := ""
	for true {
		hash_val = generate_md5(num_val)
		if strings.HasPrefix(hash_val, prefix) {
			return num_val
		}
		num_val++
	}
	return 0
}

func main() {
	lowest_5 := find_lowest_int(5)
	fmt.Println("Lowest Integer Key (5):", lowest_5)
	lowest_6 := find_lowest_int(6)
	fmt.Println("Lowest Integer Key (6):", lowest_6)
}