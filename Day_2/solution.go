package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func answer1(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	x := 0
	y := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.Fields(scanner.Text())
		value, error := strconv.Atoi(word[1])
		if error != nil {
			return 0 // error
		}
		direction := word[0]
		switch direction {
		case "up":
			y -= value
		case "down":
			y += value
		case "forward":
			x += value
		case "backward":
			x -= value
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return (x * y)
}

func answer2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	x := 0
	y := 0
	aim := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.Fields(scanner.Text())
		value, error := strconv.Atoi(word[1])
		if error != nil {
			return 0 // error
		}
		direction := word[0]
		switch direction {
		case "up":
			aim -= value
		case "down":
			aim += value
		case "forward":
			x += value
			y += aim * value
		case "backward":
			x -= value
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return (x * y)
}

func main() {
	fmt.Println("Answer for 1,", answer1("input.txt"))
	fmt.Println("Answer for 2,", answer2("input.txt"))
}
