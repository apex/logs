package logs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Error is an error returned by the client.
type Error struct {
	Status     string
	StatusCode int
	Type       string
	Message    string
}

// Error implementation.
func (e Error) Error() string {
	if e.Type == "" {
		return fmt.Sprintf("%s: %d", e.Status, e.StatusCode)
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

// call implementation.
func call(endpoint, method string, in, out interface{}) error {
	var body io.Reader

	// input params
	if in != nil {
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(in)
		if err != nil {
			return fmt.Errorf("encoding: %w", err)
		}
		body = &buf
	}

	// POST request
	req, err := http.NewRequest("POST", endpoint+"/"+method, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// response
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// error
	if res.StatusCode >= 300 {
		var e Error
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return err
		}
		e.Status = http.StatusText(res.StatusCode)
		e.StatusCode = res.StatusCode
		return e
	}

	// output params
	if out != nil {
		err = json.NewDecoder(res.Body).Decode(out)
		if err != nil {
			return err
		}
	}

	return nil
}
