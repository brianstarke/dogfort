package domain

import (
	"code.google.com/p/go.crypto/bcrypt"
	"github.com/nu7hatch/gouuid"
	"labix.org/v2/mgo"
)

type User struct {
	Uid          string `json:"uid"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	PasswordHash string `json:"passwordHash"`
}

type NewUser struct {
	Email    string
	Username string
	Password string
}

type UserDomain struct {
	Collection *mgo.Collection
}

/*
Creates a User from a NewUser struct.  Returns the new user's UID
*/
func (ud UserDomain) CreateUser(newUser *NewUser) (*string, error) {
	user := User{}

	// create our own UID (ignore Mongo ObjectId)
	uid, err := uuid.NewV4()

	if err != nil {
		return nil, err
	}

	// generate password hash
	b, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)

	if err != nil {
		return nil, err
	}

	user.Uid = uid.String()
	user.Email = newUser.Email
	user.Username = newUser.Username
	user.PasswordHash = string(b)

	err = ud.Collection.Insert(&user)

	if err != nil {
		return nil, err
	} else {
		return &user.Uid, nil
	}
}
