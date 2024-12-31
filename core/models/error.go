package models

import (
	"fmt"
	"log"
)

// HttpError represents an error returned from an API with a status code greater than or equal to 400.
type HttpError struct {
	Status  int    // HTTP status code
	Message string // Error message
}

func (ce HttpError) Error() string {
	return fmt.Sprintf("HTTP %d - %s", ce.Status, ce.Message)
}

func logError(err error) {
	log.Println("Error occurred:", err)
}

// InternalError represents an error in internal logic.
type InternalError struct {
	Source  string // Error source
	Message string // Error message
}

func (ce InternalError) Error() string {
	return fmt.Sprintf("%s - %s", ce.Source, ce.Message)
}
