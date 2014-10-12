package backend

import (
	"appengine"
	"appengine/datastore"
	"errors"
	"github.com/TetAlius/GoBike/backend/maped"
	"strconv"
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

	//err = filterDistance(context, "40", "600", routes)
	//err = filterScore(context, "4", routes)
	//err = filterSlope(context, "-15", "80", routes)
	//err = filterTotalAscent(context, "200", routes)
	//err = filterfilterDifficulty(context, "alta", routes)

	//err = filterSignal(context, false, routes)
	//err = filterGarage(context, true, routes)
	//err = filterBeginTransport(context, false, routes)
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

func filterSlope(context appengine.Context, slopeMin string, slopeMax string, routes maped.Routes) (err error) {
	context.Infof("filterSlope start")
	if len(slopeMin) > 0 {
		min, _ := strconv.ParseFloat(slopeMin, 64)
		max, _ := strconv.ParseFloat(slopeMax, 64)
		for pos, route := range routes {
			context.Infof("pos: %s, route: %s", pos, routes)
			if route.Slope < min || route.Slope > max {
				context.Infof("min -> %s < %s", route.Slope, min)
				context.Infof("max -> %s > %s", route.Slope, max)
				delete(routes, pos)
			}
		}
	} else {
		min, _ := strconv.ParseFloat("0", 64)
		max, _ := strconv.ParseFloat(slopeMax, 64)
		for pos, route := range routes {
			context.Infof("pos: %s, route: %s", pos, routes)
			if route.Slope < min || route.Slope > max {
				context.Infof("min -> %s < %s", route.Slope, min)
				context.Infof("max -> %s > %s", route.Slope, max)
				delete(routes, pos)
			}
		}

	}
	return
}

func filterTotalAscent(context appengine.Context, maxAscent string, routes maped.Routes) (err error) {
	context.Infof("filterTotalAscent start")
	max, _ := strconv.ParseFloat(maxAscent, 64)
	for pos, route := range routes {
		context.Infof("pos: %s, route: %s", pos, routes)
		if route.TotalAscent > max {
			context.Infof("max -> %s > %s", route.TotalAscent, max)
			delete(routes, pos)
		}
	}
	return
}

func filterScore(context appengine.Context, score string, routes maped.Routes) (err error) {
	context.Infof("filterScore start")
	for pos, route := range routes {
		context.Infof("score: %s", score)
		if route.Score != score {
			delete(routes, pos)
		}
	}
	return
}

func filterSignal(context appengine.Context, signal bool, routes maped.Routes) (err error) {
	context.Infof("filterSignal start")
	for pos, route := range routes {
		context.Infof("signal: %s", signal)
		if route.Signal != signal {
			delete(routes, pos)
		}
	}
	return
}

func filterBeginTransport(context appengine.Context, isTransport bool, routes maped.Routes) (err error) {
	context.Infof("filterBeginTransport start")
	for pos, route := range routes {
		if route.BeginTransport != isTransport {
			delete(routes, pos)
		}
	}
	return
}
func filterGarage(context appengine.Context, garage bool, routes maped.Routes) (err error) {
	context.Infof("filterGarage start")
	for pos, route := range routes {
		if route.Garage != garage {
			delete(routes, pos)
		}
	}
	return
}

//AddRoute adds a new Route
func AddRoute(route maped.Route, context appengine.Context) {
	key := datastore.NewIncompleteKey(context, "Routes", routeKey(context))
	route.ID = key.Enconde()
	_, err := datastore.Put(context, key, &route)
	if err != nil {
		return
	}
}

func routeKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Routes", "default_route", 0, nil)
}
