package maped

import (
	"time"
)

// User the users
type User struct {
	Username    string    // the user name
	Since       time.Time // the day when the account was created
	AlreadyDone []Route
	Email       string
	Name        string
	Surname     string
	BirthDate   string
	Password    string
}
