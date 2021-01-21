package models

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