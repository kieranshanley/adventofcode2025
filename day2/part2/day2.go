package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Start uint64
	End   uint64
}

func main() {

	const filePath = "challenge.txt"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open %s: %v", filePath, err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	// Flexible parser: allow varying number of fields per record
	r.FieldsPerRecord = -1
	// Trim leading spaces on fields
	r.TrimLeadingSpace = true
	// Allow unescaped quotes in fields (useful for malformed CSVs)
	r.LazyQuotes = true

	NumberRanges := []Range{}
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to read CSV: %v", err)
		}

		// Skip empty records (e.g., trailing blank lines)
		if len(rec) == 0 {
			continue
		}

		for i := range rec {
			// Each entry has START-END
			// Split the entry on "-"
			if rec[i] != "" {
				parts := strings.Split(rec[i], "-")
				if len(parts) != 2 {
					log.Fatalf("invalid range format: %s", rec[i])
				}

				start, err := strconv.Atoi(parts[0])
				if err != nil {
					log.Fatalf("invalid range format: %s", rec[i])
				}
				end, err := strconv.Atoi(parts[1])
				if err != nil {
					log.Fatalf("invalid range format: %s", rec[i])
				}
				rng := Range{
					Start: uint64(start),
					End:   uint64(end),
				}
				NumberRanges = append(NumberRanges, rng)
			}
		}
	}

	var total uint64 = 0
	for _, idRange := range NumberRanges {
		illegalIdsInRange := checkIds(idRange)
		fmt.Printf("Illegal Ids: %v", illegalIdsInRange)
		for _, illegalId := range illegalIdsInRange {
			total += illegalId
		}
	}
	fmt.Printf("Total = %v", total)
}

func convertStringToUint64(s string) uint64 {
	v64, err := strconv.ParseUint(s, 10, 64) // base 10, 64-bit max
	if err != nil {
		// handle error (invalid number or overflow)
		fmt.Println("parse error:", err)
	}
	v := uint64(v64)
	return v
}

func checkIds(idRange Range) []uint64 {
	illegalIds := map[uint64]struct{}{}
	for i := idRange.Start; i <= idRange.End; i++ {
		var numberAsString = strconv.FormatUint(uint64(i), 10)
		patterns := generatePatterns(numberAsString)

		// look for repetitions of these patterns in the number
		for _, pattern := range patterns {
		out:
			// check if pattern exists in numberAsString
			for i := 0; i <= len(numberAsString)-len(pattern); i++ {
				// we dont want to allow pattern to occur next to each other
				foundIndex := checkForPattern(numberAsString, i, pattern)
				startIndex := foundIndex
				if foundIndex != -1 {
					for {
						// found one occurance - keep looking for more
						startIndex += +len(pattern)
						if startIndex >= len(numberAsString) {
							break out
						}
						// search for another occurance of the pattern
						foundIndex := checkForPattern(numberAsString, startIndex, pattern)
						if foundIndex == -1 {
							break out
						}
						// found adjacent pattern
						if foundIndex == startIndex && (foundIndex == len(numberAsString)-len(pattern)) {
							// found adjacent pattern
							illegalIds[convertStringToUint64(numberAsString)] = struct{}{}
						} else if foundIndex != startIndex {
							break out
						}
					}
				}
			}
		}
	}

	return getIllegalIds(illegalIds)
}

func getIllegalIds(m map[uint64]struct{}) []uint64 {
	var keys []uint64
	for key := range m {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

func getKeys(m map[string]struct{}) []string {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

// generate all possible patterns to search for
func generatePatterns(targetString string) []string {
	searchPatterns := map[string]struct{}{} // use map to avoid duplicates

	if len(targetString) == 0 || len(targetString) == 1 {
		return getKeys(searchPatterns)
	}

	patternLength := len(targetString) / 2
	for {

		if patternLength == 0 {
			break
		}
		numTimes := len(targetString) / patternLength
		if len(targetString)%numTimes == 0 {
			searchPatterns[targetString[0:patternLength]] = struct{}{}
		}
		patternLength--
	}

	return getKeys(searchPatterns)
}

// look for an occurance pattern from startIndex onwards
func checkForPattern(numberAsString string, startIndex int, pattern string) int {
	for i := startIndex; i <= len(numberAsString)-len(pattern); i++ {
		if numberAsString[i:i+len(pattern)] == pattern {
			return i
		}
	}
	return -1
}
