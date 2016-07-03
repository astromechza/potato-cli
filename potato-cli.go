package main

import (
    "os"
    "flag"
    "fmt"

    "github.com/AstromechZA/potato-cli/config"
    "github.com/AstromechZA/potato-cli/transport"
    "github.com/AstromechZA/potato-cli/model"
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

    tasks, err := transport.Read()
    if err != nil { return err }

    tasksS := append(*tasks, model.ToDoTask{

    })
    tasks = &tasksS

    err = transport.Write(tasks)
    fmt.Println(err)

    return nil
}

func main() {
    if err := mainInner(); err != nil {
        os.Stderr.WriteString(err.Error() + "\n")
        os.Exit(1)
    }
}
