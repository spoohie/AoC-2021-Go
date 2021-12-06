package utils

import (
	"bufio"
	"log"
	"os"
)

func ParseFile() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func Sum(array [9]int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
