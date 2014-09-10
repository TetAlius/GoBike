package maped

import (
	"time"
)

// Route a route
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
	Duration       time.Time // filter //must be int, just a workarround
	Slope          float64   // filter
	TotalAscent    float64   // filter
	Photos         string
	Score          string // filter
	Signal         bool   // filer
	BeginTransport bool   // filter
	Garage         bool   // filter

}

// Routes the routes
type Routes map[int]Route

/*
// SliceHandler the routes interface
type SliceHandler interface {
	Delete() error
}

// Delete a given route in a routes slice
func (routes Routes) Delete(context appengine.Context, pos int) (err error) {
	// TODO
	//routes[pos] = routes[len(routes)-1]
	//routes = routes[:len(routes)-1]
	context.Infof("routes[%s] -> %s", pos, routes[pos])
	return
}*/
