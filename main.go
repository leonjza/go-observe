package main

// A small client library for the Mozzilla HTTP Observatory
// https://observatory.mozilla.org/

// 2017 - @leonjza
// Weekend Projects FTW

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// really, all we using this for is to parse the command line
	// and move on to the commands file for actual processing

	flag.Usage = func() {
		fmt.Printf(`
                      __
  ___ ____  _______  / /  ___ ___ _____  _____
 / _  / _ \/___/ _ \/ _ \(_-</ -_) __/ |/ / -_)
 \_, /\___/    \___/_.__/___/\__/_/  |___/\__/
/___/            @leonjza | v%s`, version)

		fmt.Printf("\n\n")
		fmt.Println("Usage:")
		fmt.Printf(" %s [command] [<args>] \n", os.Args[0])
		fmt.Println("")

		fmt.Println("Available Commands:")
		fmt.Println("	results 		- get analysis results for a url")
		fmt.Println("	bulkresults 		- get bulk analysis for urls from a file")
		fmt.Println("	submit 			- submit a url for analysis")
		fmt.Println("	bulksubmit		- submit urls for analysis read from a file")
		fmt.Println("	version 		- print the version and exit")
		fmt.Println("")

		fmt.Println("Examples:")
		fmt.Printf("	%s version\n", os.Args[0])
		fmt.Printf("	%s results -url https://www.google.com\n", os.Args[0])
		fmt.Printf("	%s results -url https://www.google.com --detail\n", os.Args[0])
		fmt.Printf("	%s submit -url https://www.google.com\n", os.Args[0])
		fmt.Printf("	%s submit -url https://www.google.com --rescan\n", os.Args[0])
		fmt.Printf("	%s bulksubmit -file urls.txt\n", os.Args[0])
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

	// go-observe bulksubmit
	bulkSubmitCommand := flag.NewFlagSet("bulksubmit", flag.ExitOnError)
	bulkSubmitFileNameFlag := bulkSubmitCommand.String("file", "", "The file to open with URLs for submission")

	// go-observe bulkresults
	bulkResultsCommand := flag.NewFlagSet("bulkresults", flag.ExitOnError)
	bulkResultsFileNameFlag := bulkResultsCommand.String("file", "", "The file to read for URLs to get results for")

	flag.Parse()

	// get the command intended to run
	command := flag.Arg(0)
	// and the rest of the arguments
	args := os.Args[2:]

	if command == "" {
		flag.Usage()
		return
	}

	// parse commandline flags based on the command we are running
	switch command {
	case "version":
		{
			versionCommand.Parse(args)
		}
	case "submit":
		{
			submitCommand.Parse(args)
		}
	case "bulksubmit":
		{
			bulkSubmitCommand.Parse(args)
		}
	case "results":
		{
			resultsCommand.Parse(args)
		}
	case "bulkresults":
		{
			bulkResultsCommand.Parse(args)
		}
	default:
		{
			fmt.Printf("'%s' is not a valid command", command)
			flag.Usage()
			return
		}
	}

	// exec the subcommand if the flags managed to parse
	if versionCommand.Parsed() {

		printVersion()
	}

	if submitCommand.Parsed() {

		execSubmitCommand(*submitURLFlag, *submitHideFlag, *submitRescanFlag)
	}

	if bulkSubmitCommand.Parsed() {

		execBulkSubmitCommand(*bulkSubmitFileNameFlag)
	}

	if resultsCommand.Parsed() {

		execResultsCommand(*resultsURLFlag, *resultsDetailFlag)
	}

	if bulkResultsCommand.Parsed() {

		execBulkResultsCommand(*bulkResultsFileNameFlag)
	}
}
