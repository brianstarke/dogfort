package domain

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"log"

	"github.com/nu7hatch/gouuid"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Message struct {
	Uid        string    `json:"uid"`
	ChannelId  string    `json:"channelId",binding:"required"`
	UserId     UserUid   `json:"userId"`
	Text       string    `json:"text",binding:"required"`
	HasImage   bool      `json:"hasImage",omitempty`
	Attachment string    `json:"attachment",omitempty`
	Timestamp  time.Time `json:"timestamp"` // Unix time, in seconds
	IsAdminMsg bool      `json:"isAdminMsg",omitempty`
	IsHtml     bool      `json:"isHtml",omitempty`
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

	// eat error, we don't care yet
	md.addAttachments(message)

	err = md.Collection.Insert(&message)

	if err != nil {
		return nil, err
	} else {
		return &message.Uid, nil
	}
}

/*
Gets all messages for a channel
*/
func (md messageDomain) MessagesByChannel(userUid *UserUid, channelId string, beforeTs string, numMessages string) (*[]Message, error) {
	m := []Message{}

	limit, err := strconv.ParseInt(numMessages, 10, 0)
	res, err := strconv.ParseInt(beforeTs, 10, 64)
	ts := time.Unix(res/1000, 0)

	if err != nil {
		return nil, err
	}

	err = md.Collection.Find(bson.M{"channelid": channelId, "timestamp": bson.M{"$lte": ts}}).Limit(int(limit)).Sort("-timestamp").All(&m)

	if err != nil {
		return nil, err
	} else {
		return &m, nil
	}
}

func (md messageDomain) MessageById(id string) (*Message, error) {
	m := Message{}

	err := md.Collection.Find(bson.M{"uid": id}).One(&m)

	if err != nil {
		return nil, err
	} else {
		return &m, nil
	}
}

func (md messageDomain) addAttachments(message *Message) error {
	re, err := regexp.Compile("https?://[^\\s<>\"]+|www\\.[^\\s<>\"]+")

	if err != nil {
		return err
	}

	b := re.Find([]byte(message.Text))

	if len(b) > 0 && md.getType(string(b)) == "IMAGE" {
		message.HasImage = true
		message.Attachment = string(b)
	} else {
		message.HasImage = false
	}

	return nil
}

/**
Only valid type right now is "IMAGE" or "UNKNOWN"
**/
func (md messageDomain) getType(url string) string {
	resp, err := http.Head(url)

	if err != nil {
		log.Printf("Error while getting resource type : %s", err.Error())
		return "UNKNOWN"
	}

	header := resp.Header["Content-Type"]

	if strings.Index(header[0], "image") != -1 {
		return "IMAGE"
	} else {
		return "UNKNOWN"
	}
}
