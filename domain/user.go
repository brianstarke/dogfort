package domain

import (
	"fmt"

	"code.google.com/p/go.crypto/bcrypt"
	"github.com/nu7hatch/gouuid"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type User struct {
	Uid          string `json:"uid"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	PasswordHash string `json:"passwordHash"`
}

type NewUser struct {
	Email    string `binding:"required"`
	Username string `binding:"required"`
	Password string `binding:"required"`
}

type UserDomain struct {
	Collection *mgo.Collection
}

/*
Creates a User from a NewUser struct.  Returns the new user's UID
*/
func (ud UserDomain) CreateUser(newUser *NewUser) (*string, error) {
	// check to see if this user already exists
	n, err := ud.Collection.Find(bson.M{"$or": []bson.M{bson.M{"email": newUser.Email}, bson.M{"username": newUser.Username}}}).Count()

	if err != nil {
		return nil, err
	}

	if n > 0 {
		return nil, fmt.Errorf("A user already exists with that username or email")
	}

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

	user := User{}
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
