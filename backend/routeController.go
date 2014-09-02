package backend

import (
	"errors"
	"net/http"
	"time"

	"github.com/TetAlius/GoBike/backend/maped"

	"appengine"
	"appengine/datastore"
)

//GetAllRoutes Returns all the routes in the DB
func GetAllRoutes(context appengine.Context) (routes []maped.Route, err error) {
	query := datastore.NewQuery("Routes").Ancestor(routeKey(context))
	routes = make([]maped.Route, 0, 10)
	_, err = query.GetAll(context, &routes)
	if err != nil {
		context.Errorf("Can´t load the routes: %e", err)
		err = errors.New("Can't load the routes")
	}
	return
}

func routeKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Routes", "default_route", 0, nil)
}

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
		Duration:       time.Now(),
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
