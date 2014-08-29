package backend

import (
	"errors"

	"backend/maped"

	"appengine"
	"appengine/datastore"
)

//GetAllRoutes Returns all the routes in the DB
func GetAllRoutes(context appengine.Context) (routes []maped.Route, err error) {
	query := datastore.NewQuery("routes")
	_, err = query.GetAll(context, routes)
	if err != nil {
		context.Errorf("Can´t load the routes: %e", err)
		err = errors.New("Can't load the routes")
	}
	return
}
