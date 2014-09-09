package frontend

import (
	"encoding/hex"
	"fmt"
	"html/template"
	"log"
	"net/http"
	//"net/smtp"

	"appengine"
	"appengine/mail"

	"strconv"

	"github.com/TetAlius/GoBike/backend"
	"github.com/TetAlius/GoBike/backend/maped"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("./frontend/resources/html/register.html", "./frontend/resources/html/meta.html", "./frontend/resources/html/login.html")
	if err != nil {
		log.Fatal("Parse file error: ", err)
	}
	t.Execute(w, map[string]string{"PageTitle": "GoBike - Register"})
}

func registerPostHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	c.Infof("Trying to register a new User")
	result := backend.AvailableUserToRegister(c, r.FormValue("Username"), r.FormValue("Email"))
	c.Infof("Available user: " + strconv.FormatBool(result))
	if result {
		user := maped.User{}
		user.Username = r.FormValue("Username")
		user.Email = r.FormValue("Email")
		user.Active = false
		//Create and insert the hash for the user
		byteHash := []byte(user.Username + user.Email + user.Username + user.Email)
		hashlink := hex.EncodeToString(byteHash)
		user.HashLink = hashlink
		c.Infof(hashlink)

		registerResult := backend.RegisterUser(c, user)
		if registerResult {
			http.Redirect(w, r, "/", http.StatusFound)

			sendActivationMail(c, r.FormValue("Email"), hashlink)

			sendActivationMail(c, r.FormValue("Email"))

		} else {
			http.Redirect(w, r, "/register", http.StatusFound)
		}
	} else {
		http.Redirect(w, r, "/register", http.StatusFound)
	}
}

func sendActivationMail(context appengine.Context, userMail string, hashlink string) {
	msg := &mail.Message{
		Sender:  "Support <noreply-gobycicle@gobycicle.appspotmail.com>",
		To:      []string{userMail},
		Subject: "Activate your account on GoBike ",
		Body:    fmt.Sprintf(activationMessage, createConfirmationURL(hashlink)),
	}
}

func sendActivationMail(context appengine.Context, userEmail string) {
	msg := &mail.Message{
		Sender:  "Support <no-reply@GoBike.com>",
		To:      userMail,
		Subject: "Activate your account on GoBike ",
		Body:    fmt.Sprintf("Test ", createConfirmationURL()),
	}
	if err := mail.Send(context, msg); err != nil {
		context.Errorf("Couldn't send email: %v", err)
	}
}

var activationMessage = "Este mensaje ha sido autogenerado :) \n para activar tu cuenta haga click en el siguiente enlace o bien copielo y pequelo en el navegador :D "

func createConfirmationURL(hashLink string) string {
	return "http://http://gobycicle.appspot.com/activateUser?hashlink=" + hashLink
}

//TODO redirects
func activateUser(w http.ResponseWriter, r *http.Request) {
	context := appengine.NewContext(r)
	hashlink := r.FormValue("hashlink")
	if result := backend.ActivateUser(context, hashlink); result {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func createConfirmationURL() (string, error) {
	return "This is a test "
}
