package main

import (
    "os"
    "fmt"

    "github.com/AstromechZA/potato-cli/transports/bitbucket"
)


func mainInner() error {

    t := bitbucket.BitBucketTransport{
        User: "benm_",
        Pass: "6dmgWkBfUsFsHrys",
        RepoSlug: "testing-todos",
    }
    err := t.Init()
    if err != nil { return err }
    tasks, err := t.Read()
    fmt.Println(tasks)
    return err

}


func main() {
    if err := mainInner(); err != nil {
        os.Stderr.WriteString(err.Error() + "\n")
        os.Exit(1)
    }
}
