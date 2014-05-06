package routes

import (
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

	m := domain.Message{}

	m.ChannelId = params["channelId"]

	m.IsAdminMsg = true
	m.IsHtml = true

	m.Text = "new commit to <strong>" + msg.Repository.Name + "</strong> from " + msg.Commits[0].Committer.Name + ": " + msg.Commits[0].Message

	domain.MessageDomain.CreateMessage(&m)

	r.JSON(200, "ok")
}
