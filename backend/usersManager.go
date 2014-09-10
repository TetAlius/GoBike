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
	query := datastore.NewQuery("Users").
		Filter("Username =", username).
		Filter("Active =", true)
	var users []maped.User
	_, err := query.GetAll(context, &users)

	if len(users) > 0 || err != nil {
		return false
	}

	query = datastore.NewQuery("Users").
		Filter("Email =", email).
		Filter("Active =", true)
	_, err = query.GetAll(context, &users)
	if len(users) > 0 || err != nil {
		return false
	}
	return true
}

// RegisterUser register a new User in the datastore, if the operation
// fails then ir will return a 'false'
func RegisterUser(context appengine.Context, user maped.User) bool {
	context.Infof("Registering " + user.Username)
	key := datastore.NewIncompleteKey(context, "Users", nil)
	context.Infof("Getting the datastore key")
	_, err := datastore.Put(context, key, &user)
	if err != nil {
		context.Errorf(err.Error())
		return false
	}
	return true
}

// ActivateUser activates a user with the hashLink provided
func ActivateUser(context appengine.Context, hashLink string) bool {
	query := datastore.NewQuery("Users").
		Filter("HashLink =", hashLink)
	var users []maped.User
	keys, err := query.GetAll(context, &users)
	if err != nil {
		context.Errorf("Could not recover the keys: %v", err)
		return false
	}
	users[0].Active = true
	datastore.Put(context, keys[0], &users[0])
	return true
}

//FUENTE: http://www.mschoebel.info/2014/03/09/snippet-golang-webapp-login-logout.html
/*
func login(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	pass := request.FormValue("password")
	if checkCredentials(name, pass) {
		setSession(name, response)
		redirectTarget := "/"
		http.Redirect(response, request, redirectTarget, 302)
	}
}

//CheckCredentials checks if the username and the password given are correct
func CheckCredentials(context appengine.Context, name string, password string) bool {
	q := datastore.NewQuery("Users").
		Filter("Username =", name)

	var user []maped.User
	keys, err := q.GetAll(c, &user)

	if err != nil {
		context.Errorf("Can't find user: %e", err)
		err = errors.New("Can't find user")
		return false
	}

	if user[0] != nil {
		if user[0].Password == password {
			return true
		}
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
