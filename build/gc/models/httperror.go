package models

import (
	"encoding/json"
	"fmt"
)

type HttpStatusError struct {
	Verb       string
	Path       string
	Body       string
	StatusCode int
	Headers    map[string][]string
}

func (e HttpStatusError) Error() string {
	return e.Body
}

func (e *HttpStatusError) ErrorDescriptive() string {
	headers, _ := json.Marshal(e.Headers)
	return fmt.Sprintf("Verb %s, Path %s, Status Code %d, Headers %v", e.Verb, e.Path, e.StatusCode, string(headers))
}
