package domain

import (
	"fmt"

	"github.com/nu7hatch/gouuid"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Channel struct {
	Uid         string    `json:"uid"`
	Name        string    `json:"name",binding:"required"`
	Description string    `json:"description",binding:"required"`
	Private     bool      `json:"isPrivate",binding:"required"`
	CreatedBy   UserUid   `json:"createdBy"` // the user UID of the admin
	Members     []UserUid `json:"members"`
}

type ChannelDomain struct {
	Collection *mgo.Collection
}

/*
Creates a new channel, returns the Channel ID
*/
func (cd ChannelDomain) CreateChannel(channel *Channel, userUid *UserUid) (*string, error) {
	// check to see if there is already a channel with this name
	n, err := cd.Collection.Find(bson.M{"name": channel.Name}).Count()

	if err != nil {
		return nil, err
	}

	if n > 0 {
		return nil, fmt.Errorf("There is already a channel with that name [%s]", channel.Name)
	}

	// create our own UID (ignore Mongo ObjectId)
	uid, err := uuid.NewV4()

	if err != nil {
		return nil, err
	}

	channel.CreatedBy = *userUid
	channel.Uid = uid.String()
	channel.Members = []UserUid{*userUid}

	err = cd.Collection.Insert(&channel)

	if err != nil {
		return nil, err
	} else {
		return &channel.Uid, nil
	}
}

func (cd ChannelDomain) ListChannels() (*[]Channel, error) {
	c := []Channel{}

	err := cd.Collection.Find(bson.M{}).All(&c)

	if err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

/*
Returns all channels this user is in
*/
func (cd ChannelDomain) GetUserChannels(userUid *UserUid) (*[]Channel, error) {
	c := []Channel{}

	err := cd.Collection.Find(bson.M{"members": *userUid}).All(&c)

	if err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}
