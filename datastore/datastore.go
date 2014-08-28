package datastore

import (
	"datastore/maped"
)


func getAllRoutes(context Context) []Route {
	query := datastore.NewQuery("routes")
	var routes []Route
	_, error := query.GetAll(context, routes)
	if error == nil {
		return routes
	}

}

