package gogovtrack

import "fmt"

// ErrBadStatusCode is returned when the API returns a non 200 error code
type ErrBadStatusCode struct {
	OriginalBody string
	Code         int
}

func (e *ErrBadStatusCode) Error() string {
	return fmt.Sprintf("Invalid status code: %d", e.Code)
}

// ErrNotExpectedJSON is returned by API calls when the response isn't expected JSON
type ErrNotExpectedJSON struct {
	OriginalBody string
	Err          error
}

func (e *ErrNotExpectedJSON) Error() string {
	return fmt.Sprintf("Unexpected JSON: %s from %s", e.Err.Error(), e.OriginalBody)
}
