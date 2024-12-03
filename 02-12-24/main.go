package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open input file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening input file", err)
	}
	defer file.Close()

	// Convert file contents into slice of int sices
	reports, err := convertInputFile(file)
	if err != nil {
		log.Fatal("Error converting file", err)
	}
	
	safeReportCount, safeReportCountDampened := countConformity(reports)

	fmt.Println("Number of safe reports: ", safeReportCount)
	fmt.Println("Number of safe reports after dampening: ", safeReportCountDampened)
}

func convertInputFile(file io.Reader) ([][]int, error) {
	// Read file line by line
	scanner := bufio.NewScanner(file)
	var reports [][]int
	for scanner.Scan() {
		// Split elements line by line and convert to integer
		report := scanner.Text()
		tmpLevels := strings.Fields(report)
		levels := make([]int, 0, len(tmpLevels))
		for _, level := range tmpLevels {
			n, err := strconv.Atoi(level)
			if err != nil {
				return nil, err
			}
			levels = append(levels, n)
		}
		reports = append(reports, levels)
	}

	// Handle scanner errors
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return reports, nil
}

func countConformity(reports [][]int) (int, int) {
	safeReportCount := 0
	safeReportCountDampened := 0
	for _, report := range reports {
		if checkForSafety(report, false) {
			safeReportCount++
		}
		if checkForSafety(report, true) {
			safeReportCountDampened++
		}
	}
	return safeReportCount, safeReportCountDampened
}

func checkForSafety(report []int, dampen bool) bool {
	if len(report) <= 1 {
		return true
	}

	isSafe := func(report []int) bool {
		increasing, decreasing := true, true
		for i := 1; i < len(report); i++ {
			// Check absolute difference condition
			diff := int(math.Abs(float64(report[i] - report[i-1])))
			if diff < 1 || diff > 3 {
				return false
			}
			
			// Check monotonicity
			if report[i] > report[i-1] {
				decreasing = false
			}
			if report[i] < report[i-1] {
				increasing = false
			}
			if !increasing && !decreasing {
				return false
			}
		}
		return true
	}

	// check slice as is, else try removing each element and check again
	if isSafe(report) {
		return true
	} else if dampen {
		for i := 0; i < len(report); i++ {
			// This operation causes memory errors without explicitly creating a new slice with make
			modifiedReport := make([]int, 0, len(report)-1)
			modifiedReport = append(modifiedReport, report[:i]...)
			modifiedReport = append(modifiedReport, report[i+1:]...)
			if isSafe(modifiedReport) {
				return true
			}
		}
	}
	return false
}