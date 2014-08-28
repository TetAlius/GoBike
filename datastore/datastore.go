package datastore

import (
	"time"
)

// Route the routes
type Route struct {
	title          string
	description    string
	creationDate   time.Time
	distance       float64
	beginLoc       string
	endLoc         string
	difficulty     string
	road           bool
	mountain       bool
	path           bool
	comments       []string
	author         string
	maps           string
	duration       time.Time
	slope          float64
	photos         string
	score          string
	signal         bool
	beginTransport bool
}
