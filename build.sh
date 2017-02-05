#!/usr/bin/env sh

# get the version
if [ -z "$1" ]
then
    echo "Provide the version!"
    exit
fi

VERSION=$1

# build references:
# https://golang.org/doc/install/source#environment

# Linux builds
FILENAME="go-observe-$VERSION-linux-386"
echo "Compiling 386 Linux to $FILENAME"
GOOS=linux GOARCH=386 go build  -o build/$FILENAME *.go

FILENAME="go-observe-$VERSION-linux-amd64"
echo "Compiling amd64 Linux to $FILENAME"
GOOS=linux GOARCH=amd64 go build  -o build/$FILENAME *.go

FILENAME="go-observe-$VERSION-linux-arm"
echo "Compiling arm Linux to $FILENAME"
GOOS=linux GOARCH=arm go build  -o build/$FILENAME *.go

FILENAME="go-observe-$VERSION-linux-arm64"
echo "Compiling arm64 Linux to $FILENAME"
GOOS=linux GOARCH=arm go build  -o build/$FILENAME *.go

# Windows build
FILENAME="go-observe-$VERSION-windows-386.exe"
echo "Compiling 386 Windows to $FILENAME"
GOOS=windows GOARCH=386 go build  -o build/$FILENAME *.go

FILENAME="go-observe-$VERSION-windows-amd64.exe"
echo "Compiling amd64 Windows to $FILENAME"
GOOS=windows GOARCH=amd64 go build  -o build/$FILENAME *.go

# macOS build
FILENAME="go-observe-$VERSION-darwin-386"
echo "Compiling 386 macOS to $FILENAME"
GOOS=darwin GOARCH=386 go build  -o build/$FILENAME *.go

FILENAME="go-observe-$VERSION-darwin-amd64"
echo "Compiling amd64 macOS to $FILENAME"
GOOS=darwin GOARCH=amd64 go build  -o build/$FILENAME *.go
