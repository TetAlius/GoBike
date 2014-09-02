package backend

import (
	"github.com/TetAlius/GoBike/backend/maped"

	"appengine"
	"appengine/datastore"
)

func registerUser(user maped.User) {

}

// AvailableUserToRegister checks if the username or email are being used in a user
func AvailableUserToRegister(context appengine.Context, username string, email string) bool {
	query := datastore.NewQuery("Users").Ancestor(userKey(context)).
		Filter("username =", username)
	var users []maped.User
	_, err := query.GetAll(context, &users)

	if len(users) > 0 {
		return false
	}

	query = datastore.NewQuery("Users").Ancestor(userKey(context)).
		Filter("email =", email)
	_, err = query.GetAll(context, &users)
	if len(users) > 0 || err != nil {
		return false
	}
	return true
}

// RegisterUser register a new User in the datastore, if the operation
// fails then ir will return a 'false'
func RegisterUser(context appengine.Context, user maped.User) bool {
	key := datastore.NewIncompleteKey(context, "Users", routeKey(context))
	_, err := datastore.Put(context, key, &user)
	if err != nil {
		return false
	}
	return true
}

func userKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Users", "default_user", 0, nil)
}

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
