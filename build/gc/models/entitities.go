package models

import (
	"encoding/json"
)

type Entities struct {
	Entities   []json.RawMessage `json:"entities"`
	PageSize   int               `json:"pageSize"`
	PageNumber int               `json:"pageNumber"`
	Total      int               `json:"total"`
	PageCount  int               `json:"pageCount"`
}
