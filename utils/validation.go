package utils

import (
	"fmt"
	"net/url"
	"strings"
)

func ValidateAndGetURLHost(urlSample string) (string, error) {

	// cheat a little by checking if we have :// in the string.
	// if we dont, its probably a hostname already
	if !strings.Contains("://", urlSample) {

		s := strings.TrimSpace(urlSample)
		return s, nil
	}

	// otherwise, if there is a ://, try and extract the hostname
	parsed, err := url.ParseRequestURI(urlSample)

	if err != nil {
		fmt.Printf("Unable to parse URL '%s' to get the host\n", urlSample)
		return "", err
	}

	// return parsed.Scheme + "://" + parsed.Host
	return parsed.Host, nil
}
