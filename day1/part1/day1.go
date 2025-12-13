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
	file, err := os.Open("rotations.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Ensure the file is closed after the function returns
	defer file.Close()

	// 2. Create a new Scanner for the file
	scanner := bufio.NewScanner(file)

	// 3. Scan() returns true as long as there is another line to read

	var dialReading int = 50

	var pointsToZero uint32 = 0
	for scanner.Scan() {

		// read in a line from the file
		line := scanner.Text()
		processRotation(&dialReading, line)
		if err != nil {
			log.Fatal(err)
		}
		if dialReading == 0 {
			pointsToZero++
		}
	}
	fmt.Printf("Code is %v", pointsToZero)

	// 4. Check for errors that might have occurred during scanning
	// (e.g., file permissions, unexpected EOF)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// Current Dial is passed into processRotation. Note: Since it is
// a pointer the value is modified.
func processRotation(currentDial *int, line string) error {
	if len(line) < 2 {
		return errors.New("Invalid line format")
	}
	direction := line[0]  // The first byte
	amountStr := line[1:] // The rest of the string
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		log.Printf("Invalid number in line %s: %v", line, err)
		return errors.New("Invalid number in line")
	}
	// handle rotations that are multiples of 100
	amount = amount % 100

	switch direction {
	case 'R':
		*currentDial += amount
	case 'L':
		*currentDial -= amount
	}

	if *currentDial < 0 {
		*currentDial += 100
	} else if *currentDial >= 100 {
		*currentDial -= 100
	}

	return nil
}
