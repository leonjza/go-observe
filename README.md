# go-observe
🌌 Go-Observe: A command line Mozilla Observatory client written in Go

![Screenshot](https://i.imgur.com/PMhaFdu.png)

## usage:

```
~ » go-observe
                      __
  ___ ____  _______  / /  ___ ___ _____  _____
 / _  / _ \/___/ _ \/ _ \(_-</ -_) __/ |/ / -_)
 \_, /\___/    \___/_.__/___/\__/_/  |___/\__/
/___/ @leonjza

Usage:
  go-observe [command]

Available Commands:
  fileresult  Get scan results for URLs / Hostnames in a file
  filesubmit  Submit URLs / Hostnames from a file to the Observatory for analysis
  result      Retrieve results for a URL / Hostname from the Observatory
  submit      Submit a URL / Hostname to the Observatory for analysis
  version     Show the version of go-observe

Use "go-observe [command] --help" for more information about a command.
```

## install
There are many options to get `go-observe`.

Download the release binaries [here](https://github.com/leonjza/go-observe/releases)

*or*

Download with `go get` if you have `$GOPATH` setup: `go get github.com/leonjza/go-observe`

*or*

Build from source!
```
$ git clone https://github.com/leonjza/go-observe.git
$ cd go-observe
$ go build -o go-observe *.go 
$ ./go-observe
```
