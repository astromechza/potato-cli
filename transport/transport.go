package transport

import (
    "fmt"
    "crypto/md5"
    "net/http"
    "io/ioutil"
    "encoding/json"

    "github.com/google/go-github/github"
    "golang.org/x/oauth2"

    "github.com/AstromechZA/potato-cli/model"
    "github.com/AstromechZA/potato-cli/config"
)


const GistNamePrefix = "potato-tasks-"


type tokenSource struct {
    AccessToken string
}


func (t *tokenSource) Token() (*oauth2.Token, error) {
    return &oauth2.Token{AccessToken: t.AccessToken}, nil
}


type Transport struct {
    Client *github.Client
    User string
    GistName string
    GistID string
}


func NewTransport(conf config.Config) *Transport {
    oauthClient := oauth2.NewClient(oauth2.NoContext, &tokenSource{
        AccessToken: conf.Token,
    })

    gistName := fmt.Sprintf(
        GistNamePrefix + "%x",
        md5.Sum([]byte(conf.User)),
    )[:len(GistNamePrefix) + 8]

    if conf.GistName != "" {
        gistName = conf.GistName
    }

    return &Transport{
        Client: github.NewClient(oauthClient),
        User: conf.User,
        GistName: gistName,
    }
}


func (t *Transport) SearchForGist(user string, name string) (*github.Gist, error) {
    lstop := &github.GistListOptions{}
    lstop.Page = 1
    lstop.PerPage = 25

    for {
        gists, _, err := t.Client.Gists.List(user, lstop)
        if err != nil { return nil, err }
        if len(gists) == 0 {
            break
        }

        for _, g := range gists {
            for _, f := range g.Files {
                if *f.Filename == name {
                    return g, nil
                }
            }
        }

        lstop.Page++
    }
    return nil, fmt.Errorf("Could not find a Gist containing the file %s", name)
}


func (t *Transport) BuildInitialGist() *github.Gist {
    gistContents := make(map[github.GistFilename]github.GistFile)
    cntn := "[]"
    gistContents[github.GistFilename(t.GistName)] = github.GistFile{
        Filename: &t.GistName,
        Content: &cntn,
    }

    public := false
    return &github.Gist{
        Public: &public,
        Files: gistContents,
    }
}


func (t *Transport) CheckAndSetup() error {
    g, err := t.SearchForGist(t.User, t.GistName)
    if err != nil {
        u, _, err := t.Client.Users.Get(t.User)
        if err != nil { return err }

        g = t.BuildInitialGist()
        g.Owner = u

        g, _, err = t.Client.Gists.Create(g)
        if err != nil { return err }
    }
    t.GistID = *g.ID
    return nil
}


type ToDoTaskArray struct {
    Tasks []model.ToDoTask     `json:"tasks"`
}


func (t *Transport) List() ([]model.ToDoTask, error) {
    g, _, err := t.Client.Gists.Get(t.GistID)
    if err != nil { return nil, err }

    f := g.Files[github.GistFilename(t.GistName)]
    res, err := http.Get(*f.RawURL)
    if err != nil { return nil, err }

    if res.StatusCode == 200 {
        defer res.Body.Close()
        content, _ := ioutil.ReadAll(res.Body)

        var tasks []model.ToDoTask
        err := json.Unmarshal(content, &tasks)
        if err != nil { return nil, err }

        return tasks, nil
    } else {
        return nil, fmt.Errorf("Could not fetch data from Github: %s", res)
    }
}


func (t *Transport) DeleteByID(id uint) error {
    return fmt.Errorf("Delete not implemented")
}


func (t *Transport) Delete(task model.ToDoTask) error {
    return t.DeleteByID(task.IssueID)
}

