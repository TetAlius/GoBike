package maped

import (
	"time"
)

// Route the routes
type Route struct {
	Title          string
	Description    string
	CreationDate   time.Time
	Distance       float64 // filter
	BeginLoc       string
	EndLoc         string
	Difficulty     string // filter
	Road           bool   // filter
	Mountain       bool   // filter
	Path           bool   // filter
	Comments       []string
	Author         string
	Maps           string
	Duration       int     // filter
	Slope          float64 // filter
	TotalAscent    float64 // filter
	Photos         string
	Score          string // filter
	Signal         bool   // filer
	BeginTransport bool   // filter
	Garage         bool   // filter

}
