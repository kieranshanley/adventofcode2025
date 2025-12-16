package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// 1. Open the file
	file, err := os.Open("challenge.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Ensure the file is closed after the function returns
	defer file.Close()

	// 2. Create a new Scanner for the file
	scanner := bufio.NewScanner(file)

	// 3. Scan() returns true as long as there is another line to read

	var totalJolts uint64 = 0
	for scanner.Scan() {

		// read in a line from the file
		line := scanner.Text()

		jolts, err := processLine(line)
		totalJolts += *jolts

		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Total Jolts = %v", totalJolts)

	// 4. Check for errors that might have occurred during scanning
	// (e.g., file permissions, unexpected EOF)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func findLargestNumber(nums []uint64) uint64 {
	if len(nums) == 0 {
		return 0 // handle empty slice case
	}
	largest := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > largest {
			largest = nums[i]
		}
	}
	return largest
}

func processLine(line string) (*uint64, error) {

	if len(line) < 15 {
		return nil, errors.New("invalid line format")
	}

	numbers := []uint64{}
	// generate all the combinations
	for i := 0; i < len(line)-1; i++ {
		for j := i + 1; j < len(line); j++ {
			firstNumber := line[i : i+1]
			firstValue, err := strconv.ParseUint(firstNumber, 10, 64)

			if err != nil {
				log.Printf("Invalid number in line %s: %v", line, err)
				return nil, errors.New("invalid number in line")
			}
			secondNumber := line[j : j+1]
			secondValue, err := strconv.ParseUint(secondNumber, 10, 64)
			if err != nil {
				log.Printf("Invalid number in line %s: %v", line, err)
				return nil, errors.New("invalid number in line")
			}
			comboValue := firstValue*10 + secondValue
			fmt.Printf("Combo: %s,%s = %d\n", firstNumber, secondNumber, comboValue)
			numbers = append(numbers, comboValue)
		}
	}
	// and find the hightest combo
	largestJolt := findLargestNumber(numbers)
	fmt.Printf("Largest Jolt = %d\n", largestJolt)

	return &largestJolt, nil
}
