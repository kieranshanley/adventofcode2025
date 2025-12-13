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
		processRotation(&dialReading, &pointsToZero, line)
		if err != nil {
			log.Fatal(err)
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
func processRotation(currentDial *int, pointsToZero *uint32, line string) error {
	if len(line) < 2 {
		return errors.New("invalid line format")
	}
	direction := line[0]  // The first byte
	amountStr := line[1:] // The rest of the string
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		log.Printf("Invalid number in line %s: %v", line, err)
		return errors.New("invalid number in line")
	}

	// how many times round the dial have we gone?
	*pointsToZero += uint32(amount / 100)

	// handle rotations that are multiples of 100
	amount = amount % 100

	rollOver := calculateNewPosition(direction, currentDial, amount)
		
	if rollOver {
		*pointsToZero++
	} else if *currentDial == 0 {
		*pointsToZero++
	}

	return nil
}

func calculateNewPosition(direction byte, currentDial *int, amount int) (bool) {
	rollOver := false

	inhibitRollOver := *currentDial == 0

	switch direction {
	case 'R':
		*currentDial += amount
	case 'L':
		*currentDial -= amount
	}

	if *currentDial < 0 {
		*currentDial += 100
		if !inhibitRollOver {
			rollOver = true
		}
	} else if *currentDial >= 100 {
		*currentDial -= 100
		if !inhibitRollOver {
			rollOver = true
		}
	}
	return rollOver
}
