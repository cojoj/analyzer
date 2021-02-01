package models

import "fmt"

//
type FetchError struct {
	URL        string
	StatusCode int
	Err        error
}

func (e *FetchError) Error() string {
	switch {
	case e.Err != nil:
		return fmt.Sprintf("Fetching %v failed with underlaying error: %v", e.URL, e.Err)
	case e.StatusCode != 0:
		return fmt.Sprintf("Fetching %v failed with status code: %v", e.URL, e.StatusCode)
	default:
		return fmt.Sprintf("Fetching %v failed", e.URL)
	}
}

func Wrap(err error, url string, statusCode int) *FetchError {
	return &FetchError{
		URL:        url,
		StatusCode: statusCode,
		Err:        err,
	}
}
