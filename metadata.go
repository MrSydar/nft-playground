package main

type metadata struct {
	Name        string         `json:"name"`
	Descriction string         `json:"description"`
	Image       string         `json:"image"`
	Attributes  map[string]any `json:"attributes"`
}
