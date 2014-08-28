package backend

import (
	"errors"

	"appengine/datastore"
)

//GetAllRoutes Returns all the routes in the DB
func GetAllRoutes(context Context) (routes []Route, err error) {
	query := datastore.NewQuery("routes")
	_, err := query.GetAll(context, routes)
	if error != nil {
		context.Errorf("Can´t load the routes: %e", err)
		err = errors.New("Can't load the routes")
	}
	return
}

//TestConnection Test connection
func TestConnection() string {
	s := "Hola ke ase"
	return s
}
