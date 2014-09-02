package backend

/*
PRIMERA VERSION DEL LOGIN
FUENTE: http://www.mschoebel.info/2014/03/09/snippet-golang-webapp-login-logout.html




func login(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	pass := request.FormValue("password")
	if checkCredentials(name, pass) {
		setSession(name, response)
		redirectTarget := "/"
		http.Redirect(response, request, redirectTarget, 302)
	}
}

func checkCredentials(name, password string) bool {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("User").
		Filter("Name =", name).
		Filter("Password =", password)

	var user []User
	usr, err := q.GetAll(c, &user)

	if err != nil {
		context.Errorf("Can't find user: %e", err)
		err = errors.New("Can't find user")
		return false
	}

	if usr != nil {
		return true
	}
	return false

}

func logout() {
	clearSession(response)
	http.Redirect(response, request, "/", 302)
}

func setSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}
*/
