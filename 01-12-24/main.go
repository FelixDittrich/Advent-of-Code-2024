package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// Open input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file", err)
		return
	}
	defer file.Close()

	// Create slices to store the input 
	var leftList []int
	var rightList []int

	// Read input line by line with scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split each line by whitespace
		line := scanner.Text()
		numbers := strings.Fields(line)
		if len(numbers) != 2 {
			fmt.Println("Invalide line format:", line)
			return
		}

		// Convert strings to integers and add to slices
		leftNum, leftErr := strconv.Atoi(numbers[0])
		rightNum, rightErr := strconv.Atoi(numbers[1])
		if leftErr != nil || rightErr != nil {
			fmt.Println("Error converting numbers:", leftErr, rightErr)
			return
		}
		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	// Handle scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Sort slices
	slices.Sort(leftList)
	slices.Sort(rightList)

	// Get answers
	absoluteDistance := calculateTotalDistance(leftList, rightList)
	fmt.Println("Absolute distance: ", absoluteDistance)

	similarityScore := calculateSimilarityScore(leftList, rightList)
	fmt.Println("Similarity score: ", similarityScore) 
}

// Calculate the absolute difference between each slice element
func calculateTotalDistance(leftList, rightList []int) int {
	var absoluteDistance int = 0

	// Calculate absolute difference between values
	for i := 0; i < len(leftList); i++ {
		if leftList[i] < rightList[i] {
			absoluteDistance = absoluteDistance + (rightList[i] - leftList[i])
		} else {
			absoluteDistance = absoluteDistance + (leftList[i] - rightList[i])
		}	
	}
	return absoluteDistance
}

func calculateSimilarityScore(leftList, rightList []int) int {
	// Create a map to store the frequency of number occurrences in the right slice
	rightFrequency := make(map[int]int)
	for _, num := range rightList {
		rightFrequency[num]++
	}

	// Calculate similarity score
	score := 0
	for _, num := range leftList {
		if frequency, exists := rightFrequency[num]; exists {
			score += num * frequency
		}
	}
	return score
}