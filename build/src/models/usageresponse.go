package models

type UsageSubmitResponse struct {
	ExecutionID string `json:"executionId"`
	ResultsURI  string `json:"resultsUri"`
}

type UsageQueryResponse struct {
	QueryStatus   string `json:"queryStatus,omitempty"`
}