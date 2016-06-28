package model

import (
    "time"
)

// TaskState represents the overall state of the task
type TaskState uint8
const (
    // StateOpen is a task that should be done soon
    StateOpen TaskState = iota + 1
    // StateClosed is for tasks that have been done (or discarded)
    StateClosed
)

// TaskDue is a type used to represent WHEN a task is due
// it is a useful ordering and is also needed for notifications / alarms
type TaskDue uint8
const (
    // DueDateTime is for tasks due on specific instants
    // the specific date time is stored separately
    DueDateTime TaskDue = iota + 1
    // DueASAP is for tasks due as soon as possible
    DueASAP
    // DueSoon is for tasks due soon
    DueSoon
    // DueLater is for tasks due in the future
    DueLater
    // NotDue is for tasks that have no requirement
    NotDue
)

// ToDoTask is the state of a current task in the system
type ToDoTask struct {
    IssueID uint            `json:"issue_id"`
    Title string            `json:"title"`
    Description string      `json:"description"`
    State TaskState         `json:"state"`
    Labels []string         `json:"labels"`
    Due TaskDue             `json:"due"`
    DueDateTime time.Time   `json:"due_datetime"`
}
