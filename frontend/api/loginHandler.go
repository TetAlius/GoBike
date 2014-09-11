package frontend

import (
	//	"fmt"
	"net/http"

	"appengine"
	"encoding/base64"
	"encoding/hex"

	"github.com/TetAlius/GoBike/backend"
	//"github.com/TetAlius/GoBike/backend/maped"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	//username := r.FormValue("username")
	username := r.FormValue("username")
	password := r.FormValue("password")
	password = hex.EncodeToString([]byte(base64.StdEncoding.EncodeToString([]byte(password))))
	if result := backend.Login(c, w, username, password); result {
		c.Infof("loged as: %s", username)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		c.Infof("not logged")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
