package main

import (
    "errors"
    "flag"
    "fmt"
    "strconv"
    "github.com/kixpanganiban/hoard/task"
)

// Holds the TaskList instance exposed to functions in this file
var taskList = &task.TaskList{Tasks: make([]task.Task, 0)}

// Handles commandline flags and arguments
func handleArgs() (error) {
    action := flag.String("action", "view", "Action [add/finish/remove/view]")
    flag.Parse()
    argTasks := flag.Args()

    switch *action {
        case "add":
            if len(argTasks) < 1 {
                    return errors.New("No tasks provided.")
                }
            for i := 0; i < len(argTasks); i ++ {
                return taskList.AddTask(task.Task{Content: argTasks[i], Status: 0})
            }
        case "finish":
            if len(argTasks) < 1 {
                    return errors.New("No tasks provided.")
                }
            for i := 0; i < len(argTasks); i++ {
                for i := 0; i < len(argTasks); i++ {
                    t, err := strconv.Atoi(argTasks[i])
                    if err != nil {
                        return errors.New("Finish arguments must be int.")
                    } else {
                        return taskList.Tasks[t].Finish()
                    }
                }
            }
            return nil
        case "remove":
            if (len(argTasks) < 1) {
                    return errors.New("No task id provided.")
            } else {
                if  len(taskList.Tasks) < 1 {
                    return errors.New("Nothing to remove.")
                }
                for i := 0; i < len(argTasks); i++ {
                    t, err := strconv.Atoi(argTasks[i])
                    if err != nil {
                        return errors.New("Remove arguments must be int.")
                    } else {
                        return taskList.RemoveTask(t)
                    }
                }
            }
        case "view":
            if  len(taskList.Tasks) < 1 {
                return errors.New("No tasks yet.")
            }
            for i := 0; i < len(taskList.Tasks); i++ {
                fmt.Printf("%d %s\n", i, taskList.Tasks[i].ToString())
            }
    }
    return nil
}

// The main function loads data from "tasks.json", ignoring any errors it might encounter.
// Then it parses the arguments to decided what to do with the data,
// then saves it back to the file.
func main() {
    _ = taskList.LoadJSON("tasks.json")
    err := handleArgs()
    if err != nil {
        fmt.Println(err)
    }
    err = taskList.SaveJSON("tasks.json")
    if err != nil {
        fmt.Println(err)
    }
}
