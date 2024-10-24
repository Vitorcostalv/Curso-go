package tasks

import (
    "encoding/json"
    "io/ioutil"
    "os"
    "fmt"
)

const dataFile = "data/tasks.json"

func LoadTasks() ([]Task, error) {
    var tasks []Task
    file, err := os.Open(dataFile)
    if err != nil {
        return tasks, err
    }
    defer file.Close()

    bytes, _ := ioutil.ReadAll(file)
    json.Unmarshal(bytes, &tasks)
    return tasks, nil
}

func SaveTasks(tasks []Task) error {
    data, _ := json.MarshalIndent(tasks, "", "    ")
    return ioutil.WriteFile(dataFile, data, 0644)
}

func AddTask(description string) error {
    tasks, err := LoadTasks()
    if err != nil {
        return err
    }
    newTask := Task{
        ID:          len(tasks) + 1,
        Description: description,
        Completed:   false,
    }
    tasks = append(tasks, newTask)
    return SaveTasks(tasks)
}

func ListTasks() ([]Task, error) {
    return LoadTasks()
}

func CompleteTask(id int) error {
    tasks, err := LoadTasks()
    if err != nil {
        return err
    }
    for i := range tasks {
        if tasks[i].ID == id {
            tasks[i].Completed = true
            break
        }
    }
    return SaveTasks(tasks)
}

func RemoveTask(id int) error {
    tasks, err := LoadTasks()
    if err != nil {
        return err
    }
    for i := range tasks {
        if tasks[i].ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            break
        }
    }
    return SaveTasks(tasks)
}
