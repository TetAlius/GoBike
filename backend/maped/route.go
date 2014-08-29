package maped

import (
	"time"
)

// Route the routes
type Route struct {
	Title          string
	Description    string
	CreationDate   time.Time
	Distance       float64
	BeginLoc       string
	EndLoc         string
	Difficulty     string
	Road           bool
	Mountain       bool
	Path           bool
	Comments       []string
	Author         string
	Maps           string
	Duration       time.Time
	Slope          float64
	Photos         string
	Score          string
	Signal         bool
	BeginTransport bool
}
