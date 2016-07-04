package main

import (
    "os"

    "github.com/AstromechZA/potato-cli/transports/bitbucket"
)


func mainInner() error {

    t := bitbucket.BitBucketTransport{
        User: "benm_",
        Pass: "6dmgWkBfUsFsHrys",
        RepoSlug: "testing-todos",
    }
    return t.Init()
}


func main() {
    if err := mainInner(); err != nil {
        os.Stderr.WriteString(err.Error() + "\n")
        os.Exit(1)
    }
}
