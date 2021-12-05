package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func answer1(input string) int64 {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Looking at our input file, the binary string is 12 bits long
	mostCommon := make([]int, 12)
	gamma := make([]string, 12)
	epsilon := make([]string, 12)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		binaryString := strings.Fields(scanner.Text())
		for _, line := range binaryString {
			for j, bit := range line {
				if string(bit) == "1" {
					mostCommon[j]++
				} else {
					mostCommon[j]--
				}
			}
		}
	}

	for i, bit := range mostCommon {
		if bit > 0 {
			gamma[i] = "1"
			epsilon[i] = "0"
		} else {
			gamma[i] = "0"
			epsilon[i] = "1"
		}
	}

	gammaDecimel, _ := strconv.ParseInt(strings.Join(gamma, ""), 2, 64)
	epsilonDecimel, _ := strconv.ParseInt(strings.Join(epsilon, ""), 2, 64)

	return (gammaDecimel * epsilonDecimel)
}

func answer2(input string) int64 {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Sample input for debugging
	// lines = []string{
	// 	"00100",
	// 	"11110",
	// 	"10110",
	// 	"10111",
	// 	"10101",
	// 	"01111",
	// 	"00111",
	// 	"11100",
	// 	"10000",
	// 	"11001",
	// 	"00010",
	// 	"01010",
	// }

	keptOxygen := lines
	oxygenCounter := 0
	for len(keptOxygen) > 1 {
		ones := 0
		for _, line := range keptOxygen {
			if string(line[oxygenCounter]) == "1" {
				ones++
			} else {
				ones--
			}
		}

		newArray := make([]string, 0)
		for _, line := range keptOxygen {
			if line[oxygenCounter] == '1' && ones >= 0 {
				newArray = append(newArray, line)
			}
			if line[oxygenCounter] == '0' && ones < 0 {
				newArray = append(newArray, line)
			}
		}
		keptOxygen = newArray
		oxygenCounter++
	}

	keptCO2 := lines
	CO2Counter := 0
	for len(keptCO2) > 1 {
		ones := 0
		for _, line := range keptCO2 {
			if string(line[CO2Counter]) == "1" {
				ones++
			} else {
				ones--
			}
		}

		newArray := make([]string, 0)
		for _, line := range keptCO2 {
			if line[CO2Counter] == '0' && ones >= 0 {
				newArray = append(newArray, line)
			}
			if line[CO2Counter] == '1' && ones < 0 {
				newArray = append(newArray, line)
			}
		}
		keptCO2 = newArray
		CO2Counter++
	}

	oxygen, _ := strconv.ParseInt(strings.Join(keptOxygen, ""), 2, 64)
	co2, _ := strconv.ParseInt(strings.Join(keptCO2, ""), 2, 64)
	return oxygen * co2
}

func main() {
	fmt.Println("Answer for 1,", answer1("input.txt"))
	fmt.Println("Answer for 2,", answer2("input.txt"))
}
