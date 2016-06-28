package main

import (
    "os"
    "flag"
    "fmt"

    "github.com/AstromechZA/potato-cli/config"
    "github.com/AstromechZA/potato-cli/transport"
)

const usageString =
`blah blah

`

func mainInner() error {
    // define flags
    confFileFlag := flag.String("conf-file", "", "A file containing Github access details")
    //verboseFlag := flag.Bool("debug", false, "Show debug messages")

    // set a more verbose usage message.
    flag.Usage = func() {
        os.Stderr.WriteString(usageString)
        flag.PrintDefaults()
    }
    // parse them
    flag.Parse()

    configObj, err := config.Load(*confFileFlag)
    if err != nil { return err }

    transport := transport.NewTransport(*configObj)

    err = transport.CheckAndSetup()
    if err != nil { return err }

    fmt.Println(transport.List())

    return nil
}

func main() {
    if err := mainInner(); err != nil {
        os.Stderr.WriteString(err.Error() + "\n")
        os.Exit(1)
    }
}
