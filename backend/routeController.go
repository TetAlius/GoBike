package backend

import (
	"errors"
	"github.com/TetAlius/GoBike/backend/maped"
	"net/http"
	"strconv"
	"time"

	"appengine"
	"appengine/datastore"
)

//GetAllRoutes Returns all the routes in the DB
func GetAllRoutes(context appengine.Context) (routes []maped.Route, err error) {
	query := datastore.NewQuery("Routes").Ancestor(routeKey(context))
	routes = make([]maped.Route, 0, 10)
	_, err = query.GetAll(context, &routes)
	if err != nil {
		context.Errorf("Can´t load the routes: %s", err)
		err = errors.New("Can't load the routes")
	}

	return
}

func filterDistance(distanceMin string, distanceMax string, routes maped.Routes) (err error) {
	min, _ := strconv.ParseFloat(distanceMin, 64)
	max, _ := strconv.ParseFloat(distanceMax, 64)
	for pos, route := range routes {
		if route.Distance < min && route.Distance > max {
			routes.Delete(pos)
		}
	}
	return
}

func filterDifficulty(difficulty string, routes maped.Routes) (err error) {
	for pos, route := range routes {
		if route.Difficulty != difficulty {
			routes.Delete(pos)
		}
	}
	return
}

func filterTypeRoad(road bool, mountain bool, path bool, routes maped.Routes) (err error) {
	for pos, route := range routes {
		switch {
		case (route.Road != road) && (route.Mountain != mountain) && (route.Path != path):
			routes.Delete(pos)
		default:
		}
	}
	return
}

func filterDuration(durationString string, routes maped.Routes) (err error) {
	duration, _ := strconv.Atoi(durationString)
	for pos, route := range routes {
		if route.Duration < duration-1 || route.Duration > duration+1 {
			routes.Delete(pos)
		}
	}
	return

}

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
		Duration:       3,
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
