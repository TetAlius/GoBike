package frontend

import (
	//	"fmt"
	"net/http"

	"encoding/base64"
	"encoding/hex"

	//"github.com/TetAlius/GoBike/backend"
	//"github.com/TetAlius/GoBike/backend/maped"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {

	//username := r.FormValue("username")
	_ = r.FormValue("username")
	password := r.FormValue("password")
	password = hex.EncodeToString([]byte(base64.StdEncoding.EncodeToString([]byte(password))))
	//Create and insert the hash for the user
	/*
		if authenticated := backend.CheckCredentials(username, password); !authenticated {
			http.Redirect(w, r, "/", http.StatusFound)
		}
	*/
}
