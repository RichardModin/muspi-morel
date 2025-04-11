package main

import (
	"fmt"
	"os"
	"reverse-lorem-ipsum/grouper"
	"reverse-lorem-ipsum/replacer"
)

func main() {
	// File paths
	inputFile := "input.txt"
	outputFile := "output.txt"
	loremFile := "lorem.txt" // File containing the lorem ipsum text

	// Read from input file
	inputText, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Read lorem ipsum text from file
	loremText, err := os.ReadFile(loremFile)
	if err != nil {
		fmt.Println("Error reading lorem file:", err)
		return
	}

	// Generate word map from lorem text
	loremWords := grouper.GroupWordsByLength(string(loremText))

	// Replace words by length
	result := replacer.ReplaceWordsByLength(string(inputText), loremWords)

	// Write to output file
	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}

	fmt.Println("Replacement complete. Output written to", outputFile)
}
