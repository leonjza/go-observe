package observatory

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
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
