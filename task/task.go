// Package task provides the Task and TaskList data structures
package task

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
)

type Task struct {
    Content string
    Status int
}

// Converts the Task to a readable string line
func (t *Task) ToString() (string) {
    var status string
    switch t.Status {
        case 0:
            status = "[ ]"
        case 1:
            status = "[x]"
    }
    return fmt.Sprintf("%s %s", status, t.Content)
}

// Sets the Task status to 1
func (t *Task) Finish() (error) {
    t.Status = 1
    return nil
}

type TaskList struct {
    Tasks []Task `json:"tasks"`
}

// Adds a Task to the TaskList's Tasks
func (t *TaskList) AddTask(task Task) (error) {
    t.Tasks = append(t.Tasks, task)
    return nil
}

// Removes a Task from the TaskList's Tasks by index
func (t *TaskList) RemoveTask(i int) (error) {
    t.Tasks = append(t.Tasks[:i], t.Tasks[i:1]...)
    return nil
}

// Loads a TaskList from a JSON file, takes file path as argument
func (t *TaskList) LoadJSON(file string) (error) {
    dat, err := ioutil.ReadFile(file)
    if err != nil {
        return err
    }
    if err = json.Unmarshal(dat, &t); err != nil {
        return err
    }
    return nil
}

// Saves a TaskList into a JSON file, takes file path as argument
func (t *TaskList) SaveJSON(file string) (error) {
    dat, err := json.Marshal(t)
    if err != nil {
        return err
    }
    err = ioutil.WriteFile(file, dat, 0644)
    if err != nil {
        return err
    }
    return nil
}
