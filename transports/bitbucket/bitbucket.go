package bitbucket

import (
    "strings"
    "net/http"
    "fmt"
    "io/ioutil"

    "github.com/AstromechZA/potato-cli/model"
    "github.com/Jeffail/gabs"
)

const BITBUCKET_API_ENDPOINT = "https://api.bitbucket.org/1.0/"

type BitBucketTransport struct {
    TaskCache []model.ToDoTask
    User string
    Pass string
    RepoSlug string
}


func (t *BitBucketTransport) Init() error {

    // check repository exists
    r, c, err := t.apiGet("repositories/" + t.User + "/" + t.RepoSlug)
    if err != nil { return err }
    if err = t.convertHttpError(r, c); err != nil { return err }

    // parse json response
    p, err := gabs.ParseJSON(*r)
    if err != nil { return fmt.Errorf("Couldn't parse JSON: %s", string(*r)) }

    // check if issues are enabled
    hasIssues, ok := p.Path("has_issues").Data().(bool)
    if ok == false { return fmt.Errorf("Repository info did not contain 'has_issues'") }
    if hasIssues == false {
        return fmt.Errorf("Repository %s/%s does not have an issue tracker enabled", t.User, t.RepoSlug)
    }

    return nil
}


func (t *BitBucketTransport) convertHttpError(jsb *[]byte, code int) error {
    if code / 200 == 1 { return nil }

    p, err := gabs.ParseJSON(*jsb)
    if err != nil { return fmt.Errorf("Couldn't parse JSON: %s", string(*jsb)) }

    m, ok := p.Path("error.message").Data().(string)
    if ok == false { return fmt.Errorf("Couldn't extract error.message string") }

    return fmt.Errorf(m)
}


func (t *BitBucketTransport) apiCall(method string, resource string) (*[]byte, int, error) {
    resource = strings.TrimLeft(resource, "/")
    finalUrl := BITBUCKET_API_ENDPOINT + resource

    client := &http.Client{}
    req, err := http.NewRequest(method, finalUrl, nil)
    if err != nil { return nil, 0, err }
    req.SetBasicAuth(t.User, t.Pass)
    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil { return nil, 0, err }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil { return nil, 0, err }

    return &body, resp.StatusCode, nil
}


func (t *BitBucketTransport) apiGet(resource string) (*[]byte, int, error) {
    return t.apiCall("GET", resource)
}


func (t *BitBucketTransport) apiPost(resource string) (*[]byte, int, error) {
    return t.apiCall("POST", resource)
}


func (t *BitBucketTransport) apiDelete(resource string) (*[]byte, int, error) {
    return t.apiCall("DELETE", resource)
}
