package models

import (
	"encoding/json"
)

// Entities is a generic structure containing entity listing, cursor and index based paginatable responses
type Entities struct {
	Entities      []json.RawMessage `json:"entities"`
	Resources     []json.RawMessage `json:"Resources"`
	Conversations []json.RawMessage `json:"conversations"`
	PageSize      int               `json:"pageSize"`
	PageNumber    int               `json:"pageNumber"`
	Total         int               `json:"total"`
	PageCount     int               `json:"pageCount"`
	NextUri       string            `json:"nextUri"`
	StartIndex    int               `json:"startIndex"`
	TotalResults  int               `json:"totalResults"`
	Cursor        string            `json:"cursor"`
	Cursors       Cursors           `json:"cursors"`
}

type Cursors struct {
	After string `json:"after"`
}