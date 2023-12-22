package main

import (
	"fmt"
	"go-reloaded/modification"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	// Check if there are 1 or 2 parameters
	if len(args) < 1 || len(args) > 2 {
		fmt.Println("Provide 2 filenames")
		return
	}

	// Check if the files are text files
	for _, fileName := range args {
		if !strings.HasSuffix(fileName, ".txt") {
			fmt.Println("Please provide txt files.")
			return
		}
	}

	// Read sample.txt
	sampleFile := args[0]
	originalText := ReadFile(sampleFile)
	finalText := modification.ModifyTheWholeText(originalText)

	// Create result.txt
	result := ""
	if len(args) == 1 {
		result = "result.txt"
	} else {
		result = args[1]
	}
	WriteFile(result, finalText)
}

func ReadFile(sampleFile string) string {
	data, err := os.ReadFile(sampleFile)
	if err != nil {
		fmt.Println("Please provide source file in txt format")
		return ""
	}

	originalText := string(data)

	return originalText
}

func WriteFile(result, finalText string) {
	resultFile, err := os.Create(result)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Write the text to result.txt
	openResultFile, _ := os.OpenFile(result, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	_, err3 := openResultFile.WriteString(finalText)
	if err3 != nil {
		fmt.Println(err)
		return
	}
	defer resultFile.Close()
}
