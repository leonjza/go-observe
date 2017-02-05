package main

import (
	"fmt"
	"net/url"
)

func validateAndGetURLHost(urlSample string) (string, error) {

	parsed, err := url.ParseRequestURI(urlSample)

	if err != nil {
		fmt.Printf("Unable to parse URL '%s' to get the host\n", urlSample)
		return "", err
	}

	// return parsed.Scheme + "://" + parsed.Host
	return parsed.Host, nil
}

func printVersion() {

	fmt.Printf("%s\n", version)
}
