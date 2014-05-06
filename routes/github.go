package routes

import (
	"bytes"
	"text/template"

	"github.com/brianstarke/dogfort/domain"
	"github.com/brianstarke/dogfort/hub"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

var (
	tmplStr = `<strong>GitHub</strong> new commits to <a href='{{.Repository.Url}}' target='_blank'>{{.Repository.Name}}</a>&nbsp;
  <small>(<a href='{{.CompareUrl}} target='_blank'>compare</a>)</small><br>
{{range .Commits}}<small>- <em>{{.Committer.Username}}</em> :: {{.Message}}</small><br>{{end}}
`
	commitTmpl, _ = template.New("commitTemplate").Parse(tmplStr)
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

	var b bytes.Buffer

	_ = commitTmpl.Execute(&b, msg)
	m.Text = b.String()

	mId, _ := domain.MessageDomain.CreateMessage(&m)
	hub.H.MessagePublish(m.ChannelId, mId)

	r.JSON(200, "ok")
}
