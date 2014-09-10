package backend

import (
	"appengine"
	"appengine/datastore"
	"errors"
	"github.com/TetAlius/GoBike/backend/maped"
	"net/http"
	"strconv"
	"time"
)

//GetAllRoutes Returns all the routes in the DB
func GetAllRoutes(context appengine.Context) (routes maped.Routes, err error) {
	query := datastore.NewQuery("Routes").Ancestor(routeKey(context))
	//routes = make([]maped.Routes, 0, 10)
	//routes = maped.Routes

	var routesTMP []maped.Route
	_, err = query.GetAll(context, &routesTMP)
	if err != nil {
		context.Errorf("Can´t load the routes: %s", err)
		err = errors.New("Can't load the routes")
	}

	routes = make(map[int]maped.Route)
	for pos, routeTMP := range routesTMP {
		routes[pos] = routeTMP
	}

	err = filterDistance(context, "40", "600", routes)
	return
}

func filterDistance(context appengine.Context, distanceMin string, distanceMax string, routes maped.Routes) (err error) {
	context.Infof("filterDistance start")
	min, _ := strconv.ParseFloat(distanceMin, 64)
	max, _ := strconv.ParseFloat(distanceMax, 64)

	for pos, route := range routes {
		context.Infof("pos: %s, route: %s", pos, routes)
		if route.Distance < min || route.Distance > max {
			context.Infof("min -> %s < %s", route.Distance, min)
			context.Infof("max -> %s > %s", route.Distance, max)
			delete(routes, pos)
		}
	}
	return
}

func filterDifficulty(context appengine.Context, difficulty string, routes maped.Routes) (err error) {
	for pos, route := range routes {
		if route.Difficulty != difficulty {
			delete(routes, pos)
		}
	}
	return
}

func filterTypeRoad(context appengine.Context, road bool, mountain bool, path bool, routes maped.Routes) (err error) {
	for pos, route := range routes {
		switch {
		case (route.Road != road) && (route.Mountain != mountain) && (route.Path != path):
			delete(routes, pos)
		default:
		}
	}
	return
}

/* uncomment when Duration is a int
func filterDuration(context appengine.Context, durationString string, routes maped.Routes) (err error) {
	duration, _ := strconv.Atoi(durationString)
	for pos, route := range routes {
		if route.Duration < duration-1 || route.Duration > duration+1 {
			delete(routes, pos)
		}
	}
	return

}*/

/*
func filterSlope(comparison string, slope string, routes []maped.Route) (err error) {

}

func filterTotalAscent(comparison string, totalAscent string, routes []maped.Route) (err error) {

}

func filterScore(comparison string, score string, routes []maped.Route) (err error) {

}

func filterSignal(comparison string, signal string, routes []maped.Route) (err error) {

}

func filterBeginTransport(comparison string, beginTransport string, routes []maped.Route) (err error) {

}

func filterGarage(comparison string, garage string, routes []maped.Route) (err error) {

}
*/

func routeKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Routes", "default_route", 0, nil)
}

// InsertRoutesHandler the routes handler
func InsertRoutesHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	g := maped.Route{
		Title:          "Salas - Pola de Allande",
		Description:    "Pues ni puta idea, porque cuando lelgamos al puticlub \"Nenas\" nos volvimos con la rodilla debajo del brazo",
		CreationDate:   time.Now(),
		Distance:       52,
		BeginLoc:       "Salas",
		EndLoc:         "Pola de Allande",
		Difficulty:     "Depende de tu rodilla",
		Road:           true,
		Mountain:       true,
		Path:           false,
		Comments:       []string{"Mola pila", "Habia gastroenteritis", "Rompi la rodilla", "No sabia que los paragüayos hablaban"},
		Author:         "Menti",
		Maps:           "mira como mola __-/^^^^^^^\\____",
		Duration:       time.Now(), //change this to 3 or another int
		Slope:          1200,
		Photos:         "nah",
		Score:          "over 9000",
		Signal:         true,
		BeginTransport: true,
	}
	key := datastore.NewIncompleteKey(c, "Routes", routeKey(c))
	_, err := datastore.Put(c, key, &g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
