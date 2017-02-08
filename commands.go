package main

import (
	"bufio"
	"fmt"
	"os"
)

func execSubmitCommand(submitURL *string, hide *bool, rescan *bool) {

	if *submitURL == "" {
		fmt.Println("Please provide a URL to submit")
		return
	}

	host, err := validateAndGetURLHost(*submitURL)
	if err != nil {
		fmt.Printf("Failed parsing host from url with error: %s", err)
		return
	}

	result := submitObservatoryAnalysis(host, *hide, *rescan)

	if result.Error != "" {
		fmt.Printf("Error: %s\n", result.Error)
		return
	}

	fmt.Printf("Scan is now:	%s\n", result.State)
}

func execBulkSubmitCommand(filename *string) {

	if *filename == "" {
		fmt.Println("Please provide a filename to read URLs from")
		return
	}

	// try to read entries from a file and submit the results
	if file, err := os.Open(*filename); err == nil {

		// close the file when we are done
		defer file.Close()

		// read the file line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			candidate := scanner.Text()

			host, err := validateAndGetURLHost(candidate)
			if err != nil {
				fmt.Printf("Skipping entry %s because of error %s\n", candidate, err)
				continue
			}

			submitObservatoryAnalysis(host, false, true)
		}

	} else {
		fmt.Printf("Failed to open file: %s", err)
		return
	}
}

func execResultsCommand(url *string, detail *bool) {

	if *url == "" {
		fmt.Println("Please provide a URL to retreive results")
		return
	}

	host, err := validateAndGetURLHost(*url)
	if err != nil {
		fmt.Printf("Failed parsing host from url with error: %s", err)
		return
	}

	result := getObservatoryResults(host)
	fmt.Println("")

	if result.Error != "" {
		fmt.Printf("Error: %s\n", result.Error)
		return
	}

	fmt.Printf("Scan State: 		%s\n", result.State)
	fmt.Printf("Scan Start Time: 	%s\n", result.StartTime)

	if result.State == "FINISHED" {
		fmt.Printf("Grade:			%s\n", result.Grade)
		fmt.Printf("Score:			%d\n", result.Score)
		fmt.Printf("Failed/Passed/Total: 	%d/%d/%d\n", result.TestsFailed, result.TestsPassed, result.TestsQuantity)
		fmt.Printf("Scan ID:		%d\n", result.ScanID)
	}

	if *detail {
		details := getObservatoryDetails(result.ScanID)
		fmt.Println("")

		for _, v := range details {

			fmt.Printf("Pass: %t	Score: %d	Description: %s\n", v.Pass, v.ScoreModifier, v.ScoreDescription)
		}
	}

}

func execBulkResultsCommand(filename *string) {

	if *filename == "" {
		fmt.Println("Please provide the file to read URLs from")
		return
	}

	// try to read entries from a file and get the results
	if file, err := os.Open(*filename); err == nil {

		// close the file when we are done
		defer file.Close()

		// read the file line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			candidate := scanner.Text()

			host, err := validateAndGetURLHost(candidate)
			if err != nil {
				fmt.Printf("Skipping entry %s because of error %s\n", candidate, err)
				continue
			}

			result := getObservatoryResults(host)
			if result.Error != "" {
				fmt.Printf("Error getting result for %s: %s\n", host, result.Error)
				continue
			}

			fmt.Printf("[%s] Scan State: 		%s\n", host, result.State)
			fmt.Printf("[%s] Scan Start Time: 	%s\n", host, result.StartTime)

			if result.State == "FINISHED" {
				fmt.Printf("[%s] Grade:			%s\n", host, result.Grade)
				fmt.Printf("[%s] Score:			%d\n", host, result.Score)
				fmt.Printf("[%s] Failed/Passed/Total: 	%d/%d/%d\n", host, result.TestsFailed, result.TestsPassed, result.TestsQuantity)
				fmt.Printf("[%s] Scan ID:		%d\n", host, result.ScanID)
			}
		}

	} else {
		fmt.Printf("Failed to open file: %s", err)
		return
	}

}
