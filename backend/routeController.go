package backend

import (
	"errors"

	"backend/maped"

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
	// The string "default_guestbook" here could be varied to have multiple guestbooks.
	return datastore.NewKey(c, "Routes", "default_route", 0, nil)
}
