package TMCHttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Sends an HTTP(S) request to whatever endpoint specified.
func SendHTTPReq(endpoint string, request_type string, post_data map[string]string, headers map[string]string, form bool) (string, int, error) {
	// Initialize data and return code (status code).
	d := ""
	rc := -1
	var req *http.Request
	var err error

	var post_body_form io.Reader
	var post_body_json io.Reader

	// Check to see if we need to send post data.
	if request_type == "POST" || request_type == "PUT" {
		// Convert to JSON and use as body.'
		var j []byte
		var err error

		if !form {
			j, err = json.Marshal(post_data)

			if err != nil {
				return d, rc, err
			}
		} else {
			urlParams := url.Values{}

			for key, val := range post_data {
				urlParams.Set(key, val)
			}

			j = []byte(urlParams.Encode())
		}

		// Read byte array into IO reader.
		if form {
			post_body_form = strings.NewReader(string(j))
		} else {
			post_body_json = bytes.NewBuffer(j)
		}

		if err != nil {
			return d, rc, err
		}
	}

	// Setup HTTP request.
	client := &http.Client{Timeout: time.Second * 5}
	if form {
		req, err = http.NewRequest(request_type, endpoint, post_body_form)
	} else {
		req, err = http.NewRequest(request_type, endpoint, post_body_json)
	}

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
