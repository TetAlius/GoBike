package frontend

import (
	"appengine"
	"appengine/blobstore"
	"github.com/TetAlius/GoBike/backend"
	"github.com/TetAlius/GoBike/backend/maped"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
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
	//t.Execute(w, map[string]string{"ID": "Hola k ase?"})
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
	road, _ := strconv.ParseBool(r.FormValue("road"))
	distance, _ := strconv.ParseFloat(r.FormValue("distance"), 64)
	mountain, _ := strconv.ParseBool(r.FormValue("mountain"))
	path, _ := strconv.ParseBool(r.FormValue("path"))
	signal, _ := strconv.ParseBool(r.FormValue("signal"))
	beginTransport, _ := strconv.ParseBool(r.FormValue("signal"))
	garage, _ := strconv.ParseBool(r.FormValue("garage"))
	totalAscent, _ := strconv.ParseFloat(r.FormValue("totalAscent"), 64)

	g := maped.Route{
		Title:          r.FormValue("title"),
		Description:    r.FormValue("description"),
		CreationDate:   time.Now(),
		Distance:       distance,
		BeginLoc:       r.FormValue("beginLoc"),
		EndLoc:         r.FormValue("endLoc"),
		Difficulty:     r.FormValue("difficulty"),
		Road:           road,
		Mountain:       mountain,
		Path:           path,
		Comments:       []string{""},
		Author:         "Menti",
		Maps:           r.FormValue("blobKey"),
		Duration:       time.Now(), //change this to 3 or another int
		Slope:          -12,
		Photos:         []string{string(file[0].BlobKey)},
		Score:          "4",
		Signal:         signal,
		BeginTransport: beginTransport,
		Garage:         garage,
		TotalAscent:    totalAscent,
	}

	http.Redirect(w, r, "/", http.StatusFound)
	backend.AddRoute(g, c)
}
