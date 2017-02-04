package main

// A small client library for the Mozzilla HTTP Observatory
// https://observatory.mozilla.org/

// 2017 - @leonjza
// Weekend Projects FTW

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var version = "0.1"

type scanObject struct {
	// If an error occured, the Error key will be set. This is how
	// we will know if shit went south.
	Error string `json:"error"`

	// Values from:
	//	https://github.com/mozilla/http-observatory/blob/master/httpobs/docs/api.md#scan
	EndTime             string            `json:"end_time"`
	Grade               string            `json:"grade"`
	Hidden              bool              `json:"hidden"`
	LikelihoodIndicator string            `json:"likelihood_indicator"`
	ResponseHeaders     map[string]string `json:"response_headers"`
	ScanID              int               `json:"scan_id"`
	Score               int               `json:"score"`
	StartTime           string            `json:"start_time"`
	State               string            `json:"state"`
	TestsFailed         int               `json:"tests_failed"`
	TestsPassed         int               `json:"tests_passed"`
	TestsQuantity       int               `json:"tests_quantity"`
}

type scanDetail struct {
	Expectation      string            `json:"expectation"`
	Name             string            `json:"name"`
	Output           map[string]string `json:"output"`
	Pass             bool              `json:"pass"`
	Result           string            `json:"result"`
	ScoreDescription string            `json:"score_description"`
	ScoreModifier    int               `json:"score_modifier"`
}

func validateAndGetURLHost(urlSample string) (string, error) {

	parsed, err := url.ParseRequestURI(urlSample)

	if err != nil {
		fmt.Printf("Unable to parse URL '%s' to get the host\n", urlSample)
		return "", err
	}

	// return parsed.Scheme + "://" + parsed.Host
	return parsed.Host, nil
}

func callObservatory(method string, endpoint string, target interface{}, queryString map[string]string, requestBody map[string]string) {

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

func printVersion() {

	fmt.Printf("%s\n", version)
}

func main() {

	flag.Usage = func() {

		fmt.Printf("go-observe v%s\n", version)
		fmt.Println("Usage:")
		fmt.Printf(" %s [command] [<args>] \n", os.Args[0])
		fmt.Println("")

		fmt.Println("Available Commands:")
		fmt.Println("	results 	- get analysis results for a url")
		fmt.Println("	submit 		- submit a url for analysis")
		fmt.Println("	version 	- print the version and exit")
		fmt.Println("")

		fmt.Println("Examples:")
		fmt.Printf("	%s version\n", os.Args[0])
		fmt.Printf("	%s results -url https://www.google.com\n", os.Args[0])
		fmt.Printf("	%s results -url https://www.google.com\n --detail", os.Args[0])
		fmt.Printf("	%s submit -url https://www.google.com\n", os.Args[0])
		fmt.Printf("	%s submit -url https://www.google.com --rescan\n", os.Args[0])
	}

	// strucutre the application subcommands and flags
	// go-observe version
	versionCommand := flag.NewFlagSet("version", flag.ExitOnError)

	// go-observe results
	resultsCommand := flag.NewFlagSet("results", flag.ExitOnError)
	resultsURLFlag := resultsCommand.String("url", "", "URL to retreive results for")
	resultsDetailFlag := resultsCommand.Bool("detail", false, "Get full test details")

	// go-observe submit
	submitCommand := flag.NewFlagSet("submit", flag.ExitOnError)
	submitURLFlag := submitCommand.String("url", "", "URL to submit an analysis for")
	submitHideFlag := submitCommand.Bool("no-hide", false, "Prevent results from being hidden from the Observatory site")
	submitRescanFlag := submitCommand.Bool("rescan", false, "Should a rescan be forced")

	// generic flags
	versionFlag := flag.Bool("version", false, "Display the version number and exit")

	// Parse the subcommands and flags
	flag.Parse()

	// Get the command intended to run
	command := flag.Arg(0)

	if command == "" {
		flag.Usage()
		return
	}

	// If we have -version or not commands, just print and leave asap
	if *versionFlag {
		printVersion()
		return
	}

	// Ensure the command we got is valid and parse the flags
	switch command {
	case "version":
		{
			versionCommand.Parse(os.Args[2:])
		}
	case "submit":
		{
			submitCommand.Parse(os.Args[2:])
		}
	case "results":
		{
			resultsCommand.Parse(os.Args[2:])
		}
	default:
		{
			fmt.Printf("'%s' is not a valid command\n\n", command)
			return
		}
	}

	// sub command logic
	if versionCommand.Parsed() {
		printVersion()
	}

	if submitCommand.Parsed() {
		if *submitURLFlag == "" {
			fmt.Println("Please provide a URL to submit")
			return
		}

		host, _ := validateAndGetURLHost(*submitURLFlag)
		result := submitObservatoryAnalysis(host, *submitHideFlag, *submitRescanFlag)

		if result.Error != "" {
			fmt.Printf("Error: %s\n", result.Error)
			return
		}

		fmt.Printf("Scan is now:	%s\n", result.State)
	}

	if resultsCommand.Parsed() {
		if *resultsURLFlag == "" {
			fmt.Println("Please provide a URL to retreive results")
			return
		}

		host, _ := validateAndGetURLHost(*resultsURLFlag)
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

		if *resultsDetailFlag {
			details := getObservatoryDetails(result.ScanID)
			fmt.Println("")

			for _, v := range details {

				fmt.Printf("Pass: %t	Score: %d	Description: %s\n", v.Pass, v.ScoreModifier, v.ScoreDescription)
			}
		}
	}
}
