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

type channelDomain struct {
	Collection *mgo.Collection
}

/*
Creates a new channel, returns the Channel ID
*/
func (cd channelDomain) CreateChannel(channel *Channel, userUid *UserUid) (*string, error) {
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

func (cd channelDomain) ListChannels() (*[]Channel, error) {
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
func (cd channelDomain) GetUserChannels(userUid *UserUid) (*[]Channel, error) {
	c := []Channel{}

	err := cd.Collection.Find(bson.M{"members": *userUid}).All(&c)

	if err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

/*
Checks if the user is in this channel
*/
func (cd channelDomain) UserInChannel(userUid *UserUid, channelId string) (bool, error) {
	c := Channel{}

	err := cd.Collection.Find(bson.M{"uid": channelId}).One(&c)

	if err != nil {
		return false, err
	} else {
		return isMember(&c.Members, userUid) != -1, nil
	}
}

/*
Subscribes a user to a channel, unless it's private
*/
func (cd channelDomain) SubscribeToChannel(userUid *UserUid, channelId string) error {
	c := Channel{}

	err := cd.Collection.Find(bson.M{"uid": channelId}).One(&c)

	if err != nil {
		return err
	}

	if c.Private {
		return fmt.Errorf("%s is a private channel, you need to request membership", c.Name)
	}

	if isMember(&c.Members, userUid) != -1 {
		return fmt.Errorf("User [%s] is already subscribed", userUid)
	}

	c.Members = append(c.Members, *userUid)

	err = cd.Collection.Update(bson.M{"uid": c.Uid}, &c)

	return nil
}

/*
Unsubscribes a user from a channel
*/
func (cd channelDomain) UnsubscribeFromChannel(userUid *UserUid, channelId string) error {
	c := Channel{}

	err := cd.Collection.Find(bson.M{"uid": channelId}).One(&c) // TODO i feel like I've written this before...

	if err != nil {
		return err
	}

	pos := isMember(&c.Members, userUid)

	if pos != -1 {
		c.Members = c.Members[:pos+copy(c.Members[pos:], c.Members[pos+1:])]

		err = cd.Collection.Update(bson.M{"uid": c.Uid}, &c)

		if err != nil {
			return err
		} else {
			return nil
		}
	} else {
		return fmt.Errorf("User [%s] is not subscribed to this channel", userUid)
	}
}

func isMember(members *[]UserUid, user *UserUid) int {
	for p, u := range *members {
		if u == *user {
			return p
		}
	}
	return -1
}
