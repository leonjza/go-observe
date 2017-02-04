# go-observe
🌌 Go-Observe: A command line Mozilla Observatory client written in Go

![Screenshot](https://i.imgur.com/GaL056o.png)

## usage:

```
go-observe v0.1
Usage:
 ./go-observe-0.1-darwin-amd64 [command] [<args>]

Available Commands:
	results 		- get analysis results for a url
	bulkresults 		- get bulk analysis for urls from a file
	submit 			- submit a url for analysis
	bulksubmit		- submit urls for analysis read from a file
	version 		- print the version and exit

Examples:
	./go-observe-0.1-darwin-amd64 version
	./go-observe-0.1-darwin-amd64 results -url https://www.google.com
	./go-observe-0.1-darwin-amd64 submit -url https://www.google.com
	./go-observe-0.1-darwin-amd64 submit -url https://www.google.com --rescan
	./go-observe-0.1-darwin-amd64 bulksubmit -file urls.txt
```
