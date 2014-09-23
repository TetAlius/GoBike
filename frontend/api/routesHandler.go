package frontend

import (
	"appengine"
	"github.com/TetAlius/GoBike/backend"
	"github.com/TetAlius/GoBike/backend/maped"
	"html/template"
	"log"
	"net/http"
	"time"
	"strconv"
	"appengine/blobstore"
)

func routesHandler(w http.ResponseWriter, r *http.Request) {
	// 404 page
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	c := appengine.NewContext(r)

	t, err := template.ParseFiles("./frontend/resources/html/all-routes.html", "./frontend/resources/html/login.html")
	if err != nil {
		log.Fatal("Parse file error: ", err)
	}

	routes, err := backend.GetAllRoutes(c)
	if err != nil {
		log.Fatal("GetAllRoutes error: ", err)
	}
	t.Execute(w, routes)
}


func insertRoutesHandler(w http.ResponseWriter, r *http.Request) {
	//t, err := template.ParseFiles("./frontend/resources/html/addRoute.html")
	c := appengine.NewContext(r)
	t, err := template.ParseFiles("./frontend/resources/html/addRoute.html", "./frontend/resources/html/meta.html", "./frontend/resources/html/login.html")
	if err != nil {
		log.Fatal("Parse file error: ", err)
	}
 	uploadURL, erro := blobstore.UploadURL(c, "/uploadPhoto", nil)
        if erro != nil {
        	log.Fatal("unable to generate upload url: ", erro)
                return
        }
    t.Execute(w, map[string]string{"PageTitle": "GoBike - Insert Route", 
    			"UploadURL": uploadURL.String()})    
	//t.Execute(w, map[string]string{"PageTitle": "GoBike - Insert Route"})
}

func upload(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        blobs, _, err := blobstore.ParseUpload(r)
        if err != nil {
                return
        }
        file := blobs["file"]
        if len(file) == 0 {
                c.Errorf("no file uploaded")
                http.Redirect(w, r, "/", http.StatusFound)
                return
        }
		g := maped.Route{
			Title:          r.FormValue("title"),
			Description:    r.FormValue("description"),
			CreationDate:   time.Now(),
			Distance:       r.FormValue("distance"),
			BeginLoc:       r.FormValue("beginLoc"),
			EndLoc:         r.FormValue("endLoc"),
			Difficulty:     r.FormValue("difficulty"),
			Road:           strconv.ParseBool(r.FormValue("road")),
			Mountain:       r.FormValue("mountain"),
			Path:           r.FormValue("path"),
			Comments:       []string{""},
			Author:         r.FormValue("blobKey"),
			Maps:           r.FormValue("blobKey"),
			Duration:       time.Now(), //change this to 3 or another int
			Slope:          -12,
			Photos:         []string{string(file[0].BlobKey)},
			Score:          1,
			Signal:         r.FormValue("signal"),
			BeginTransport: r.FormValue("beginTransport"),
			TotalAscent:    r.FormValue("totalAscent"),
		}
	key := datastore.NewIncompleteKey(c, "Routes", routeKey(c))
	_, err := datastore.Put(c, key, &g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
	http.Redirect(w, r, "/", http.StatusFound)
}
