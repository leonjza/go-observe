package utils

import (
	"bufio"
	"fmt"
	"os"
)

// ParseHostsFile parses a file and attempts to extract
// valid hosts from it
func ParseHostsFile(fileName string) ([]string, error) {

	hosts := []string{}

	// try to read entries from a file and submit the results
	if file, err := os.Open(fileName); err == nil {

		// close the file when we are done
		defer file.Close()

		// read the file line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			candidate := scanner.Text()

			host, err := ValidateAndGetURLHost(candidate)
			if err != nil {
				fmt.Printf("Skipping entry %s because of error %s\n", candidate, err)
				continue
			}

			hosts = append(hosts, host)
		}

	} else {
		return hosts, err
	}

	return hosts, nil
}
