# go-observe
🌌 Go-Observe: A command line Mozilla Observatory client written in Go

![Screenshot](https://i.imgur.com/avK1Mja.png)

## usage:

```
» ./go-observe-0.2-darwin-amd64

                      __
  ___ ____  _______  / /  ___ ___ _____  _____
 / _  / _ \/___/ _ \/ _ \(_-</ -_) __/ |/ / -_)
 \_, /\___/    \___/_.__/___/\__/_/  |___/\__/
/___/            @leonjza | v0.2

Usage:
 ./go-observe-0.2-darwin-amd64 [command] [<args>]

Available Commands:
	results 		- get analysis results for a url
	bulkresults 		- get bulk analysis for urls from a file
	submit 			- submit a url for analysis
	bulksubmit		- submit urls for analysis read from a file
	version 		- print the version and exit

Examples:
	./go-observe-0.2-darwin-amd64 version
	./go-observe-0.2-darwin-amd64 results -url https://www.google.com
	./go-observe-0.2-darwin-amd64 results -url https://www.google.com --detail
	./go-observe-0.2-darwin-amd64 submit -url https://www.google.com
	./go-observe-0.2-darwin-amd64 submit -url https://www.google.com --rescan
	./go-observe-0.2-darwin-amd64 bulksubmit -file urls.txt
```

## install
Download the release binaries [here](https://github.com/leonjza/go-observe/releases), or, build from source!

```
$ git clone https://github.com/leonjza/go-observe.git
$ cd go-observe
$ go build -o go-observe *.go 
$ ./go-observe
```
