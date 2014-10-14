package frontend

import (
	"appengine"
	"github.com/TetAlius/GoBike/backend"
	"github.com/TetAlius/GoBike/backend/maped"
	"net/http"
	"strconv"
	"time"
)

// TODO delete, only for tests
func bulkHandler(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 5; i++ {
		context := appengine.NewContext(r)
		route := maped.Route{
			Title:          "Title: " + strconv.Atoi(i),
			Description:    "Description",
			CreationDate:   time.Now(),
			Distance:       69,
			BeginLoc:       "beginLoc",
			EndLoc:         "endLoc",
			Difficulty:     "difficulty",
			Road:           true,
			Mountain:       true,
			Path:           true,
			Comments:       []string{""},
			Author:         "The Author",
			Maps:           "blobKey",
			Duration:       42,
			Slope:          -12,
			Photos:         []string{"Fotuquis"}, //[]string{string(file[0].BlobKey)},
			Score:          "42",
			Signal:         false,
			BeginTransport: false,
			Garage:         false,
			TotalAscent:    9000,
		}
		http.Redirect(w, r, "/", http.StatusFound)
		backend.AddRoute(route, context)
	}
}
