package models

import "net/http"

type Parameters struct {
	Name        string `json:"name,omitempty"`
	In          string `json:"in,omitempty"`
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Type        string `json:"type,omitempty"`
}

type Method struct {
	Summary     string       `json:"summary,omitempty"`
	Description string       `json:"description,omitempty"`
	Parameters  []Parameters `json:"parameters,omitempty"`
}

type Path struct {
	Get    Method `json:"get,omitempty"`
	Put    Method `json:"put,omitempty"`
	Post   Method `json:"post,omitempty"`
	Patch  Method `json:"patch,omitempty"`
	Delete Method `json:"delete,omitempty"`
}

func (p Path) GetMethod(method string) *Method {
	switch method {
	case http.MethodGet:
		return &p.Get
	case http.MethodPut:
		return &p.Put
	case http.MethodPost:
		return &p.Post
	case http.MethodPatch:
		return &p.Patch
	case http.MethodDelete:
		return &p.Delete
	}

	return nil
}