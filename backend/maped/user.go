package maped

import (
	"time"
)

// User the users
type User struct {
	username        string    // the user name
	since       time.Time // the day when the account was created
	alreadyDone []Route
	email		string
	name 		string
	surname 	string
	birthDate	string
	password	string
}
