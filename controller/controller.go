package controller

import (
    "github.com/AstromechZA/potato-cli/transport"
    "github.com/AstromechZA/potato-cli/config"
    "github.com/AstromechZA/potato-cli/model"
)

type Controller struct {
    Transport *transport.Transport,
    TasksCache []model.Tasks,
    NextTaskId uint
}

func NewController(conf config.Config) *Controller {
    return &Controller{
        Transport: transport.NewTransport(conf)
        TasksCache: nil
    }
}

func (c *Controller) Init() error {
    isnew, err := c.Transport.CheckAndSetup()
    if err != nil { return err }

    if isnew {
        c.TasksCache = make([]model.ToDoTask, 0)
    } else {
        tasks, err := c.Transport.Read()
        if err != nil { return err }
        c.TasksCache = *tasks
    }
    biggestTaskId := -1
    for _, t := range c.TasksCache {
        if t.ID > biggestTaskId {
            biggestTaskId = t.ID
        }
    }
    c.NextTaskId = biggestTaskId + 1
    return nil
}

