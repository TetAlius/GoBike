package frontend

import (
	"appengine"
	"github.com/TetAlius/GoBike/backend"
	"html/template"
	"log"
	"net/http"
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

func serve(w http.ResponseWriter, r *http.Request) {
        blobstore.Send(w, appengine.BlobKey(r.FormValue("blobKey")))
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
        http.Redirect(w, r, "/serve/?blobKey="+string(file[0].BlobKey), http.StatusFound)
}
