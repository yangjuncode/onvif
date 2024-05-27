package networking

import (
	"bytes"
	"net/http"

	"github.com/juju/errors"
)

// SendSoap send soap message
func SendSoap(httpClient *http.Client, endpoint, message string) (*http.Response, error) {
	resp, err := httpClient.Post(endpoint, "application/soap+xml; charset=utf-8", bytes.NewBufferString(message))
	if err != nil {
		return resp, errors.Annotate(err, "Post")
	}

	// if resp.StatusCode is 4xx,5xx, return error
	if resp.StatusCode >= 400 && resp.StatusCode < 600 {
		return resp, errors.Errorf("Server error: %d: %s", resp.StatusCode, resp.Status)
	}

	return resp, nil
}
