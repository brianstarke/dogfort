package domain

import (
	"time"

	"github.com/nu7hatch/gouuid"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Message struct {
	Uid       string    `json:"uid"`
	ChannelId string    `json:"channelId",binding:"required"`
	UserId    UserUid   `json:"userId"`
	Text      string    `json:"text",binding:"required"`
	Timestamp time.Time `json:"timestamp"` // Unix time, in seconds
}

type messageDomain struct {
	Collection *mgo.Collection
}

/*
Creates a new message, returns the message ID
*/
func (md messageDomain) CreateMessage(message *Message) (*string, error) {
	// TODO validate channel ID and user ID

	uid, err := uuid.NewV4()

	if err != nil {
		return nil, err
	}

	message.Uid = uid.String()
	message.Timestamp = time.Now()

	err = md.Collection.Insert(&message)

	if err != nil {
		return nil, err
	} else {
		return &message.Uid, nil
	}
}

/*
Gets all messages for a channel

TODO add filtering, pagination
*/
func (md messageDomain) MessagesByChannel(userUid *UserUid, channelId string) (*[]Message, error) {
	m := []Message{}

	err := md.Collection.Find(bson.M{"channelid": channelId}).All(&m)

	if err != nil {
		return nil, err
	} else {
		return &m, nil
	}
}
