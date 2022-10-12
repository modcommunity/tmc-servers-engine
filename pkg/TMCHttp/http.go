package TMCHttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// Sends an HTTP(S) request to whatever endpoint specified.
func SendHTTPReq(endpoint string, request_type string, post_data map[string]string, headers map[string]string) (string, int, error) {
	// Initialize data and return code (status code).
	d := ""
	rc := -1

	var post_body io.Reader

	// Check to see if we need to send post data.
	if request_type == "POST" || request_type == "PUT" {
		// Convert to JSON and use as body.
		j, err := json.Marshal(post_data)

		if err != nil {
			return d, rc, err
		}

		// Read byte array into IO reader.
		post_body = bytes.NewBuffer(j)

		if err != nil {
			return d, rc, err
		}
	}

	// Setup HTTP request.
	client := &http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest(request_type, endpoint, post_body)

	// Check for error.
	if err != nil {
		fmt.Println(err)

		return d, rc, err
	}

	// Add headers.
	for key, val := range headers {
		req.Header.Set(key, val)
	}

	// Perform HTTP request and check for errors.
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)

		return d, rc, err
	}

	// Set return code.
	rc = resp.StatusCode

	// Read body.
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)

		return d, rc, err
	}

	// Return data as a string.
	d = string(body)

	return d, rc, nil
}
