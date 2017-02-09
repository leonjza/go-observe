package observatory

import (
	"errors"
	"fmt"
	"strconv"
)

// SubmitObservatoryAnalysis submits a URL for analysis
func SubmitObservatoryAnalysis(observatoryHost string, nohide bool, rescan bool) (ScanObject, error) {

	fmt.Printf("Submitting analysis for %s\n", observatoryHost)

	queryString := make(map[string]string)
	queryString["host"] = observatoryHost

	// by default, hide results from the observatory
	requestBody := make(map[string]string)
	requestBody["hidden"] = "true"

	if nohide {
		fmt.Println("Removing 'hidden' flag for this scan on the observatory")
		delete(requestBody, "hidden")
	}

	if rescan {
		fmt.Println("Forcing a rescan on the observatory")
		requestBody["rescan"] = "true"
	}

	results := ScanObject{}
	err := callObservatory("post", "analyze", &results, queryString, requestBody)
	if err != nil {
		return results, err
	}

	// If there was an error, return that too
	if results.Error != "" {
		return results, errors.New(results.Error)
	}

	return results, nil
}

// GetObservatoryResults gets the results of a hostname
func GetObservatoryResults(observatoryHost string) (ScanObject, error) {

	fmt.Printf("Getting results for: %s\n", observatoryHost)

	m := make(map[string]string)
	m["host"] = observatoryHost

	result := ScanObject{}
	err := callObservatory("get", "analyze", &result, m, make(map[string]string))
	if err != nil {
		return result, err
	}

	// If there was an error, return that too
	if result.Error != "" {
		return result, errors.New(result.Error)
	}

	return result, nil
}

// GetObservatoryDetails gets the details of a hostname
func GetObservatoryDetails(scanID int) (map[string]ScanDetail, error) {

	fmt.Printf("Getting scan details for scan: %d\n", scanID)

	m := make(map[string]string)
	m["scan"] = strconv.Itoa(scanID)

	result := map[string]ScanDetail{}
	err := callObservatory("get", "getScanResults", &result, m, make(map[string]string))
	if err != nil {
		return result, err
	}

	return result, nil
}
