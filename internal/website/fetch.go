package website

import (
	"io"
	"net/http"
	"time"

	"github.com/cojoj/analyzer/internal/models"
)

// Fetch performs a GET request on provided string with timeout of 5 seconds. If response's
// status code is in the range from 200 to 300 it returns it's Body. It's up to the user to close it
// once finished using it.
func Fetch(u string) (io.ReadCloser, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := client.Get(u)
	if err != nil {
		return nil, models.Wrap(err, u, 0)
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		return nil, models.Wrap(nil, u, res.StatusCode)
	}

	return res.Body, nil
}

// Reachable performs HEAD request on provided URL and checks the status code. If it's 200 it'll
// return `true` and status code. In any other case it'll return `false` adn status code 0.
func Reachable(u string) (bool, int) {
	res, err := http.Head(u)
	if err != nil {
		return false, 0
	}

	return res.StatusCode == http.StatusOK, res.StatusCode
}
