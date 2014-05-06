package domain

import (
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"crypto/md5"

	"code.google.com/p/go.crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"github.com/nu7hatch/gouuid"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type UserUid string

type User struct {
	Uid          UserUid `json:"uid"`
	Email        string  `json:"email"`
	Username     string  `json:"username"`
	PasswordHash string  `json:"passwordHash,omitempty"`
	GravatarHash string  `json:"gravatarHash"`
}

type NewUser struct {
	Email    string `binding:"required"`
	Username string `binding:"required"`
	Password string `binding:"required"`
}

type AuthenticationRequest struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

type userDomain struct {
	Collection *mgo.Collection
}

/*
Creates a User from a NewUser struct.  Returns the new user's UID
*/
func (ud userDomain) CreateUser(newUser *NewUser) (*UserUid, error) {
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
	user.Uid = UserUid(uid.String())
	user.Email = newUser.Email
	user.Username = newUser.Username
	user.PasswordHash = string(b)
	user.GravatarHash = getGravatarHash(&newUser.Email)

	err = ud.Collection.Insert(&user)

	if err != nil {
		return nil, err
	} else {
		return &user.Uid, nil
	}
}

/*
Attempts to authenticate a user and returns a JWT if successful
*/
func (ud userDomain) Authenticate(ar *AuthenticationRequest) (*string, error) {
	u, err := ud.UserByUsername(ar.Username)

	if err != nil {
		return nil, err
	}

	// check password against hash
	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(ar.Password))

	if err != nil {
		return nil, fmt.Errorf("Invalid password for %s", ar.Username)
	}

	token := jwt.New(jwt.GetSigningMethod("HS256"))

	token.Header["user_id"] = u.Uid
	token.Claims["iat"] = time.Now().Unix()
	token.Claims["exp"] = time.Now().Add(time.Hour * 24 * 30 * 6).Unix()

	// TODO!  move this signing key to .env (and maybe use rsa key)
	tokenString, err := token.SignedString([]byte("dogfort"))

	if err != nil {
		return nil, err
	} else {
		return &tokenString, nil
	}
}

func (ud userDomain) UserByUsername(username string) (*User, error) {
	u := User{}

	err := ud.Collection.Find(bson.M{"username": username}).One(&u)

	if err != nil {
		return nil, err
	} else {
		return &u, nil
	}
}

func (ud userDomain) UserByUid(uid UserUid) (*User, error) {
	u := User{}
	err := ud.Collection.Find(bson.M{"uid": uid}).One(&u)

	if err != nil {
		return nil, err
	} else {
		u.PasswordHash = "" // wipe this before sending it anywhere front end
		return &u, nil
	}
}

func getGravatarHash(email *string) string {
	hasher := md5.New()

	hasher.Write([]byte(strings.ToLower(*email)))

	return hex.EncodeToString(hasher.Sum(nil))
}

/*
Parses a token, returns user id if valid
*/
func getUserUidFromToken(token string) (*UserUid, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) ([]byte, error) {
		return []byte("dogfort"), nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %s", err.Error())
	} else {
		if t.Valid {
			uid := UserUid(t.Header["user_id"].(string))

			return &uid, nil
		} else {
			return nil, fmt.Errorf("token is not valid: %s", err.Error())
		}
	}
}
