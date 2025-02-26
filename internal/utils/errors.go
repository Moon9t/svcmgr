package utils

import (
	"errors"
	"fmt"
)

var ErrUnsupportedService = errors.New("unsupported service")

// HumanizeError converts an error into a human-readable string.
func HumanizeError(err error) string {
	if err == nil {
		return "No error"
	}
	return fmt.Sprintf("An error occurred: %v", err)
}
