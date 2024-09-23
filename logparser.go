package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Parselogs(filename string) (int, []string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return 0, nil, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()
	var errorLogs []string
	var errorCount int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "ERROR") {
			errorCount++
			errorLogs = append(errorLogs, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, nil, fmt.Errorf("could not read file: %w", err)
	}
	return errorCount, errorLogs, nil
}

func printMostRecentErrors(errorLogs []string, n int) {
	if len(errorLogs) == 0 {
		fmt.Println("No ERROR logs found.")
		return
	}

	// Get the most recent `n` ERROR logs
	numLogs := len(errorLogs)
	startIndex := numLogs - n
	if startIndex < 0 {
		startIndex = 0
	}
	recentErrors := errorLogs[startIndex:numLogs]
	fmt.Printf("\nMost recent %d ERROR logs:\n", n)
	for _, log := range recentErrors {
		fmt.Println(log)
	}
}
func main() {
	// Take filename input from the user
	fmt.Print("Enter the log filename: ")
	var filename string
	fmt.Scanln(&filename)

	// Parse logs and get error count and logs
	errorCount, errorLogs, err := Parselogs(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the count of ERROR logs
	fmt.Printf("Total number of ERROR logs: %d\n", errorCount)

	// Print the most recent ERROR logs
	printMostRecentErrors(errorLogs, 6) // Print the most recent 5 ERROR logs

}
