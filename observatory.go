package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func callObservatory(method string, endpoint string, target interface{}, queryString map[string]string,
	requestBody map[string]string) {

	observatoryBase := "https://http-observatory.security.mozilla.org/api/v1/"

	// prepare the final url we will call
	u, err := url.Parse(observatoryBase + endpoint)

	if err != nil {
		fmt.Println("Somehow failed to parse base Observatory URL")
		fmt.Println(err)

		os.Exit(2)
	}

	// prepare to add the query string parameters
	params := url.Values{}
	for k, v := range queryString {
		params.Add(k, v)
	}

	// slap the query string on to u
	u.RawQuery = params.Encode()
	fmt.Printf("[%s] %s\n", method, u.String())

	// prepare to add the body values
	bodyData := url.Values{}
	for k, v := range requestBody {
		bodyData.Add(k, v)
	}

	// build the request
	req, err := http.NewRequest(method, u.String(), strings.NewReader(bodyData.Encode()))

	// If we are using the post method, set the content type
	// as form-encoded so that the body could be understood on the
	// other end.
	if method == "post" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	if err != nil {
		fmt.Println("New Request:", err)
		return
	}

	// make the request
	client := http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Do: ", err)
		return
	}

	defer resp.Body.Close()

	// marshal the response into the struct type at target
	json.NewDecoder(resp.Body).Decode(&target)
}

// Submit a host for analysis.
func submitObservatoryAnalysis(observatoryHost string, nohide bool, rescan bool) scanObject {

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

	results := scanObject{}
	callObservatory("post", "analyze", &results, queryString, requestBody)

	return results
}

// Get analysis results
func getObservatoryResults(observatoryHost string) scanObject {

	fmt.Printf("Getting results for: %s\n", observatoryHost)

	m := make(map[string]string)
	m["host"] = observatoryHost

	result := scanObject{}
	callObservatory("get", "analyze", &result, m, make(map[string]string))

	return result
}

func getObservatoryDetails(scanID int) map[string]scanDetail {

	fmt.Printf("Getting scan details for scan: %d\n", scanID)

	m := make(map[string]string)
	m["scan"] = strconv.Itoa(scanID)

	result := map[string]scanDetail{}
	callObservatory("get", "getScanResults", &result, m, make(map[string]string))

	return result
}
