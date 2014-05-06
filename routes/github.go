package routes

import (
	"log"

	"github.com/brianstarke/dogfort/domain"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type GithubMsg struct {
	CompareUrl string           `json:"compare"`
	Commits    []GithubCommits  `json:"commits"`
	Repository GithubRepository `json:"repository"`
}

type GithubCommits struct {
	Message   string     `json:"message"`
	Author    GithubUser `json:"author"`
	Committer GithubUser `json:"committer"`
}

type GithubUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type GithubRepository struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func GithubHandler(msg GithubMsg, params martini.Params, r render.Render) {
	userUid, err := getGithubUserId()

	if err != nil {
		r.JSON(400, err.Error())
		return
	}

	addToChannel(params["channelId"], userUid)

	m := domain.Message{}

	m.ChannelId = params["channelId"]
	m.UserId = *userUid
	m.Text = "new commit to " + msg.Repository.Name + " from " + msg.Commits[0].Committer.Name + ": " + msg.Commits[0].Message

	domain.MessageDomain.CreateMessage(&m)

	r.JSON(200, "ok")
}

/*
Add the github user to this channel if it's not there already
*/
func addToChannel(channelId string, user *domain.UserUid) {
	inChannel, _ := domain.ChannelDomain.UserInChannel(user, channelId)

	if !inChannel {
		err := domain.ChannelDomain.SubscribeToChannel(user, channelId)
		if err != nil {
			log.Print(err.Error())
		}
	}
}

func getGithubUserId() (*domain.UserUid, error) {
	user, err := domain.UserDomain.UserByUsername("GitHub")

	if err != nil {
		// no github user yet, create
		n := domain.NewUser{"github@dogfort.io", "GitHub", "lolineverlogin123456789"}

		id, err := domain.UserDomain.CreateUser(&n)

		if err != nil {
			return nil, err
		} else {
			return id, nil
		}
	} else {
		return &user.Uid, nil
	}

}
