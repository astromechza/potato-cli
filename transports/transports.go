package transports

import (
    "github.com/AstromechZA/potato-cli/model"
)

// A TodoTaskTransport is a way of storing and fetching a list
// of TodoTask models from some external source.
type TodoTaskTransport interface {
    Init() error
    Read() (*[]model.TodoTask, error)
    Write(*[]model.TodoTask) error
}
