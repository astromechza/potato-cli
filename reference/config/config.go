package config

import (
    "os"
    "os/user"
    "path/filepath"
    "encoding/json"
)

const DefaultConfigFile = "potato-todo.conf"

type Config struct {
    Token string `json:"token"`
    User string `json:"user"`
    GistName string `json:"gistname"`
}

func Load(path string) (*Config, error) {
    if path == "" {
        u, err := user.Current()
        if err != nil { return nil, err }
        path = filepath.Join(u.HomeDir, ".ssh", DefaultConfigFile)
    }

    fd, err := os.Open(path)
    if err != nil { return nil, err }

    var c Config
    d := json.NewDecoder(fd)
    if err = d.Decode(&c); err != nil { return nil, err }

    return &c, nil
}

