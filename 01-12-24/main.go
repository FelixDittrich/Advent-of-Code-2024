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

	absoluteDistance, err := totalDistance(leftList, rightList)
	if err != nil {
		fmt.Println("Error calculating absolute distance between slice items:", err)
		return
	}

	fmt.Println("Absolute distance: ", absoluteDistance)
}

// Calculate the absolute difference between each slice element
func totalDistance(leftList []int, rightList []int) (int, error){
	var absoluteDistance int = 0

	if len(leftList) != len(rightList) {
		err := fmt.Errorf("Slices are not of equal lenght")
		return 0, err
	}

	for i := 0; i < len(leftList); i++ {
		if leftList[i] < rightList[i] {
			absoluteDistance = absoluteDistance + (rightList[i] - leftList[i])
		} else {
			absoluteDistance = absoluteDistance + (leftList[i] - rightList[i])
		}	
	}
	return absoluteDistance, nil
}